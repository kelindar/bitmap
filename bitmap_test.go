// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"math"
	"testing"

	"github.com/klauspost/cpuid/v2"
	"github.com/stretchr/testify/assert"
)

/*
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkBitmap/set-8         	289575987	         4.122 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/remove-8      	761542122	         1.561 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/contains-8    	908733610	         1.312 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/clear-8       	131705713	         9.073 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/ones-8        	40716474	        30.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/first-zero-8  	23821670	        50.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/min-8         	420276520	         2.860 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/max-8         	777923565	         1.542 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/count-8       	42097286	        27.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/count-to-8    	27890955	        42.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/clone-8       	75004218	        15.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/and-8         	70589480	        17.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/andnot-8      	70588650	        17.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/or-8          	66248564	        17.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitmap/xor-8         	70587819	        17.46 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkBitmap(b *testing.B) {
	other := make(Bitmap, 100000/64)
	other.Set(100000)

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

	run(b, "min", func(index Bitmap) {
		index.Min()
	})

	run(b, "max", func(index Bitmap) {
		index.Max()
	})

	run(b, "min-zero", func(index Bitmap) {
		index.MinZero()
	})

	run(b, "max-zero", func(index Bitmap) {
		index.MaxZero()
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
		a.AndNot(Bitmap{0b0101})
		assert.Equal(t, 0b0010, int(a[0]))
	}
	{ // OR
		a := Bitmap{0b0011, 0, 0, 0}
		a.Or(Bitmap{0b0101})
		assert.Equal(t, 0b0111, int(a[0]))
	}
	{ // XOR
		a := Bitmap{0b0011, 0, 0, 0}
		a.Xor(Bitmap{0b0101})
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

func TestMinZero(t *testing.T) {
	{
		a := Bitmap{0xffffffffffffffff, 0xffffffffffffffff, 0xf0ffffffffffff0f}
		v, ok := a.MinZero()
		assert.True(t, ok)
		assert.Equal(t, 64+64+4, int(v))
		assert.False(t, a.Contains(v))
	}

	{
		a := Bitmap{0xffffffffffffffff, 0xffffffffffffffff}
		v, ok := a.MinZero()
		assert.False(t, ok)
		assert.Equal(t, 0, int(v))
	}
}

func TestMaxZero(t *testing.T) {
	{
		a := Bitmap{0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffff0f}
		v, ok := a.MaxZero()
		assert.True(t, ok)
		assert.Equal(t, 64+64+7, int(v))
		assert.False(t, a.Contains(v))
	}

	{
		a := Bitmap{0xffffffffffffffff, 0xffffffffffffffff}
		v, ok := a.MaxZero()
		assert.False(t, ok)
		assert.Equal(t, 0, int(v))
	}
}

func TestCount(t *testing.T) {
	testCount(t, true)
	testCount(t, false)
}

func testCount(t *testing.T, x64 bool) {
	popc = x64
	defer func() {
		popc = cpuid.CPU.Supports(cpuid.POPCNT)
	}()

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
