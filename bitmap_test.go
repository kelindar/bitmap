// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchmarkBitmap/set-8         	607986201	         2.004 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/remove-8      	759495112	         1.564 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/contains-8    	916055708	         1.306 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/and-8         	30769704	        39.00 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/andnot-8      	23518297	        51.11 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/or-8          	24000480	        50.92 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/xor-8         	23529734	        51.24 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/clear-8       	206332054	         5.845 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/ones-8        	39965895	        29.47 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/first-zero-8  	30001574	        40.15 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/min-8         	375470745	         3.295 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/max-8         	681821280	         1.740 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/count-8       	32433483	        35.64 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/clone-8       	100000000	        11.77 ns/op	       0 B/op	       0 allocs/op
func BenchmarkBitmap(b *testing.B) {
	other := make(Bitmap, 100)
	other.Set(5000)

	run(b, "set", func(index Bitmap) {
		index.Set(5000)
	})

	run(b, "remove", func(index Bitmap) {
		index.Remove(5000)
	})

	run(b, "contains", func(index Bitmap) {
		index.Contains(5000)
	})

	run(b, "and", func(index Bitmap) {
		index.And(other)
	})

	run(b, "andnot", func(index Bitmap) {
		index.AndNot(other)
	})

	run(b, "or", func(index Bitmap) {
		index.AndNot(other)
	})

	run(b, "xor", func(index Bitmap) {
		index.AndNot(other)
	})

	run(b, "clear", func(index Bitmap) {
		index.Clear()
	})

	run(b, "ones", func(index Bitmap) {
		index.Ones()
	})

	run(b, "first-zero", func(index Bitmap) {
		index.FirstZero()
	})

	run(b, "min", func(index Bitmap) {
		index.Min()
	})

	run(b, "max", func(index Bitmap) {
		index.Max()
	})

	run(b, "count", func(index Bitmap) {
		index.Count()
	})

	var into Bitmap
	run(b, "clone", func(index Bitmap) {
		index.Clone(&into)
	})

}

func TestSetRemove(t *testing.T) {
	index := Bitmap{}
	for i := uint32(100); i < 200; i++ {
		index.Set(i)
		assert.True(t, index.Contains(i))
	}

	for i := uint32(150); i < 180; i++ {
		index.Remove(i)
		assert.False(t, index.Contains(i))
	}
}

func TestClear(t *testing.T) {
	index := Bitmap{}
	for i := uint32(100); i < 200; i++ {
		index.Set(i)
		assert.True(t, index.Contains(i))
	}

	index.Clear()
	for i := uint32(100); i < 200; i++ {
		assert.False(t, index.Contains(i))
	}
}

func TestTruthTables(t *testing.T) {
	{ // AND
		a := Bitmap{0b0011}
		a.And(Bitmap{0b0101})
		assert.Equal(t, 0b0001, int(a[0]))
	}
	{ // AND NOT
		a := Bitmap{0b0011}
		a.AndNot(Bitmap{0b0101})
		assert.Equal(t, 0b0010, int(a[0]))
	}
	{ // OR
		a := Bitmap{0b0011}
		a.Or(Bitmap{0b0101})
		assert.Equal(t, 0b0111, int(a[0]))
	}
	{ // XOR
		a := Bitmap{0b0011}
		a.Xor(Bitmap{0b0101})
		assert.Equal(t, 0b0110, int(a[0]))
	}
}

func TestAnd(t *testing.T) {
	a, b := Bitmap{}, Bitmap{}
	for i := uint32(0); i < 100; i += 2 {
		a.Set(i)
		b.Set(i)
	}

	a.And(b)
	assert.False(t, a.Contains(1))
	for i := uint32(0); i < 100; i += 2 {
		assert.True(t, a.Contains(i))
	}
}

func TestAndNot(t *testing.T) {
	a, b := Bitmap{}, Bitmap{}
	for i := uint32(0); i < 100; i += 2 {
		a.Set(i)
		b.Set(i)
	}

	a.AndNot(b)
	assert.False(t, a.Contains(1))
	for i := uint32(0); i < 100; i += 2 {
		assert.False(t, a.Contains(i))
	}
}

func TestOr(t *testing.T) {
	a, b := Bitmap{}, Bitmap{}
	for i := uint32(0); i < 100; i += 2 {
		b.Set(i)
	}

	a.Or(b)
	assert.False(t, a.Contains(1))
	for i := uint32(0); i < 100; i += 2 {
		assert.True(t, a.Contains(i))
	}
}

func TestXor(t *testing.T) {
	a, b := Bitmap{}, Bitmap{}
	for i := uint32(0); i < 100; i += 2 {
		b.Set(i)
	}

	a.Xor(b)
	assert.False(t, a.Contains(1))
	for i := uint32(0); i < 100; i += 2 {
		assert.True(t, a.Contains(i))
	}
}

func TestMin(t *testing.T) {
	{
		a := Bitmap{0x0, 0x0, 0xffffffffffffff00}
		v, ok := a.Min()
		assert.True(t, ok)
		assert.Equal(t, 64+64+8, int(v))
		assert.False(t, a.Contains(v-1))
		assert.True(t, a.Contains(v))
	}

	{
		a := Bitmap{0x0, 0x0}
		v, ok := a.Min()
		assert.False(t, ok)
		assert.Equal(t, 0, int(v))
	}
}

func TestMax(t *testing.T) {
	{
		a := Bitmap{0x0, 0x0, 0x00000000000000f0}
		v, ok := a.Max()
		assert.True(t, ok)
		assert.Equal(t, 64+64+7, int(v))

		assert.False(t, a.Contains(v-4))
		assert.True(t, a.Contains(v-3))
		assert.True(t, a.Contains(v-2))
		assert.True(t, a.Contains(v-1))
		assert.True(t, a.Contains(v))
		assert.False(t, a.Contains(v+1))
		assert.False(t, a.Contains(v+2))
	}

	{
		a := Bitmap{0x0, 0x0}
		v, ok := a.Max()
		assert.False(t, ok)
		assert.Equal(t, 0, int(v))
	}
}

func TestFirstZero(t *testing.T) {
	{
		a := Bitmap{0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffff0f}
		v, ok := a.FirstZero()
		assert.True(t, ok)
		assert.Equal(t, 64+64+4, int(v))
		assert.False(t, a.Contains(v))
	}

	{
		a := Bitmap{0xffffffffffffffff, 0xffffffffffffffff}
		v, ok := a.FirstZero()
		assert.False(t, ok)
		assert.Equal(t, 0, int(v))
	}
}
