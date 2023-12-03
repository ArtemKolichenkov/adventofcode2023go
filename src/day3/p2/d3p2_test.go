package main

import (
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve("input.txt")
	}
}

func BenchmarkSolveWithImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolveWithImage("input.txt")
	}
}
