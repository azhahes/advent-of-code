package day4

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		name     string
		filePath string
		expected int
	}{
		{
			name:     "Sample input",
			filePath: "sample-input.txt",
			expected: 18,
		},
		{
			name:     "Actual input",
			filePath: "input.txt",
			expected: 2583,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := part1(tc.filePath)

			if result != tc.expected {
				t.Errorf("Failed %s: got %d, want %d", tc.name, result, tc.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		name     string
		filePath string
		expected int
	}{
		{
			name:     "Sample input",
			filePath: "sample-input.txt",
			expected: 9,
		},
		{
			name:     "Actual input",
			filePath: "input.txt",
			expected: 1978,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := part2(tc.filePath)

			if result != tc.expected {
				t.Errorf("Failed %s: got %d, want %d", tc.name, result, tc.expected)
			}
		})
	}
}
