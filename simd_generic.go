// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import "math/bits"

// Count counts the number of bits set to one
func count(arr []uint64) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += bits.OnesCount64(arr[i])
	}
	return sum
}

// and computes the intersection of multiple bitmaps
func and(a Bitmap, upper int, other Bitmap, extra []Bitmap) {
	for i := 0; i < upper; i++ {
		a[i] = a[i] & other[i]
	}

	for _, b := range extra {
		for i := 0; i < upper; i++ {
			a[i] = a[i] & b[i]
		}
	}
}

// AndNot computes the difference between two bitmaps and stores the result in the current bitmap
func andn(a Bitmap, upper int, other Bitmap, extra []Bitmap) {
	for i := 0; i < upper; i++ {
		a[i] = a[i] &^ other[i]
	}

	for _, b := range extra {
		for i := 0; i < upper; i++ {
			a[i] = a[i] &^ b[i]
		}
	}
}

// or computes the union between two bitmaps and stores the result in the current bitmap
func or(a Bitmap, other Bitmap, extra []Bitmap) {
	for i := 0; i < len(other); i++ {
		a[i] = a[i] | other[i]
	}

	for _, b := range extra {
		for i := 0; i < len(b); i++ {
			a[i] = a[i] | b[i]
		}
	}
}

// Xor computes the symmetric difference between two bitmaps and stores the result in the current bitmap
func xor(a Bitmap, other Bitmap, extra []Bitmap) {
	for i := 0; i < len(other); i++ {
		a[i] = a[i] ^ other[i]
	}

	for _, b := range extra {
		for i := 0; i < len(b); i++ {
			a[i] = a[i] ^ b[i]
		}
	}
}
