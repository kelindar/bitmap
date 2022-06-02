// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkCodec(b *testing.B) {
	tmp := bytes.NewBuffer(nil)
	run(b, "write-to", func(index Bitmap) {
		tmp.Reset()
		index.WriteTo(tmp)
	})

	run(b, "read-from", func(index Bitmap) {
		ReadFrom(tmp)
	})
}

func TestSaveLoad(t *testing.T) {
	m := Bitmap{}
	for i := 0; i <= 5000; i += 10 {
		m.Set(uint32(i))
	}

	// Save the map
	enc := new(bytes.Buffer)
	cloned := m.Clone(nil)
	n, err := cloned.WriteTo(enc)
	assert.NoError(t, err)
	assert.Equal(t, int64(636), n)

	// Load the map back
	out, err := ReadFrom(enc)
	assert.NoError(t, err)
	assert.Equal(t, len(m), len(out))
	assert.Equal(t, m, out)
}

func TestFromBytes(t *testing.T) {
	m := Bitmap{}
	for i := 0; i <= 5000; i += 10 {
		m.Set(uint32(i))
	}

	out := FromBytes(m.ToBytes())
	assert.Equal(t, m, out)
}

func TestFromBytesNil(t *testing.T) {
	out := FromBytes(nil)
	assert.Nil(t, out)
}

func TestFromBytesInvalid(t *testing.T) {
	m := make([]byte, 10)
	for i := 1; i < 8; i++ {
		assert.Panics(t, func() {
			FromBytes(m[:i])
		})
	}
}

func TestToBytesNil(t *testing.T) {
	var m Bitmap
	out := m.ToBytes()
	assert.Nil(t, out)
}

func TestFromHex(t *testing.T) {
	bm, err := FromHex("FFA001")
	assert.NoError(t, err)
	assert.Equal(t, Bitmap{0xFFA001}, bm)

	bm, err = FromHex("000000000000000000000000000000000001")
	assert.NoError(t, err)
	assert.Equal(t, Bitmap{1, 0, 0}, bm)

	_, err = FromHex("Not Valid")
	assert.Error(t, err)

	bm, err = FromHex("")
	assert.NoError(t, err)
	assert.Nil(t, bm)
}
