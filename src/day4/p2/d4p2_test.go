package main

import (
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	// No pre-parsing results:
	// Run: 1
	// op time: 13673700250 ns/op
	// Total time: 13.843s

	// Results with:
	// - pre-parsing string content into Card structs
	// copy:
	// Run: 1
	// op time: 4327610667 ns/op
	// Total time: 4.499s
	// append
	// Run: 1
	// op time: 3837964958 ns/op
	// Total time: 4.008s

	// Reults with:
	// - pre-computing winCount for each card in advance
	// Run: 2
	// op time: 550838250 ns/op
	// Total time: 1.854s

	for i := 0; i < b.N; i++ {
		Solve("input.txt")
	}
}
