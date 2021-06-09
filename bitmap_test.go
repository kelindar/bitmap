// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"testing"

	"github.com/kelindar/bitmap/simd"
	"github.com/stretchr/testify/assert"
)

// BenchmarkBitmap/set-8         	608127316	         1.979 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/remove-8      	775627708	         1.562 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/contains-8    	907577592	         1.299 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/clear-8       	231583378	         5.163 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/ones-8        	39476930	        29.77 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/first-zero-8  	23612611	        50.82 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/min-8         	415250632	         2.916 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/max-8         	683142546	         1.763 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/count-8       	33334074	        34.88 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/clone-8       	100000000	        11.46 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/simd-and-8    	74337927	        15.47 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/simd-andnot-8 	80220294	        14.92 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/simd-or-8     	81321524	        14.81 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/simd-xor-8    	80181888	        14.81 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/and-8         	29650201	        41.68 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/andnot-8      	26496499	        51.72 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/or-8          	20629934	        50.83 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/xor-8         	23786632	        51.46 ns/op	       0 B/op	       0 allocs/op
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

	// With AVX (should be default)
	simd.Supported = true
	run(b, "simd-and", func(index Bitmap) {
		index.And(other)
	})

	run(b, "simd-andnot", func(index Bitmap) {
		index.AndNot(other)
	})

	run(b, "simd-or", func(index Bitmap) {
		index.AndNot(other)
	})

	run(b, "simd-xor", func(index Bitmap) {
		index.AndNot(other)
	})

	// Disable AVX
	simd.Supported = false
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
		a := Bitmap{0b0011, 0b1011, 0b1100, 0b0000, 0b0011, 0b1011, 0b1100, 0b0000, 0b0011}
		a.And(Bitmap{0b0101, 0b1101, 0b1010, 0b1111, 0b0101, 0b1101, 0b1010, 0b1111, 0b0101})
		assert.Equal(t, 0b0001, int(a[0]))
		assert.Equal(t, 0b1001, int(a[1]))
		assert.Equal(t, 0b1000, int(a[2]))
		assert.Equal(t, 0b0000, int(a[3]))
		assert.Equal(t, 0b0001, int(a[4]))
		assert.Equal(t, 0b1001, int(a[5]))
		assert.Equal(t, 0b1000, int(a[6]))
		assert.Equal(t, 0b0000, int(a[7]))
		assert.Equal(t, 0b0001, int(a[8]))
	}
	{ // AND NOT
		a := Bitmap{0b0011, 0, 0, 0}
		a.AndNot(Bitmap{0b0101, 0, 0, 0})
		assert.Equal(t, 0b0010, int(a[0]))
	}
	{ // OR
		a := Bitmap{0b0011, 0, 0, 0}
		a.Or(Bitmap{0b0101, 0, 0, 0})
		assert.Equal(t, 0b0111, int(a[0]))
	}
	{ // XOR
		a := Bitmap{0b0011, 0, 0, 0}
		a.Xor(Bitmap{0b0101, 0, 0, 0})
		assert.Equal(t, 0b0110, int(a[0]))
	}
}

func TestTruthTables_NoSIMD(t *testing.T) {
	simd.Supported = false
	TestTruthTables(t)
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
