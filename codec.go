// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"encoding/binary"
	"io"
	"reflect"
	"unsafe"
)

// FromBytes reads a bitmap from a byte buffer without copying the buffer.
func FromBytes(buffer []byte) (out Bitmap) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&out))
	hdr.Len = len(buffer) >> 3
	hdr.Cap = hdr.Len
	hdr.Data = uintptr(unsafe.Pointer(&(buffer)[0]))
	return out
}

// ToBytes converts the bitmap to binary representation without copying the underlying
// data. The output buffer should not be modified, since it would also change the bitmap.
func (dst *Bitmap) ToBytes() (out []byte) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&out))
	hdr.Len = len(*dst) * 8
	hdr.Cap = hdr.Len
	hdr.Data = uintptr(unsafe.Pointer(&(*dst)[0]))
	return out
}

// ReadFrom reads the bitmap from the reader.
func ReadFrom(rdr io.Reader) (Bitmap, error) {
	header := make([]byte, 4)
	if _, err := io.ReadFull(rdr, header); err != nil {
		return Bitmap{}, err
	}

	buffer := make([]byte, binary.BigEndian.Uint32(header[:4]))
	if _, err := io.ReadFull(rdr, buffer); err != nil {
		return Bitmap{}, err
	}

	return FromBytes(buffer), nil
}

// WriteTo writes the bitmap to a specified writer.
func (dst *Bitmap) WriteTo(w io.Writer) (int64, error) {
	buffer := dst.ToBytes()
	header := make([]byte, 4)

	// Write the header into the stream
	binary.BigEndian.PutUint32(header[0:4], uint32(len(buffer)))
	n1, err := w.Write(header)
	if err != nil {
		return int64(n1), err
	}

	// Write the buffer into the stream
	n2, err := w.Write(buffer)
	if err != nil {
		return int64(n2), err
	}

	return int64(n1 + n2), err
}

// Clone clones the bitmap. If a destination bitmap is provided, the bitmap will be
// cloned inside, otherwise a new Bitmap will be allocated and returned
func (dst Bitmap) Clone(into *Bitmap) Bitmap {
	if into == nil {
		newm := make(Bitmap, len(dst))
		into = &newm
	}

	if into.balance(dst); len(*into) < len(dst) {
		return nil // Elliminate bounds check
	}

	copy(*into, dst)
	*into = (*into)[:len(dst)]
	return *into
}

// Clear clears the bitmap and resizes it to zero.
func (dst *Bitmap) Clear() {
	if size := len(*dst); size > 0 {
		ptr := unsafe.Pointer(&(*dst)[0])
		memclrNoHeapPointers(ptr, uintptr(size))
		*dst = (*dst)[:0]
	}
}

//go:linkname memclrNoHeapPointers runtime.memclrNoHeapPointers
func memclrNoHeapPointers(p unsafe.Pointer, n uintptr)
