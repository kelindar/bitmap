// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

// +build ignore

package main

//go:generate go run amd64.go -out ../simd_amd64.s -stubs ../simd_amd64.go -pkg=bitmap

import (
	"strings"

	. "github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/buildtags"
	. "github.com/mmcloughlin/avo/operand"
	. "github.com/mmcloughlin/avo/reg"
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
	makeOp("and")
	makeOp("andnot")
	makeOp("or")
	makeOp("xor")

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

// makeOp generates an SIMD "and", "or" , "andnot", "xor" operations.
func makeOp(op string) {
	switch op {
	case "and":
		TEXT("x64and", NOSPLIT, "func(a []uint64, b []uint64)")
		Doc("x64and (AND) computes the intersection between two slices and stores the result in the first one")
	case "or":
		TEXT("x64or", NOSPLIT, "func(a []uint64, b []uint64)")
		Doc("x64or (OR) computes the union between two slices and stores the result in the first one")
	case "andnot":
		TEXT("x64andn", NOSPLIT, "func(a []uint64, b []uint64)")
		Doc("x64andn (AND NOT) computes the difference between two slices and stores the result in the first one")
	case "xor":
		TEXT("x64xor", NOSPLIT, "func(a []uint64, b []uint64)")
		Doc("x64xor (XOR) computes the symmetric difference between two slices and stores the result in the first one")
	}

	// Load the a and b addresses as well as the current len(a). Assume len(a) == len(b)
	Pragma("noescape")
	a := Mem{Base: Load(Param("a").Base(), GP64())}
	b := Mem{Base: Load(Param("b").Base(), GP64())}
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
	for i := 0; i < unroll; i++ {
		VMOVUPD(b.Offset(size*i), vector[i])
	}

	// Perform the actual operation
	for i := 0; i < unroll; i++ {
		switch op {
		case "and":
			VPAND(a.Offset(size*i), vector[i], vector[i])
		case "or":
			VPOR(a.Offset(size*i), vector[i], vector[i])
		case "andnot":
			VPANDN(a.Offset(size*i), vector[i], vector[i])
		case "xor":
			VPXOR(a.Offset(size*i), vector[i], vector[i])
		}
	}

	// Move the result to "a" by copying the vector
	for i := 0; i < unroll; i++ {
		VMOVUPD(vector[i], a.Offset(size*i))
	}

	// Continue the iteration
	Comment("continue the interation by moving read pointers")
	ADDQ(U32(blocksize), a.Base)
	ADDQ(U32(blocksize), b.Base)
	SUBQ(U32(4*unroll), n)
	JMP(LabelRef("body"))

	// Now, we only have less than 512 bits left, use normal scalar operation
	Label("tail")
	CMPQ(n, Imm(0))
	JE(LabelRef("done"))

	// Perform the actual operation
	Commentf("perform the logical \"%v\" operation", strings.ToUpper(op))
	MOVQ(Mem{Base: b.Base}, s)
	switch op {
	case "and":
		ANDQ(Mem{Base: a.Base}, s)
	case "or":
		ORQ(Mem{Base: a.Base}, s)
	case "andnot":
		ANDNQ(Mem{Base: a.Base}, s, s)
	case "xor":
		XORQ(Mem{Base: a.Base}, s)
	}
	MOVQ(s, Mem{Base: a.Base})

	// Continue the iteration
	Comment("continue the interation by moving read pointers")
	ADDQ(U32(8), a.Base)
	ADDQ(U32(8), b.Base)
	SUBQ(U32(1), n)
	JMP(LabelRef("tail"))

	Label("done")
	RET()
}
