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

func TestAnd(t *testing.T) {
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

// run runs a benchmark
func run(b *testing.B, name string, f func(index Bitmap)) {
	b.Run(name, func(b *testing.B) {
		index := make(Bitmap, 100)
		b.ReportAllocs()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			f(index)
		}
	})
}
