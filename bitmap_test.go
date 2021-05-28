// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package bitmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchmarkBitmap/set-8         	600116121	         1.987 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/remove-8      	828642868	         1.534 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/contains-8    	910330890	         1.305 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/and-8         	29602585	        38.39 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/andnot-8      	29372960	        49.75 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/or-8          	23535918	        50.01 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/xor-8         	23752498	        50.69 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/clear-8       	203721861	         5.910 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/ones-8        	41331565	        29.70 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitmap/clone-8       	 8291484	       137.6 ns/op	     896 B/op	       1 allocs/op
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

	run(b, "clone", func(index Bitmap) {
		index.Clone(nil)
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

// run runs a benchmark
func run(b *testing.B, name string, f func(index Bitmap)) {
	b.Run(name, func(b *testing.B) {
		index := make(Bitmap, 100)
		for i := 0; i < len(index); i++ {
			index[i] = 0xffffffffffffffff
		}

		b.ReportAllocs()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			f(index)
		}
	})
}
