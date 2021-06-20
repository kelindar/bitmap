// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkBitmap/set-12         	271467813	         4.335 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/remove-12      	677338836	         1.779 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/contains-12    	806530189	         1.497 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/clear-12       	100000000	        10.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/ones-12        	36235255	        33.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/first-zero-12  	26520604	        45.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/min-12         	369584089	         3.258 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/max-12         	608877200	         1.987 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/count-12       	29074789	        41.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/count-to-12    	24887761	        48.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/clone-12       	63780740	        18.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/and-12         	59301794	        20.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/andnot-12      	53553556	        22.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/or-12          	52102473	        22.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/xor-12         	51208834	        22.55 ns/op	       0 B/op	       0 allocs/op
*/
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

	run(b, "count-to", func(index Bitmap) {
		index.CountTo(5001)
	})

	var into Bitmap
	run(b, "clone", func(index Bitmap) {
		index.Clone(&into)
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
	for i := uint32(0); i < 500; i++ {
		index.Set(i)
		assert.True(t, index.Contains(i))
	}

	index.Clear()
	index.Set(500)
	for i := uint32(0); i < 500; i++ {
		assert.False(t, index.Contains(i), i)
	}
	assert.True(t, index.Contains(500))
}

func testTruthTables(t *testing.T) {
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
	avx2 = false
	testTruthTables(t)
}

func TestTruthTables_SIMD(t *testing.T) {
	avx2 = true
	testTruthTables(t)
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

func TestCount(t *testing.T) {
	a := Bitmap{}
	assert.Equal(t, 0, a.Count())
	assert.Equal(t, 0, a.CountTo(math.MaxUint32))

	b := Bitmap{}
	b.Set(1)
	b.Set(2)
	b.Set(5)
	b.Set(6)
	b.Set(101)
	b.Set(102)
	b.Set(105)
	b.Set(106)

	assert.Equal(t, 8, b.Count())
	assert.Equal(t, 1, b.CountTo(2))
	assert.Equal(t, 2, b.CountTo(4))
	assert.Equal(t, 4, b.CountTo(100))
	assert.Equal(t, 4, b.CountTo(101))
	assert.Equal(t, 5, b.CountTo(102))
	assert.Equal(t, 8, b.CountTo(math.MaxUint32))
}

func TestGrow(t *testing.T) {
	bitmap := make(Bitmap, 1, 5)
	bitmap[0] = 42

	assert.Equal(t, 1, len(bitmap))
	assert.Equal(t, 5, cap(bitmap))
	assert.Equal(t, Bitmap{42}, bitmap)

	bitmap.grow(0)
	assert.Equal(t, 1, len(bitmap))
	assert.Equal(t, 5, cap(bitmap))
	assert.Equal(t, Bitmap{42}, bitmap)

	bitmap.grow(4)
	assert.Equal(t, 5, len(bitmap))
	assert.Equal(t, 5, cap(bitmap))
	assert.Equal(t, Bitmap{42, 0, 0, 0, 0}, bitmap)

	bitmap.grow(5)
	assert.Equal(t, 6, len(bitmap))
	assert.Equal(t, Bitmap{42, 0, 0, 0, 0, 0}, bitmap)
	bitmap.Grow(6)
}
