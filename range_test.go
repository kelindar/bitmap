package bitmap

import (
	"testing"

	"github.com/kelindar/simd"
	"github.com/stretchr/testify/assert"
)

/*
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkRange/range-8         	    1756	    678070 ns/op	       0 B/op	       0 allocs/op
BenchmarkRange/filter-8        	    2104	    539290 ns/op	       0 B/op	       0 allocs/op
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

/*
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkAggregate/sum-8         	    1996	    645131 ns/op	       0 B/op	       0 allocs/op
BenchmarkAggregate/sum-full-8    	    7982	    155056 ns/op	       0 B/op	       0 allocs/op
BenchmarkAggregate/min-8         	    1280	    874967 ns/op	       0 B/op	       0 allocs/op
BenchmarkAggregate/min-full-8    	    2974	    427653 ns/op	       0 B/op	       0 allocs/op
BenchmarkAggregate/max-8         	    1392	    877381 ns/op	       0 B/op	       0 allocs/op
BenchmarkAggregate/max-full-8    	    2880	    420351 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkAggregate(b *testing.B) {
	target := make([]float32, 1000100)
	run(b, "sum", func(index Bitmap) {
		Sum(target, index)
	})

	runFull(b, "sum-full", func(index Bitmap) {
		Sum(target, index)
	})

	run(b, "min", func(index Bitmap) {
		Min(target, index)
	})

	runFull(b, "min-full", func(index Bitmap) {
		Min(target, index)
	})

	run(b, "max", func(index Bitmap) {
		Max(target, index)
	})

	runFull(b, "max-full", func(index Bitmap) {
		Max(target, index)
	})
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

// ----------------------------- Aggregation -----------------------------

func TestAggSum(t *testing.T) {
	{ // Empty Bitmap
		arr, index := makeAggregateInput(0x0)
		assert.Equal(t, sumIf(arr, index), Sum(arr, index))
	}

	{ // Partial Bitmap
		arr, index := makeAggregateInput(0x0123456789abcdef)
		assert.Equal(t, sumIf(arr, index), Sum(arr, index))
	}

	{ // Full Bitmap
		arr, index := makeAggregateInput(0xffffffffffffffff)
		assert.Equal(t, sumIf(arr, index), Sum(arr, index))
	}

	{ // Nil Bitmap
		arr, _ := makeAggregateInput(0x0)
		assert.Equal(t, sumIf(arr, nil), Sum(arr, nil))
	}

	{ // Nil Array
		_, index := makeAggregateInput(0x0)
		assert.Equal(t, sumIf([]int{}, index), Sum([]int{}, index))
	}
}

func TestAggMin(t *testing.T) {
	{ // Empty Bitmap
		arr, index := makeAggregateInput(0x0)
		expect, ok1 := minIf(arr, index)
		result, ok2 := Min(arr, index)
		assert.Equal(t, expect, result)
		assert.Equal(t, ok1, ok2)
	}

	{ // Partial Bitmap
		arr, index := makeAggregateInput(0x0123456789abcdef)
		expect, ok1 := minIf(arr, index)
		result, ok2 := Min(arr, index)
		assert.Equal(t, expect, result)
		assert.Equal(t, ok1, ok2)
	}

	{ // Full Bitmap
		arr, index := makeAggregateInput(0xffffffffffffffff)
		expect, ok1 := minIf(arr, index)
		result, ok2 := Min(arr, index)
		assert.Equal(t, expect, result)
		assert.Equal(t, ok1, ok2)
	}

	{ // Nil Bitmap
		arr, _ := makeAggregateInput(0x0)
		expect, ok1 := minIf(arr, nil)
		result, ok2 := Min(arr, nil)
		assert.Equal(t, expect, result)
		assert.Equal(t, ok1, ok2)
	}

	{ // Nil Array
		_, index := makeAggregateInput(0x0)
		expect, ok1 := minIf([]int{}, index)
		result, ok2 := Min([]int{}, index)
		assert.Equal(t, expect, result)
		assert.Equal(t, ok1, ok2)
	}
}

func TestAggMax(t *testing.T) {
	{ // Empty Bitmap
		arr, index := makeAggregateInput(0x0)
		expect, ok1 := maxIf(arr, index)
		result, ok2 := Max(arr, index)
		assert.Equal(t, expect, result)
		assert.Equal(t, ok1, ok2)
	}

	{ // Partial Bitmap
		arr, index := makeAggregateInput(0x0123456789abcdef)
		expect, ok1 := maxIf(arr, index)
		result, ok2 := Max(arr, index)
		assert.Equal(t, expect, result)
		assert.Equal(t, ok1, ok2)
	}

	{ // Full Bitmap
		arr, index := makeAggregateInput(0xffffffffffffffff)
		expect, ok1 := maxIf(arr, index)
		result, ok2 := Max(arr, index)
		assert.Equal(t, expect, result)
		assert.Equal(t, ok1, ok2)
	}

	{ // Nil Bitmap
		arr, _ := makeAggregateInput(0x0)
		expect, ok1 := maxIf(arr, nil)
		result, ok2 := Max(arr, nil)
		assert.Equal(t, expect, result)
		assert.Equal(t, ok1, ok2)
	}

	{ // Nil Array
		_, index := makeAggregateInput(0x0)
		expect, ok1 := maxIf([]int{}, index)
		result, ok2 := Max([]int{}, index)
		assert.Equal(t, expect, result)
		assert.Equal(t, ok1, ok2)
	}
}

func TestLeftPack(t *testing.T) {
	src, index := makeAggregateInput(0x0123456789abcdef)
	dst := leftPack([]int{}, src, index[0])
	assert.Equal(t, 32, len(dst))
}

// ----------------------------- Naive Aggregation Funcs -----------------------------

func sumIf[T simd.Number](src []T, index Bitmap) (out T) {
	size := min(len(src), len(index)*64)
	for i := 0; i < size; i++ {
		if index.Contains(uint32(i)) {
			out += src[i]
		}
	}
	return
}

func minIf[T simd.Number](src []T, index Bitmap) (T, bool) {
	if len(src) == 0 || index.Count() == 0 {
		return 0, false
	}

	size := min(len(src), len(index)*64)
	out := src[0]
	for i := 0; i < size; i++ {
		if index.Contains(uint32(i)) && src[i] < out {
			out = src[i]
		}
	}
	return out, true
}

func maxIf[T simd.Number](src []T, index Bitmap) (T, bool) {
	if len(src) == 0 || index.Count() == 0 {
		return 0, false
	}

	size := min(len(src), len(index)*64)
	out := src[0]
	for i := 0; i < size; i++ {
		if index.Contains(uint32(i)) && src[i] > out {
			out = src[i]
		}
	}
	return out, true
}

func makeAggregateInput(filter uint64) ([]int, Bitmap) {
	index := make(Bitmap, 0, 8)
	for i := 0; i < 8; i++ {
		index = append(index, filter)
	}

	var arr []int
	for i := 0; i < 500; i++ {
		arr = append(arr, 100+i)
	}
	arr[102] = 50
	arr[101] = 1000
	arr[490] = 40
	arr[491] = 2000
	return arr, index
}

// ----------------------------- Benchmark -----------------------------

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

// run runs a benchmark on a full bitmap
func runFull(b *testing.B, name string, f func(index Bitmap)) {
	count := 1000000
	b.Run(name, func(b *testing.B) {
		index := make(Bitmap, count/64)
		index.Grow(uint32(count))
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
