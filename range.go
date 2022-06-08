// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"unsafe"

	"github.com/kelindar/simd"
)

const full = 0xffffffffffffffff

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

// Sum computes a horizontal sum of a slice, filtered by the provided bitmap
func Sum[T simd.Number](src []T, filter Bitmap) (sum T) {
	tail := minint(len(src)/64, len(filter)) << 6 // End of 64-byte blocks
	last := minint(len(src), len(filter)*64)      // End of slice or mask

	var frame [64]T
	var i0, i1 int
	for i1 = 0; i1 < tail; i1 += 64 {
		switch filter[i1>>6] {
		case full:
			continue // Continue buffering
		case 0:
		default:
			sum += simd.Sum(leftPack(&frame, src[i1:i1+64], filter[i1>>6]))
		}

		// Flush the current buffer
		if (i1 - i0) > 0 {
			sum += simd.Sum(src[i0:i1])
		}
		i0 = i1 + 64
	}

	// Flush the accumulated buffer so far
	if (i1 - i0) > 0 {
		sum += simd.Sum(src[i0:i1])
	}

	// Process the tail
	for i := tail; i < last; i++ {
		if filter.Contains(uint32(i)) {
			sum += src[i]
		}
	}
	return sum
}

// Min finds the smallest value in a slice, filtered by the provided bitmap
func Min[T simd.Number](src []T, filter Bitmap) (min T, hit bool) {
	tail := minint(len(src)/64, len(filter)) << 6 // End of 64-byte blocks
	last := minint(len(src), len(filter)*64)      // End of slice or mask

	var frame [64]T
	var i0, i1 int
	for i1 = 0; i1 < tail; i1 += 64 {
		switch filter[i1>>6] {
		case full:
			continue // Continue buffering
		case 0:
		default:
			if m := simd.Min(leftPack(&frame, src[i1:i1+64], filter[i1>>6])); m < min || !hit {
				hit = true
				min = m
			}
		}

		// Flush the current buffer
		if (i1 - i0) > 0 {
			if m := simd.Min(src[i0:i1]); m < min || !hit {
				hit = true
				min = m
			}
		}
		i0 = i1 + 64
	}

	// Flush the accumulated buffer so far
	if (i1 - i0) > 0 {
		if m := simd.Min(src[i0:i1]); m < min || !hit {
			hit = true
			min = m
		}
	}

	// Process the tail
	for i := tail; i < last; i++ {
		if filter.Contains(uint32(i)) && (src[i] < min || !hit) {
			hit = true
			min = src[i]
		}
	}
	return
}

// Max finds the largest value in a slice, filtered by the provided bitmap
func Max[T simd.Number](src []T, filter Bitmap) (max T, hit bool) {
	tail := minint(len(src)/64, len(filter)) << 6 // End of 64-byte blocks
	last := minint(len(src), len(filter)*64)      // End of slice or mask

	var frame [64]T
	var i0, i1 int
	for i1 = 0; i1 < tail; i1 += 64 {
		switch filter[i1>>6] {
		case full:
			continue // Continue buffering
		case 0:
		default:
			if m := simd.Max(leftPack(&frame, src[i1:i1+64], filter[i1>>6])); m > max || !hit {
				hit = true
				max = m
			}
		}

		// Flush the current buffer
		if (i1 - i0) > 0 {
			if m := simd.Max(src[i0:i1]); m > max || !hit {
				hit = true
				max = m
			}
		}
		i0 = i1 + 64
	}

	// Flush the accumulated buffer so far
	if (i1 - i0) > 0 {
		if m := simd.Max(src[i0:i1]); m > max || !hit {
			hit = true
			max = m
		}
	}

	// Process the tail
	for i := tail; i < last; i++ {
		if filter.Contains(uint32(i)) && (src[i] > max || !hit) {
			hit = true
			max = src[i]
		}
	}
	return
}

// leftPack left-packs a src slice into a dst for a single block blk
func leftPack[T any](dst *[64]T, src []T, blk uint64) []T {
	offset := 0
	cursor := 0
	for ; blk > 0; blk = blk >> 4 {
		switch blk & 0b1111 {
		case 0b0001:
			dst[cursor] = src[offset+0]
			cursor += 1
		case 0b0010:
			dst[cursor] = src[offset+1]
			cursor += 1
		case 0b0011:
			dst[cursor] = src[offset+0]
			dst[cursor+1] = src[offset+1]
			cursor += 2
		case 0b0100:
			dst[cursor] = src[offset+2]
			cursor += 1
		case 0b0101:
			dst[cursor] = src[offset+0]
			dst[cursor+1] = src[offset+2]
			cursor += 2
		case 0b0110:
			dst[cursor] = src[offset+1]
			dst[cursor+1] = src[offset+2]
			cursor += 2
		case 0b0111:
			dst[cursor] = src[offset+0]
			dst[cursor+1] = src[offset+1]
			dst[cursor+2] = src[offset+2]
			cursor += 3
		case 0b1000:
			dst[cursor] = src[offset+3]
			cursor += 1
		case 0b1001:
			dst[cursor] = src[offset+0]
			dst[cursor+1] = src[offset+3]
			cursor += 2
		case 0b1010:
			dst[cursor] = src[offset+1]
			dst[cursor+1] = src[offset+3]
			cursor += 2
		case 0b1011:
			dst[cursor] = src[offset+0]
			dst[cursor+1] = src[offset+1]
			dst[cursor+2] = src[offset+3]
			cursor += 3
		case 0b1100:
			dst[cursor] = src[offset+2]
			dst[cursor+1] = src[offset+3]
			cursor += 2
		case 0b1101:
			dst[cursor] = src[offset+0]
			dst[cursor+1] = src[offset+2]
			dst[cursor+2] = src[offset+3]
			cursor += 3
		case 0b1110:
			dst[cursor] = src[offset+1]
			dst[cursor+1] = src[offset+2]
			dst[cursor+2] = src[offset+3]
			cursor += 3
		case 0b1111:
			dst[cursor] = src[offset+0]
			dst[cursor+1] = src[offset+1]
			dst[cursor+2] = src[offset+2]
			dst[cursor+3] = src[offset+3]
			cursor += 4
		}

		offset += 4
	}

	return (*dst)[:cursor]
}
