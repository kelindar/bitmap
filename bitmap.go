// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"unsafe"
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
func (dst *Bitmap) Contains(x uint32) bool {
	blkAt := int(x >> 6)
	if size := len(*dst); blkAt >= size {
		return false
	}

	bitAt := int(x % 64)
	return ((*dst)[blkAt] & (1 << bitAt)) > 0
}

// And computes the intersection between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) And(b Bitmap) {
	if dst.balance(b); len(*dst) != len(b) {
		return // Elliminate bounds check
	}

	a := *dst
	for i := 0; i < len(a); i++ {
		a[i] = a[i] & b[i]
	}
}

// AndNot computes the difference between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) AndNot(b Bitmap) {
	if dst.balance(b); len(*dst) != len(b) {
		return // Elliminate bounds check
	}

	a := *dst
	for i := 0; i < len(a); i++ {
		a[i] = a[i] &^ b[i]
	}
}

// Or computes the union between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) Or(b Bitmap) {
	if dst.balance(b); len(*dst) != len(b) {
		return // Elliminate bounds check
	}

	a := *dst
	for i := 0; i < len(a); i++ {
		a[i] = a[i] | b[i]
	}
}

// Xor computes the symmetric difference between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) Xor(b Bitmap) {
	if dst.balance(b); len(*dst) != len(b) {
		return // Elliminate bounds check
	}

	a := *dst
	for i := 0; i < len(a); i++ {
		a[i] = a[i] ^ b[i]
	}
}

// Clear clears the bitmap and resizes it to zero.
func (dst *Bitmap) Clear() {
	if size := len(*dst); size > 0 {
		ptr := unsafe.Pointer(&(*dst)[0])
		memclrNoHeapPointers(ptr, uintptr(size))
		*dst = (*dst)[:0]
	}
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

//go:linkname memclrNoHeapPointers runtime.memclrNoHeapPointers
func memclrNoHeapPointers(p unsafe.Pointer, n uintptr)
