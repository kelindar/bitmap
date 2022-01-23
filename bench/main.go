package main

import (
	"fmt"
	"time"

	"github.com/kelindar/bitmap"
)

func main() {
	const size = 10000000
	const iter = 500
	const inner = 500

	a := createBitmap(size)
	b := createBitmap(size)

	for i := 0; i < iter; i++ {
		start := time.Now()

		for j := 0; j < inner; j++ {
			a.And(b, b, b, b)
			//a.And(b)
			//a.And(b)
			//a.And(b)
			//a.And(b)
		}

		fmt.Printf("iteration %v took %v...\n", i*inner, time.Now().Sub(start))
	}

}

func createBitmap(size int) bitmap.Bitmap {
	index := make(bitmap.Bitmap, size/64)
	index.Grow(uint32(size - 1))
	for i := 0; i < len(index); i++ {
		index[i] = 0xf0f0f0f0f0f0f0f0
	}
	return index
}
