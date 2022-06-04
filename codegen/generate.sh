#!/bin/bash

PATH="$PATH:./bin"

# Generates AMD64 
function build_amd64 {
    ASM_FILE=$(basename -- ${FILE%.cpp}.s)

    clang -S -Ofast -mavx2 -fno-exceptions -fno-rtti -masm=intel -fno-asynchronous-unwind-tables -mstackrealign -o $ASM_FILE  $FILE
    c2goasm -a -f $ASM_FILE ../$ASM_FILE
    rm $ASM_FILE
}

# Generates ARM64 (untested)
function build_arm64 {
    ASM_FILE=$(basename -- ${FILE%.cpp}.s)

    clang -c -Ofast -arch arm64 -march=armv8+sve -fno-exceptions -fno-rtti -fno-asynchronous-unwind-tables -mstackrealign -o $ASM_FILE $FILE
    c2goasm -a -f $ASM_FILE ../$ASM_FILE
    rm $ASM_FILE
}

# iterate over all source files and generate
for FILE in amd64/*.cpp; do
    build_amd64 $FILE...
done
