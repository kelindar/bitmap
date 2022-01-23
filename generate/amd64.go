// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.
// Initial code adapted from Marko Kevac (see https://github.com/mkevac/gopherconrussia2019)

//go:build ignore
// +build ignore

package main

//go:generate go run amd64.go -out ../simd_amd64.s -stubs ../simd_amd64.go -pkg=bitmap

import (
	"fmt"
	"strings"

	. "github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/buildtags"
	. "github.com/mmcloughlin/avo/operand"
	. "github.com/mmcloughlin/avo/reg"
)

const (
	opAnd  = "and"
	opAndn = "andn"
	opOr   = "or"
	opXor  = "xor"
)

func main() {
	Constraint(buildtags.Not("appengine").ToConstraint())
	Constraint(buildtags.Not("noasm").ToConstraint())
	Constraint(buildtags.Term("gc").ToConstraint())

	// Generate count
	// Useful examples: https://github.com/WojciechMula/sse-popcount
	// TODO: explore Harley-Seal with AVX2/512 https://arxiv.org/pdf/1611.07612.pdf
	makeCount()

	// Generate boolean algebra
	for i := 1; i <= 4; i++ {
		makeOpN(opAnd, i)
		makeOpN(opAndn, i)
		makeOpN(opOr, i)
		makeOpN(opXor, i)
	}

	Generate()
}

// makeCount generates a faster way of counting ones using popcount
func makeCount() {
	TEXT("x64count", NOSPLIT, "func(a []uint64) uint64")
	Doc("x64count counts the bits set to one")

	// Load the array and its length
	Pragma("noescape")
	a := Mem{Base: Load(Param("a").Base(), GP64())}
	n := Load(Param("a").Len(), GP64())

	// The register for the sum, reset to zero
	sum := GP64()
	XORQ(sum, sum)

	const size, unroll = 8, 4 // bytes (64bit * 4)
	const blocksize = size * unroll

	// Create a vector
	var vector []GPVirtual
	for i := 0; i < unroll; i++ {
		vector = append(vector, GP64())
	}

	Commentf("perform vectorized operation for every block of %v bits", blocksize*8)
	Label("body")
	CMPQ(n, U32(1*unroll))
	JL(LabelRef("tail"))

	// Do the unrolled counting
	Commentf("count the bits, 4 numbers at a time")
	for i := 0; i < unroll; i++ {
		POPCNTQ(a.Offset(size*i), vector[i])
	}

	// Sum it up
	for i := 0; i < unroll; i++ {
		ADDQ(vector[i], sum)
	}

	// Continue the iteration
	Comment("continue the interation by moving read pointers")
	ADDQ(U32(blocksize), a.Base)
	SUBQ(U32(1*unroll), n)
	JMP(LabelRef("body"))

	// Now, we only have less than 512 bits left, use normal scalar operation
	Label("tail")
	CMPQ(n, Imm(0))
	JE(LabelRef("done"))

	POPCNTQ(Mem{Base: a.Base}, R15)
	ADDQ(R15, sum)

	// Continue the iteration
	Comment("continue the interation by moving read pointers")
	ADDQ(U32(8), a.Base)
	SUBQ(U32(1), n)
	JMP(LabelRef("tail"))

	// Return the sum
	Label("done")
	Store(sum, ReturnIndex(0))
	RET()
}

// makeOpN generates an SIMD "and", "or" , "andnot", "xor" operations.
func makeOpN(op string, param int) {
	names := []string{}
	for i := 0; i <= param; i++ {
		names = append(names, nameOf(i))
	}

	TEXT(fmt.Sprintf("x64%v%d", op, param), NOSPLIT,
		fmt.Sprintf("func (%v []uint64)", strings.Join(names, ", ")))
	// Doc(name + " (AND) computes the intersection between two slices and stores the result in the first one")

	// Load the a and b addresses as well as the current len(a). Assume len(a) == len(b)
	Pragma("noescape")
	ptr := []Mem{}
	for i := 0; i <= param; i++ {
		ptr = append(ptr, Mem{Base: Load(Param(nameOf(i)).Base(), GP64())})
	}
	n := Load(Param("b").Len(), GP64())

	// The register for the tail, we xor it with itself to zero out
	s := GP64()
	XORQ(s, s)

	const size, unroll = 32, 2 // bytes (256bit * 2)
	const blocksize = size * unroll

	Commentf("perform vectorized operation for every block of %v bits", blocksize*8)
	Label("body")
	CMPQ(n, U32(4*unroll))
	JL(LabelRef("tail"))

	// Create a vector
	vector := make([]VecVirtual, unroll)
	for i := 0; i < unroll; i++ {
		vector[i] = YMM()
	}

	// Move memory vector into position
	Commentf("perform the logical \"%v\" operation", strings.ToUpper(op))
	for i := 1; i <= param; i++ {
		for r := 0; r < unroll; r++ {
			VMOVUPD(ptr[i].Offset(size*r), vector[r])
		}

		// Perform the actual operation
		for r := 0; r < unroll; r++ {
			switch op {
			case opAnd:
				VPAND(ptr[0].Offset(size*r), vector[r], vector[r])
			case opOr:
				VPOR(ptr[0].Offset(size*r), vector[r], vector[r])
			case opAndn:
				VPANDN(ptr[0].Offset(size*r), vector[r], vector[r])
			case opXor:
				VPXOR(ptr[0].Offset(size*r), vector[r], vector[r])
			}
		}

		// Move the result to "a" by copying the vector
		for r := 0; r < unroll; r++ {
			VMOVUPD(vector[r], ptr[0].Offset(size*r))
		}
	}

	// Continue the iteration
	Comment("continue the interation by moving read pointers")
	for i := 0; i <= param; i++ {
		ADDQ(U32(blocksize), ptr[i].Base)
	}
	SUBQ(U32(4*unroll), n)
	JMP(LabelRef("body"))

	// Now, we only have less than 512 bits left, use normal scalar operation
	Label("tail")
	CMPQ(n, Imm(0))
	JE(LabelRef("done"))

	// Perform the actual operation
	Commentf("perform the logical \"%v\" operation", strings.ToUpper(op))
	for i := 1; i <= param; i++ {
		MOVQ(Mem{Base: ptr[i].Base}, s)
		switch op {
		case opAnd:
			ANDQ(Mem{Base: ptr[0].Base}, s)
		case opOr:
			ORQ(Mem{Base: ptr[0].Base}, s)
		case opAndn:
			ANDNQ(Mem{Base: ptr[0].Base}, s, s)
		case opXor:
			XORQ(Mem{Base: ptr[0].Base}, s)
		}
		MOVQ(s, Mem{Base: ptr[0].Base})
	}

	// Continue the iteration
	Comment("continue the interation by moving read pointers")
	for i := 0; i <= param; i++ {
		ADDQ(U32(8), ptr[i].Base)
	}
	SUBQ(U32(1), n)
	JMP(LabelRef("tail"))

	Label("done")
	RET()
}

func nameOf(i int) string {
	return string(byte(0x61 + i))
}
