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
func ReadFrom(r io.Reader) (Bitmap, error) {
	var output Bitmap
	_, err := output.ReadFrom(r)
	return output, err
}

// WriteTo writes the bitmap to a specified writer.
func (dst *Bitmap) WriteTo(w io.Writer) (int64, error) {
	buffer := dst.ToBytes()

	// Write the header into the stream
	var header [4]byte
	binary.BigEndian.PutUint32(header[:4], uint32(len(buffer)))
	n1, err := w.Write(header[:4])
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

// ReadFrom reads data from r until EOF or error. The return value n is the number of
// bytes read. Any error except EOF encountered during the read is also returned.
func (dst *Bitmap) ReadFrom(r io.Reader) (int64, error) {
	var header [4]byte
	if n, err := io.ReadFull(r, header[:]); err != nil {
		return int64(n), err
	}

	// If bitmap is too small, create one of the required size
	if size := int(binary.BigEndian.Uint32(header[:4])) / 8; size > len(*dst) {
		*dst = make(Bitmap, size)
	}

	// Read into the buffer
	buffer := dst.ToBytes()
	n, err := io.ReadFull(r, buffer)
	return int64(n + 4), err
}

// Clone clones the bitmap. If a destination bitmap is provided, the bitmap will be
// cloned inside, otherwise a new Bitmap will be allocated and returned
func (dst Bitmap) Clone(into *Bitmap) Bitmap {
	if into == nil {
		newm := make(Bitmap, len(dst))
		into = &newm
	}

	max := maxlen(*into, dst, nil)
	into.grow(max - 1)

	copy(*into, dst)
	return (*into)[:len(dst)]
}

// Clear clears the bitmap and resizes it to zero.
func (dst *Bitmap) Clear() {
	for i := range *dst {
		(*dst)[i] = 0
	}
	*dst = (*dst)[:0]
}
