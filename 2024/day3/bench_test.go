package day3

import "testing"

func BenchmarkPart1(b *testing.B) {
	// Setup code: Ensure you have a valid file for benchmarking.
	b.ReportAllocs()
	filePath := "input.txt" // Path to a sample input file.

	// Run the benchmark N times.
	for i := 0; i < b.N; i++ {
		_ = part1(filePath) // Run the function you're benchmarking.
	}
}

func BenchmarkPart2(b *testing.B) {
	// Setup code: Ensure you have a valid file for benchmarking.
	b.ReportAllocs()
	filePath := "input.txt" // Path to a sample input file.

	// Run the benchmark N times.
	for i := 0; i < b.N; i++ {
		_ = part2(filePath) // Run the function you're benchmarking.
	}
}
