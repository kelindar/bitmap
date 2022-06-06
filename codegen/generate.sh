#!/bin/bash

PATH="$PATH:./bin"
CLANG_OPTS="-mno-red-zone -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -ffast-math -O1"

# Generates AMD64 
function build_avx2_amd64 {
    SRC="bitmap_avx2_amd64.cpp"
    ASM="bitmap_avx2_amd64.s"
    clang-14 -S -mavx2 -masm=intel $CLANG_OPTS -o $ASM $SRC
    c2goasm -a -f $ASM ../$ASM
    rm $ASM
}

# Generates ARM64 (untested)
function build_neon_arm64 {
    SRC="bitmap_neon_arm64.cpp"
    ASM="bitmap_neon_arm64.s"
    clang-14 -S -arch arm64 $CLANG_OPTS -o $ASM $SRC
    c2goasm -a -f $ASM ../$ASM
    rm $ASM
}

build_avx2_amd64