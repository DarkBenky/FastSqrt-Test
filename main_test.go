package main

import (
	"fmt"
	"math"
	"testing"
)

// Standard sqrt using math.Sqrt
func StandardSqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

// Iterative sqrt using 5 iterations of Newton's method
func IterativeSqrt(x float32) float32 {
	if x == 0 {
		return 0
	}
	guess := x / 2.0
	for i := 0; i < 5; i++ {
		guess = 0.5 * (guess + x/guess)
	}
	return guess
}

// Optimized sqrt using bit manipulation
func OptimizedSqrt(x float32) float32 {
	if x == 0 {
		return 0
	}
	halfX := 0.5 * x
	i := math.Float32bits(x)
	i = 0x5f3759df - (i >> 1) // Magic number for float32
	guess := math.Float32frombits(i)
	guess = guess * (1.5 - halfX*guess*guess)
	guess = guess * (1.5 - halfX*guess*guess)
	return 1 / guess
}

// Benchmarking
func BenchmarkStandardSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StandardSqrt(525.75)
	}
}

func BenchmarkIterativeSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IterativeSqrt(525.75)
	}
}

func BenchmarkOptimizedSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = OptimizedSqrt(525.75)
	}
}

func main() {
	// Run benchmarks
	fmt.Println("Run `go test -bench=.` to see the benchmark results.")
}
