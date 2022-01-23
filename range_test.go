package bitmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkRange/range-8         	   15823	     75607 ns/op	       0 B/op	       0 allocs/op
BenchmarkRange/filter-8        	   20366	     59646 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkRange(b *testing.B) {
	var i uint32
	run(b, "range", func(index Bitmap) {
		index.Range(func(x uint32) {
			i = x
			return
		})
	})

	run(b, "filter", func(index Bitmap) {
		index.Filter(func(x uint32) bool {
			return x%2 == 0
		})
	})

	_ = i
}

func TestFilter(t *testing.T) {
	a := make(Bitmap, 4)
	a.Ones()
	assert.Equal(t, 256, a.Count())

	// Filter out odd
	a.Filter(func(x uint32) bool {
		return x%2 == 0
	})
	assert.Equal(t, 128, a.Count())

	// Filter out even
	a.Filter(func(x uint32) bool {
		assert.Equal(t, 0, int(x%2)) // Must be odd
		return x%2 == 1
	})
	assert.Equal(t, 0, a.Count())

	// Filter cases
	for i := 0; i < 512; i++ {
		b := Bitmap{uint64(i)}
		c1 := b.Count()
		c2 := 0
		b.Filter(func(x uint32) bool {
			c2++
			return true
		})

		// We must have the minimum number of function calls
		assert.Equal(t, c1, c2)
		assert.Equal(t, uint64(i), b[0])
	}
}

func TestRangeCases(t *testing.T) {
	for i := 0; i < 512; i++ {
		b := Bitmap{uint64(i)}
		c1 := b.Count()
		c2 := 0
		b.Range(func(x uint32) {
			c2++
			return
		})

		// We must have the minimum number of function calls
		assert.Equal(t, c1, c2)
		assert.Equal(t, uint64(i), b[0])
	}
}

func TestRangeIndex(t *testing.T) {
	a := make(Bitmap, 2)
	a.Ones()

	triangular := 0
	a.Range(func(x uint32) {
		triangular += int(x)
		return
	})
	assert.Equal(t, 8128, triangular)
}

// run runs a benchmark
func run(b *testing.B, name string, f func(index Bitmap)) {
	count := 1000000
	b.Run(name, func(b *testing.B) {
		index := make(Bitmap, count/64)
		index.Grow(uint32(count))
		for i := 0; i < len(index); i++ {
			index[i] = 0xf0f0f0f0f0f0f0f0
		}

		b.ReportAllocs()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			f(index)
		}
	})
}
