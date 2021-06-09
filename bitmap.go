// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"math/bits"

	"github.com/kelindar/bitmap/simd"
)

// Bitmap represents a scalar-backed bitmap index
type Bitmap []uint64

// Set sets the bit x in the bitmap and grows it if necessary.
func (dst *Bitmap) Set(x uint32) {
	blkAt := int(x >> 6)
	bitAt := int(x % 64)
	if size := len(*dst); blkAt >= size {
		dst.grow(blkAt)
	}

	(*dst)[blkAt] |= (1 << bitAt)
}

// Remove removes the bit x from the bitmap, but does not shrink it.
func (dst *Bitmap) Remove(x uint32) {
	if blkAt := int(x >> 6); blkAt < len(*dst) {
		bitAt := int(x % 64)
		(*dst)[blkAt] &^= (1 << bitAt)
	}
}

// Contains checks whether a value is contained in the bitmap or not.
func (dst Bitmap) Contains(x uint32) bool {
	blkAt := int(x >> 6)
	if size := len(dst); blkAt >= size {
		return false
	}

	bitAt := int(x % 64)
	return (dst[blkAt] & (1 << bitAt)) > 0
}

// And computes the intersection between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) And(b Bitmap) {
	if dst.balance(b); len(*dst) != len(b) {
		return // Elliminate bounds check
	}

	a := *dst
	if simd.Supported {
		simd.And(a, b)
	} else {
		a := *dst
		for i := 0; i < len(a); i++ {
			a[i] = a[i] & b[i]
		}
	}
}

// AndNot computes the difference between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) AndNot(b Bitmap) {
	if dst.balance(b); len(*dst) != len(b) {
		return // Elliminate bounds check
	}

	a := *dst
	if simd.Supported {
		simd.AndNot(a, b)
	} else {
		for i := 0; i < len(a); i++ {
			a[i] = a[i] &^ b[i]
		}
	}
}

// Or computes the union between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) Or(b Bitmap) {
	if dst.balance(b); len(*dst) != len(b) {
		return // Elliminate bounds check
	}

	a := *dst
	if simd.Supported {
		simd.Or(a, b)
	} else {
		for i := 0; i < len(a); i++ {
			a[i] = a[i] | b[i]
		}
	}
}

// Xor computes the symmetric difference between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) Xor(b Bitmap) {
	if dst.balance(b); len(*dst) != len(b) {
		return // Elliminate bounds check
	}

	a := *dst
	if simd.Supported {
		simd.Xor(a, b)
	} else {
		for i := 0; i < len(a); i++ {
			a[i] = a[i] ^ b[i]
		}
	}
}

// Ones sets the entire bitmap to one
func (dst Bitmap) Ones() {
	for i := 0; i < len(dst); i++ {
		dst[i] = 0xffffffffffffffff
	}
}

// Min get the smallest value stored in this bitmap, assuming the bitmap is not empty.
func (dst Bitmap) Min() (uint32, bool) {
	for blkAt, blk := range dst {
		if blk != 0x0 {
			return uint32(blkAt<<6 + bits.TrailingZeros64(blk)), true
		}
	}

	return 0, false
}

// Max get the largest value stored in this bitmap, assuming the bitmap is not empty.
func (dst Bitmap) Max() (uint32, bool) {
	var blk uint64
	for blkAt := len(dst) - 1; blkAt >= 0; blkAt-- {
		if blk = dst[blkAt]; blk != 0x0 {
			return uint32(blkAt<<6 + (63 - bits.LeadingZeros64(blk))), true
		}
	}
	return 0, false
}

// FirstZero finds the first zero bit and returns its index, assuming the bitmap is not empty.
func (dst Bitmap) FirstZero() (uint32, bool) {
	for blkAt, blk := range dst {
		if blk != 0xffffffffffffffff {
			return uint32(blkAt<<6 + bits.TrailingZeros64(^blk)), true
		}
	}
	return 0, false
}

// grow gros whe size of the bitmap until we reach the desired block offset
func (dst *Bitmap) grow(blkAt int) {
	for i := len(*dst); i <= blkAt; i++ {
		*dst = append(*dst, 0)
	}
}

// balance grows the destination bitmap to match the size of the source bitmap.
func (dst *Bitmap) balance(src Bitmap) {
	if len(*dst) < len(src) {
		dst.grow(len(src) - 1)
	}
}
