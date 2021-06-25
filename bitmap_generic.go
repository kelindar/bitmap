// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

// +build !amd64

package bitmap

// And computes the intersection between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) And(b Bitmap) {
	if dst.balance(b); len(*dst) >= len(b) {
		and(dst, b)
	}

}

// AndNot computes the difference between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) AndNot(b Bitmap) {
	if dst.balance(b); len(*dst) >= len(b) {
		andn(dst, b)
	}
}

// Or computes the union between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) Or(b Bitmap) {
	if dst.balance(b); len(*dst) >= len(b) {
		or(dst, b)
	}
}

// Xor computes the symmetric difference between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) Xor(b Bitmap) {
	if dst.balance(b); len(*dst) >= len(b) {
		xor(dst, b)
	}
}

// Count returns the number of elements in this bitmap
func (dst Bitmap) Count() int {
	return count(dst)
}
