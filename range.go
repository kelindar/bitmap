// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

// Range iterates over the bitmap elements. If the callback returns false it halts
// the iteration.
func (dst Bitmap) Range(f func(x uint32) bool) {
	for blkAt := 0; blkAt < len(dst); blkAt++ {
		bitAt, blk := 0, (dst)[blkAt]
		if (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			return
		}
	}
}

// Filter iterates over the bitmap elements and calls a predicate provided for each
// containing element. If the predicate returns false, the bitmap at the element's
// position is set to zero.
func (dst *Bitmap) Filter(f func(x uint32) bool) {
	for blkAt := 0; blkAt < len(*dst); blkAt++ {
		bitAt, blk := 0, (*dst)[blkAt]
		if (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
		if bitAt++; (blk&(1<<bitAt)) > 0 && !f(uint32(blkAt<<6+bitAt)) {
			(*dst)[blkAt] &^= (1 << bitAt)
		}
	}
}
