// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"bytes"
	"encoding/json"
	"math"
	"math/rand"
	"strings"
	"testing"

	"github.com/klauspost/cpuid/v2"
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

func TestJSON(t *testing.T) {
	mp := Bitmap{}

	for i := 0; i < 1000; i++ {
		mp.Set(uint32(rand.Intn(10000)))
	}

	data, err := json.Marshal(mp)
	assert.NoError(t, err)

	newMp := Bitmap{}
	assert.NoError(t, json.Unmarshal(data, &newMp))
	assert.Equal(t, mp, newMp)

	assert.NoError(t, mp.UnmarshalJSON(nil))
	assert.Empty(t, mp)

	assert.Error(t, mp.UnmarshalJSON([]byte("\"notvalid")))
	assert.Error(t, mp.UnmarshalJSON([]byte("\"Z\"")))
}

func TestToHexadecimal(t *testing.T) {
	type Case struct {
		Input  uint64
		Pad    bool
		Output string
	}
	tests := []Case{{
		Input:  0,
		Pad:    false,
		Output: "0",
	}, {
		Input:  42,
		Pad:    false,
		Output: "2A",
	}, {
		Input:  math.MaxUint64,
		Pad:    false,
		Output: "FFFFFFFFFFFFFFFF",
	}, {
		Input:  15,
		Pad:    true,
		Output: "000000000000000F",
	},
	}

	for _, tc := range tests {
		sb := strings.Builder{}
		writeHexdecimal(&sb, tc.Input, tc.Pad)
		assert.Equal(t, tc.Output, sb.String())
	}

}

func TestFromHex(t *testing.T) {
	bm, err := fromHex("FFA001")
	assert.NoError(t, err)
	assert.Equal(t, Bitmap{0xFFA001}, bm)

	bm, err = fromHex("000000000000000000000000000000000001")
	assert.NoError(t, err)
	assert.Equal(t, Bitmap{1, 0, 0}, bm)

	_, err = fromHex("Not Valid")
	assert.Error(t, err)

	bm, err = fromHex("")
	assert.NoError(t, err)
	assert.Nil(t, bm)
}

func TestDimensionsOf(t *testing.T) {
	testCases := []struct {
		n        int
		m        int
		expected uint64
	}{
		{0, 0, 0},
		{10, 11, 0xb0000000a},
	}

	for _, tc := range testCases {
		d := dimensionsOf(tc.n, tc.m)
		assert.Equal(t, tc.expected, d)
	}
}

func TestPointersOf(t *testing.T) {
	testCases := []struct {
		inputOther Bitmap
		inputExtra []Bitmap
	}{
		{Bitmap{1}, []Bitmap{{2}}},
		{Bitmap{1}, []Bitmap{{2}, {3}}},
	}

	for _, tc := range testCases {
		ptr, max := pointersOf(tc.inputOther, tc.inputExtra)
		assert.NotNil(t, ptr)
		assert.NotZero(t, uintptr(ptr))
		assert.NotZero(t, max)
	}
}

func TestLevelOfWithEnabledFeatures(t *testing.T) {
	testCases := []struct {
		name       string
		featureIDs []cpuid.FeatureID
		expected   int
	}{
		{
			name:       "AVX-512F, AVX-512BW, and AVX-512DQ support",
			featureIDs: []cpuid.FeatureID{cpuid.AVX512F, cpuid.AVX512BW, cpuid.AVX512DQ},
			expected:   isAVX512,
		},
		{
			name:       "AVX2 and FMA3 support",
			featureIDs: []cpuid.FeatureID{cpuid.AVX2, cpuid.FMA3},
			expected:   isAccelerated,
		},
		{
			name:       "NEON support on ARM64",
			featureIDs: []cpuid.FeatureID{cpuid.ASIMD},
			expected:   isAccelerated,
		},
		{
			name:       "Unsupported feature combination",
			featureIDs: []cpuid.FeatureID{cpuid.AVX512F, cpuid.AESARM},
			expected:   isUnsupported,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cpu := cpuid.CPUInfo{}
			for _, feature := range tc.featureIDs {
				cpu.Enable(feature)
			}

			level := levelOf(cpu)
			assert.Equal(t, tc.expected, level, "expected to return %d, but got %d", tc.expected, level)
		})
	}
}
