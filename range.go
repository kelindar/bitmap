// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"unsafe"
)

// Range iterates over all of the bits set to one in this bitmap.
func (dst Bitmap) Range(fn func(x uint32)) {
	for blkAt := 0; blkAt < len(dst); blkAt++ {
		blk := (dst)[blkAt]
		if blk == 0x0 {
			continue // Skip the empty page
		}

		// Iterate in a 4-bit chunks so we can reduce the number of function calls and skip
		// the bits for which we should not call our range function.
		offset := uint32(blkAt << 6)
		for ; blk > 0; blk = blk >> 4 {
			switch blk & 0b1111 {
			case 0b0001:
				fn(offset + 0)
			case 0b0010:
				fn(offset + 1)
			case 0b0011:
				fn(offset + 0)
				fn(offset + 1)
			case 0b0100:
				fn(offset + 2)
			case 0b0101:
				fn(offset + 0)
				fn(offset + 2)
			case 0b0110:
				fn(offset + 1)
				fn(offset + 2)
			case 0b0111:
				fn(offset + 0)
				fn(offset + 1)
				fn(offset + 2)
			case 0b1000:
				fn(offset + 3)
			case 0b1001:
				fn(offset + 0)
				fn(offset + 3)
			case 0b1010:
				fn(offset + 1)
				fn(offset + 3)
			case 0b1011:
				fn(offset + 0)
				fn(offset + 1)
				fn(offset + 3)
			case 0b1100:
				fn(offset + 2)
				fn(offset + 3)
			case 0b1101:
				fn(offset + 0)
				fn(offset + 2)
				fn(offset + 3)
			case 0b1110:
				fn(offset + 1)
				fn(offset + 2)
				fn(offset + 3)
			case 0b1111:
				fn(offset + 0)
				fn(offset + 1)
				fn(offset + 2)
				fn(offset + 3)
			}
			offset += 4
		}
	}
}

// Filter predicate
type predicate = func(x uint32) byte

// Filter iterates over the bitmap elements and calls a predicate provided for each
// containing element. If the predicate returns false, the bitmap at the element's
// position is set to zero.
func (dst *Bitmap) Filter(f func(x uint32) bool) {
	fn := *(*predicate)(unsafe.Pointer(&f))
	for blkAt := 0; blkAt < len(*dst); blkAt++ {
		blk := (*dst)[blkAt]
		if blk == 0x0 {
			continue // Skip the empty page
		}

		offset := uint32(blkAt << 6)
		var mask uint64
		var i uint32

		// Iterate in a 4-bit chunks so we can reduce the number of function calls and skip
		// the bits for which we should not call our filter function.
		for ; blk > 0; blk = blk >> 4 {
			switch blk & 0b1111 {
			case 0b0001:
				mask |= uint64(fn(offset)) << i
			case 0b0010:
				mask |= uint64(fn(offset+1)<<1) << i
			case 0b0011:
				mask |= uint64(fn(offset)|(fn(offset+1)<<1)) << i
			case 0b0100:
				mask |= uint64(fn(offset+2)<<2) << i
			case 0b0101:
				mask |= uint64(fn(offset)|fn(offset+2)<<2) << i
			case 0b0110:
				mask |= uint64((fn(offset+1)<<1)|(fn(offset+2)<<2)) << i
			case 0b0111:
				mask |= uint64(fn(offset)|(fn(offset+1)<<1)|(fn(offset+2)<<2)) << i
			case 0b1000:
				mask |= uint64(fn(offset+3)<<3) << i
			case 0b1001:
				mask |= uint64(fn(offset)|(fn(offset+3)<<3)) << i
			case 0b1010:
				mask |= uint64((fn(offset+1)<<1)|(fn(offset+3)<<3)) << i
			case 0b1011:
				mask |= uint64(fn(offset)|(fn(offset+1)<<1)|(fn(offset+3)<<3)) << i
			case 0b1100:
				mask |= uint64((fn(offset+2)<<2)|(fn(offset+3)<<3)) << i
			case 0b1101:
				mask |= uint64(fn(offset)|(fn(offset+2)<<2)|(fn(offset+3)<<3)) << i
			case 0b1110:
				mask |= uint64((fn(offset+1)<<1)|(fn(offset+2)<<2)|(fn(offset+3)<<3)) << i
			case 0b1111:
				mask |= uint64(fn(offset)|(fn(offset+1)<<1)|(fn(offset+2)<<2)|(fn(offset+3)<<3)) << i
			}

			i += 4
			offset += 4
		}

		// Apply the mask
		(*dst)[blkAt] &= mask
	}
}
