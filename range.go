// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import "math/bits"

// Minimum get the smallest value stored in this roaring bitmap, assumes that it is not empty

// FirstZero finds the first zero bit and returns its index, assuming the map is not empty.
func (dst Bitmap) FirstZero() (uint32, bool) {
	for blkAt, blk := range dst {
		if blk != 0xffffffffffffffff {
			return uint32(blkAt<<6 + bits.LeadingZeros64(blk)), true
		}
	}
	return 0, false
}

// Range iterates over the bitmap elements. If the callback returns false it halts
// the iteration.
func (dst Bitmap) Range(f func(x uint32) bool) {
	for blkAt, blk := range dst {
		offset := 0
		if (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
		if offset++; (blk&(1<<offset)) > 0 && !f(uint32(blkAt+offset)) {
			return
		}
	}

	// Naive implementation:
	/*
		idx := uint32(0)
		for _, blk := range dst {
			for bitAt := 0; bitAt < 64; bitAt++ {
				if (blk & (1 << bitAt)) > 0 {
					if !f(idx) {
						return
					}
				}
				idx++
			}
		}
	*/
}
