// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// FromBytes reads a bitmap from a byte buffer without copying the buffer.
func FromBytes(buffer []byte) (out Bitmap) {
	switch {
	case len(buffer) == 0:
		return nil
	case len(buffer)%8 != 0:
		panic(fmt.Sprintf("bitmap: buffer length expected to be multiple of 8, was %d", len(buffer)))
	}

	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&out))
	hdr.Len = len(buffer) >> 3
	hdr.Cap = hdr.Len
	hdr.Data = uintptr(unsafe.Pointer(&(buffer)[0]))
	return out
}

// ToBytes converts the bitmap to binary representation without copying the underlying
// data. The output buffer should not be modified, since it would also change the bitmap.
func (dst *Bitmap) ToBytes() (out []byte) {
	if len(*dst) == 0 {
		return nil
	}

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

// MarshalJSON returns encoded string representation for the bitmap
func (dst Bitmap) MarshalJSON() ([]byte, error) {
	var sb strings.Builder
	for i := len(dst) - 1; i >= 0; i-- {
		// convert each uint64 into 16 * 4-bit hexadecimal character
		writeHexdecimal(&sb, dst[i], true)
	}

	return json.Marshal(sb.String())
}

// writeHexdecimal write the hexdecimal representation for given value in buffer
func writeHexdecimal(sb *strings.Builder, value uint64, pad bool) {
	maxLen := 16 // 64 bits / 4

	hexadecimal := strings.ToUpper(strconv.FormatUint(value, 16))
	hexaLen := len(hexadecimal)

	if !pad || hexaLen == maxLen {
		sb.WriteString(hexadecimal)
		return
	}

	// Add padding
	for i := hexaLen; i < maxLen; i++ {
		sb.WriteString("0")
	}

	sb.WriteString(hexadecimal)
}

// UnmarshalJSON decodes the received bytes and loads it to bitmap object
func (dst *Bitmap) UnmarshalJSON(data []byte) (err error) {
	var str string
	if data == nil {
		*dst = make(Bitmap, 0)
		return
	}

	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	mp, err := fromHex(str)
	if err != nil {
		return err
	}

	*dst = mp
	return nil

}

// fromHex reads a hexadecimal string and converts it to bitmap, character at index 0 is the most significant
func fromHex(hexString string) (Bitmap, error) {
	bytes, err := hex.DecodeString(hexString)

	switch {
	case err != nil:
		return nil, err
	case len(bytes) == 0:
		return nil, nil
	}

	// reverse bytes to maintain bytes significance order (least significant = hexString tail = list head)
	for l, r := 0, len(bytes)-1; l < r; l, r = l+1, r-1 {
		bytes[l], bytes[r] = bytes[r], bytes[l]
	}

	for len(bytes)%8 != 0 {
		bytes = append(bytes, 0)
	}
	return FromBytes(bytes), nil
}
