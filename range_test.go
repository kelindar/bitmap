package bitmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchmarkRange/range-8         	   80383	     14849 ns/op	       0 B/op	       0 allocs/op
// BenchmarkRange/first-zero-8    	22536818	        51.76 ns/op	       0 B/op	       0 allocs/op
func BenchmarkRange(b *testing.B) {
	run(b, "range", func(index Bitmap) {
		index.Range(func(x uint32) bool {
			return true
		})
	})

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
