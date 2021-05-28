# Bitmap index in Go

This package contaisn a bitmap index which is backed by `uint64` slice, easily encodable to/from a `[]byte` without copying memory around so it can be present
in both disk and memory. As opposed to something as [roaring bitmaps](github.com/RoaringBitmap/roaring), this is a simple impementation designed to be used for small to medium dense collections.


## Benchmarks
Benchmarks below were run on a large pre-allocated bitmap (slice of 100 elements).

```
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkBitmap/set-8        596428585   1.975 ns/op   0 B/op   0 allocs/op
BenchmarkBitmap/remove-8     765963639   1.541 ns/op   0 B/op   0 allocs/op
BenchmarkBitmap/contains-8   909556022   1.295 ns/op   0 B/op   0 allocs/op
BenchmarkBitmap/and-8        29702161    38.92 ns/op   0 B/op   0 allocs/op
BenchmarkBitmap/andnot-8     23749771    50.42 ns/op   0 B/op   0 allocs/op
BenchmarkBitmap/or-8         23796303    50.42 ns/op   0 B/op   0 allocs/op
BenchmarkBitmap/xor-8        23703469    50.33 ns/op   0 B/op   0 allocs/op
```