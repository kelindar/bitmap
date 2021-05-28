package bitmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchmarkRange/range-8         	  944895	      1246 ns/op	       0 B/op	       0 allocs/op
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
