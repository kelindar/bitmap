#!/bin/bash

./bin/gocc simd_avx.c --arch amd64 -O1 -mavx2 -mfma -masm=intel -o ../
./bin/gocc simd_neon.c --arch arm64 -O3 -o ../
./bin/gocc simd_apple.c --arch arm64 -O3 -o ../