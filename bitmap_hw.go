// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

//go:build !noasm && (amd64 || arm64)

package bitmap

import "unsafe"

func dimensionsOf(n, m int) uint64 {
	return uint64(n) | (uint64(m) << 32)
}

func convertToPointerSlice(other Bitmap, extra []Bitmap) (unsafe.Pointer, int) {
	out := make([]unsafe.Pointer, len(extra)+1)
	out[0] = unsafe.Pointer(&other[0])
	max := 0

	for i := range extra {
		out[i+1] = unsafe.Pointer(&extra[i][0])
		if len(extra[i]) > max {
			max = len(extra[i])
		}
	}

	return unsafe.Pointer(&out[0]), max
}

// And computes the intersection between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) And(other Bitmap, extra ...Bitmap) {
	max := minlen(*dst, other, extra)
	dst.shrink(max)

	switch hardware {
	case isAccelerated:
		switch len(extra) {
		case 0:
			_and(unsafe.Pointer(&(*dst)[0]), unsafe.Pointer(&other[0]), uint64(max))
		default:
			vx, _ := convertToPointerSlice(other, extra)
			_and_many(unsafe.Pointer(&(*dst)[0]), vx, dimensionsOf(max, len(extra)+1))
		}
	case isAVX512:
		switch len(extra) {
		case 0:
			_and_avx512(unsafe.Pointer(&(*dst)[0]), unsafe.Pointer(&other[0]), uint64(max))
		default:
			vx, _ := convertToPointerSlice(other, extra)
			_and_many_avx512(unsafe.Pointer(&(*dst)[0]), vx, dimensionsOf(max, len(extra)+1))
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

	switch hardware {
	case isAccelerated:
		switch len(extra) {
		case 0:
			_andn(unsafe.Pointer(&(*dst)[0]), unsafe.Pointer(&other[0]), uint64(max))
		default:
			vx, _ := convertToPointerSlice(other, extra)
			_andn_many(unsafe.Pointer(&(*dst)[0]), vx, dimensionsOf(max, len(extra)+1))
		}
	case isAVX512:
		switch len(extra) {
		case 0:
			_andn_avx512(unsafe.Pointer(&(*dst)[0]), unsafe.Pointer(&other[0]), uint64(max))
		default:
			vx, _ := convertToPointerSlice(other, extra)
			_andn_many_avx512(unsafe.Pointer(&(*dst)[0]), vx, dimensionsOf(max, len(extra)+1))
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

	switch hardware {
	case isAccelerated:
		switch len(extra) {
		case 0:
			_or(unsafe.Pointer(&(*dst)[0]), unsafe.Pointer(&other[0]), uint64(len(other)))
		default:
			vx, max := convertToPointerSlice(other, extra)
			_or_many(unsafe.Pointer(&(*dst)[0]), vx, dimensionsOf(max, len(extra)+1))
		}
	case isAVX512:
		switch len(extra) {
		case 0:
			_or_avx512(unsafe.Pointer(&(*dst)[0]), unsafe.Pointer(&other[0]), uint64(len(other)))
		default:
			vx, max := convertToPointerSlice(other, extra)
			_or_many_avx512(unsafe.Pointer(&(*dst)[0]), vx, dimensionsOf(max, len(extra)+1))
		}
	default:
		or(*dst, other, extra)
	}
}

// Xor computes the symmetric difference between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) Xor(other Bitmap, extra ...Bitmap) {
	max := maxlen(*dst, other, extra)
	dst.grow(max - 1)

	switch hardware {
	case isAccelerated:
		switch len(extra) {
		case 0:
			_xor(unsafe.Pointer(&(*dst)[0]), unsafe.Pointer(&other[0]), uint64(len(other)))
		default:
			vx, max := convertToPointerSlice(other, extra)
			_xor_many(unsafe.Pointer(&(*dst)[0]), vx, dimensionsOf(max, len(extra)+1))
		}
	case isAVX512:
		switch len(extra) {
		case 0:
			_xor_avx512(unsafe.Pointer(&(*dst)[0]), unsafe.Pointer(&other[0]), uint64(len(other)))
		default:
			vx, max := convertToPointerSlice(other, extra)
			_xor_many_avx512(unsafe.Pointer(&(*dst)[0]), vx, dimensionsOf(max, len(extra)+1))
		}
	default:
		xor(*dst, other, extra)
	}
}

// Count returns the number of elements in this bitmap
func (dst Bitmap) Count() int {
	if len(dst) == 0 {
		return 0
	}

	switch hardware {
	case isAccelerated:
		var res uint64
		_count(unsafe.Pointer(&dst[0]), uint64(len(dst)), unsafe.Pointer(&res))
		return int(res)
	default:
		return count(dst)
	}
}
