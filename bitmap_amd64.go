// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"github.com/klauspost/cpuid/v2"
)

var (
	avx2 = cpuid.CPU.Supports(cpuid.AVX2)
)

// And computes the intersection between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) And(other Bitmap, extra ...Bitmap) {
	max := minlen(*dst, other, extra)
	dst.shrink(max)

	switch avx2 {
	case true:
		switch len(extra) {
		case 0:
			x64and1(*dst, other[:max])
		case 1:
			x64and2(*dst, other[:max], extra[0][:max])
		case 2:
			x64and3(*dst, other[:max], extra[0][:max], extra[1][:max])
		case 3:
			x64and4(*dst, other[:max], extra[0][:max], extra[1][:max], extra[2][:max])
		default:
			x64and4(*dst, other[:max], extra[0][:max], extra[1][:max], extra[2][:max])
			for i := 3; i < len(extra); i++ {
				x64and1(*dst, extra[i][:max])
			}
		}
	default:
		and(*dst, max, other, extra)
		return
	}
}

// AndNot computes the difference between two bitmaps and stores the result in the current bitmap.
// Operation works as set subtract: dst - b
func (dst *Bitmap) AndNot(other Bitmap, extra ...Bitmap) {
	max := minlen(*dst, other, extra)

	switch avx2 {
	case true:
		switch len(extra) {
		case 0:
			x64andn1(*dst, other[:max])
		case 1:
			x64andn2(*dst, other[:max], extra[0][:max])
		case 2:
			x64andn3(*dst, other[:max], extra[0][:max], extra[1][:max])
		case 3:
			x64andn4(*dst, other[:max], extra[0][:max], extra[1][:max], extra[2][:max])
		default:
			x64andn4(*dst, other[:max], extra[0][:max], extra[1][:max], extra[2][:max])
			for i := 3; i < len(extra); i++ {
				x64andn1(*dst, extra[i][:max])
			}
		}
	default:
		andn(*dst, max, other, extra)
		return
	}
}

// Or computes the union between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) Or(other Bitmap, extra ...Bitmap) {
	max := maxlen(*dst, other, extra)
	dst.grow(max - 1)

	switch avx2 {
	case true:
		switch len(extra) {
		case 0:
			x64or1(*dst, other)
		case 1:
			x64or2(*dst, other, extra[0])
		case 2:
			x64or3(*dst, other, extra[0], extra[1])
		case 3:
			x64or4(*dst, other, extra[0], extra[1], extra[2])
		default:
			x64or4(*dst, other, extra[0], extra[1], extra[2])
			for i := 3; i < len(extra); i++ {
				x64or1(*dst, extra[i])
			}
		}
	default:
		or(*dst, other, extra)
	}
}

// Xor computes the symmetric difference between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) Xor(other Bitmap, extra ...Bitmap) {
	max := maxlen(*dst, other, extra)
	dst.grow(max - 1)

	switch avx2 {
	case true:
		switch len(extra) {
		case 0:
			x64xor1(*dst, other)
		case 1:
			x64xor2(*dst, other, extra[0])
		case 2:
			x64xor3(*dst, other, extra[0], extra[1])
		case 3:
			x64xor4(*dst, other, extra[0], extra[1], extra[2])
		default:
			x64xor4(*dst, other, extra[0], extra[1], extra[2])
			for i := 3; i < len(extra); i++ {
				x64xor1(*dst, extra[i])
			}
		}
	default:
		xor(*dst, other, extra)
	}
}

// Count returns the number of elements in this bitmap
func (dst Bitmap) Count() int {
	switch avx2 {
	case true:
		return x64count_avx2(dst)
	default:
		return count(dst)
	}
}
