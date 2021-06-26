package bitmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkRange/range-8         	   12834	     92889 ns/op	       0 B/op	       0 allocs/op
BenchmarkRange/filter-8        	   18354	     64827 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkRange(b *testing.B) {
	run(b, "range", func(index Bitmap) {
		index.Range(func(x uint32) bool {
			return true
		})
	})

	run(b, "filter", func(index Bitmap) {
		index.Filter(func(x uint32) bool {
			return x%2 == 0
		})
	})
}

func TestFilter(t *testing.T) {
	a := make(Bitmap, 4)
	a.Ones()
	assert.Equal(t, 256, a.Count())

	{ // Filter out odd
		count := 0
		a.Filter(func(x uint32) bool {
			count++
			return x%2 == 0
		})
		assert.Equal(t, 256, count)
		assert.Equal(t, 128, a.Count())
	}

	{
		// Filter out even
		count := 0
		a.Filter(func(x uint32) bool {
			count++
			return x%2 == 1
		})
		assert.Equal(t, 128, count)
		assert.Equal(t, 0, a.Count())
	}
}

func TestRange(t *testing.T) {
	a := Bitmap{}
	for i := uint32(0); i < 100; i += 2 {
		a.Set(i)
	}

	count := 0
	a.Range(func(x uint32) bool {
		count++
		return true
	})
	assert.Equal(t, 50, count)
	assert.Equal(t, 50, a.Count())
}

func TestRangeCases(t *testing.T) {
	a := Bitmap{}
	for i := uint32(0); i < 100; i++ {
		a.Set(i)
	}

	for i := 0; i < 100; i++ {
		count := 0
		a.Range(func(x uint32) bool {
			if count == i {
				return false
			}

			count++
			return true
		})
		assert.Equal(t, i, count)
	}
}

func TestRangeIndex(t *testing.T) {
	a := make(Bitmap, 2)
	a.Ones()

	triangular := 0
	a.Range(func(x uint32) bool {
		triangular += int(x)
		return true
	})
	assert.Equal(t, 8128, triangular)
}

// run runs a benchmark
func run(b *testing.B, name string, f func(index Bitmap)) {
	count := 100000
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
