// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"github.com/klauspost/cpuid/v2"
)

var (
	avx2 = cpuid.CPU.Supports(cpuid.AVX2)
	popc = cpuid.CPU.Supports(cpuid.POPCNT)
)

// And computes the intersection between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) And(b Bitmap) {
	if dst.balance(b); len(*dst) >= len(b) {
		switch avx2 {
		case true:
			x64and(*dst, b)
		default:
			and(dst, b)
		}
	}
}

// AndNot computes the difference between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) AndNot(b Bitmap) {
	if dst.balance(b); len(*dst) >= len(b) {
		switch avx2 {
		case true:
			x64andn(*dst, b)
		default:
			andn(dst, b)
		}
	}
}

// Or computes the union between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) Or(b Bitmap) {
	if dst.balance(b); len(*dst) >= len(b) {
		switch avx2 {
		case true:
			x64or(*dst, b)
		default:
			or(dst, b)
		}
	}
}

// Xor computes the symmetric difference between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) Xor(b Bitmap) {
	if dst.balance(b); len(*dst) >= len(b) {
		switch avx2 {
		case true:
			x64xor(*dst, b)
		default:
			xor(dst, b)
		}
	}
}

// Count returns the number of elements in this bitmap
func (dst Bitmap) Count() int {
	if popc {
		return int(x64count(dst))
	}
	return count(dst)
}
