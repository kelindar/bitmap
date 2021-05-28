# Zero-Allocation Bitmap Index (Bitset) in Go

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kelindar/bitmap)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/kelindar/bitmap)](https://pkg.go.dev/github.com/kelindar/bitmap)
[![Go Report Card](https://goreportcard.com/badge/github.com/kelindar/bitmap)](https://goreportcard.com/report/github.com/kelindar/bitmap)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

This package contaisn a bitmap index which is backed by `uint64` slice, easily encodable to/from a `[]byte` without copying memory around so it can be present
in both disk and memory. As opposed to something as [roaring bitmaps](github.com/RoaringBitmap/roaring), this is a simple impementation designed to be used for small to medium dense collections.

I've used this package to build a columnar in-memory datastore, so if you want to see how it can be used for indexing, have a look at [kelindar/columnar](https://github.com/kelindar/columnar). I'd like to specifically point out the indexing part and how bitmaps can be used as a good alternative to B*Trees and Hash Maps.

## Features

 * Zero-allocation (see benchmarks below) on almost all of the important APIs.
 * 1-2 nanosecond on single-bit operations (set/remove/contains).
 * Support for `and`, `and not`, `or` and `xor` which allows you to compute intersect, difference, union and symmetric difference between 2 bitmaps.
 * Support for `min`, `max`, `count`, and `first-zero` which is very useful for building free-lists using a bitmap index.
 * Reusable and can be pooled, providing `clone` with a destination and `clear` operations.
 * Can be encoded to binary without copy as well as optional stream `WriteTo` method.
 * Support for iteration via `Range` method and filtering via `Filter` method.

## Example Usage

In its simplest form, you can use the bitmap as a bitset, set and remove bits. This is quite useful as an index (free/fill-list) for an array of data.

```go
import "github.com/kelindar/bitmap"
```

```go
bitmap := make(bitmap.Bitmap, 0, 8) // 8*64 = 512 elements pre-allocated
bitmap.Set(300)         // sets 250-th bit
bitmap.Set(400)         // sets 400-th bit
bitmap.Set(600)         // sets 600-th bit (auto-resized)
bitmap.Contains(300)    // returns true
bitmap.Contains(301)    // returns false
bitmap.Remove(400)      // clears 400-th bit

// Min, Max, Count
min, ok :=bitmap.Min()  // returns 300
max, ok := bitmap.Max() // returns 600
count := bitmap.Count() // returns 2
```

The bits in the bitmap can also be iterated over using the `Range` method. It is a simple loop which iterates over and calls a callback. If the callback returns false, then the iteration is halted (similar to `sync.Map`).

```go
// Iterate over the bits in the bitmap
bitmap.Range(func(x uint32) bool {
    println(x)
    return true
})
```

Another way of iterating is using the `Filter` method. It iterates similarly to `Range` but the callback returns a boolean value, and if it returns `false` then the current bit will be cleared in the underlying bitmap. You could accomplish the same using `Range` and `Remove` but `Filter` is significantly faster.

```go
// Filter iterates over the bits and applies a callback
bitmap.Filter(func(x uint32) bool {
    return x % 2 == 0
})
```

Bitmaps are also extremely useful as they support boolean operations very efficiently. This library contains `And`, `AndNot`, `Or` and `Xor`.

```go
// And computes the intersection between two bitmaps and stores the result in the current bitmap
a := Bitmap{0b0011}
a.And(Bitmap{0b0101})

// AndNot computes the difference between two bitmaps and stores the result in the current bitmap
a := Bitmap{0b0011}
a.AndNot(Bitmap{0b0101})

// Or computes the union between two bitmaps and stores the result in the current bitmap
a := Bitmap{0b0011}
a.Or(Bitmap{0b0101})

// Xor computes the symmetric difference between two bitmaps and stores the result in the current bitmap
a := Bitmap{0b0011}
a.Xor(Bitmap{0b0101})
```

## Benchmarks
Benchmarks below were run on a large pre-allocated bitmap (slice of 100 pages, 6400 items).

```
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkBitmap/set-8          607986201    2.004 ns/op    0 B/op    0 allocs/op
BenchmarkBitmap/remove-8       759495112    1.564 ns/op    0 B/op    0 allocs/op
BenchmarkBitmap/contains-8     916055708    1.306 ns/op    0 B/op    0 allocs/op
BenchmarkBitmap/and-8          30769704     39.00 ns/op    0 B/op    0 allocs/op
BenchmarkBitmap/andnot-8       23518297     51.11 ns/op    0 B/op    0 allocs/op
BenchmarkBitmap/or-8           24000480     50.92 ns/op    0 B/op    0 allocs/op
BenchmarkBitmap/xor-8          23529734     51.24 ns/op    0 B/op    0 allocs/op
BenchmarkBitmap/clear-8        206332054    5.845 ns/op    0 B/op    0 allocs/op
BenchmarkBitmap/ones-8         39965895	    29.47 ns/op    0 B/op    0 allocs/op
BenchmarkBitmap/first-zero-8   30001574	    40.15 ns/op    0 B/op    0 allocs/op
BenchmarkBitmap/min-8          375470745    3.295 ns/op    0 B/op    0 allocs/op
BenchmarkBitmap/max-8          681821280    1.740 ns/op    0 B/op    0 allocs/op
BenchmarkBitmap/count-8        32433483     35.64 ns/op    0 B/op    0 allocs/op
BenchmarkBitmap/clone-8        100000000    11.77 ns/op    0 B/op    0 allocs/op
```