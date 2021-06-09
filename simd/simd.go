//go:generate go run asm.go -out simd.s -stubs stub.go

package simd

import (
	"github.com/klauspost/cpuid/v2"
)

// Supported returns whether the SIMD instructions required are available or not
var Supported = cpuid.CPU.Supports(cpuid.AVX, cpuid.AVX2)
