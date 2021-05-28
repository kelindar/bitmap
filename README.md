# Bitmap index in Go

This package contaisn a bitmap index which is backed by `uint64` slice, easily encodable to/from a `[]byte` without copying memory around so it can be present
in both disk and memory. As opposed to something as [roaring bitmaps](github.com/RoaringBitmap/roaring), this is a simple impementation designed to be used for small to medium dense collections.


## Benchmarks
Benchmarks below were run on a large pre-allocated bitmap (slice of 100 elements).

```
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkBitmap/set-8           606936064    1.984 ns/op      0 B/op    0 allocs/op
BenchmarkBitmap/remove-8        763688641    1.559 ns/op      0 B/op    0 allocs/op
BenchmarkBitmap/contains-8      915751614    1.299 ns/op      0 B/op    0 allocs/op
BenchmarkBitmap/and-8           31582188     38.38 ns/op      0 B/op    0 allocs/op
BenchmarkBitmap/andnot-8        23472387     50.65 ns/op      0 B/op    0 allocs/op
BenchmarkBitmap/or-8            23695324     50.98 ns/op      0 B/op    0 allocs/op
BenchmarkBitmap/xor-8           23197326     51.23 ns/op      0 B/op    0 allocs/op
BenchmarkBitmap/clear-8         205686652    5.897 ns/op      0 B/op    0 allocs/op
BenchmarkBitmap/ones-8          40554514     29.33 ns/op      0 B/op    0 allocs/op
BenchmarkBitmap/first-zero-8    30062228     40.39 ns/op      0 B/op    0 allocs/op
BenchmarkBitmap/min-8           381730626    3.131 ns/op      0 B/op    0 allocs/op
BenchmarkBitmap/max-8           684005834    1.756 ns/op      0 B/op    0 allocs/op
BenchmarkRange/count-8          34223232     35.20 ns/op      0 B/op    0 allocs/op
BenchmarkBitmap/clone-8         7676722      143.9 ns/op    896 B/op    1 allocs/op
```