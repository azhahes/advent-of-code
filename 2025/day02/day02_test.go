package day02

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
			expected: 0,
		},
		{
			name:     "Actual input",
			filePath: "input.txt",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Part1(tc.filePath)

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
			expected: 0,
		},
		{
			name:     "Actual input",
			filePath: "input.txt",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Part2(tc.filePath)

			if result != tc.expected {
				t.Errorf("Failed %s: got %d, want %d", tc.name, result, tc.expected)
			}
		})
	}
}
