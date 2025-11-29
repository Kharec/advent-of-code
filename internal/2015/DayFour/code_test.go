package dayfour

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected int
	}{
		{"abcdef", []byte("abcdef"), 609043},
		{"pqrstuv", []byte("pqrstuv"), 1048970},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PartOne(tt.input)
			if got != tt.expected {
				t.Errorf("PartOne() = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected int
	}{
		{"abcdef", []byte("abcdef"), 6742839},
		{"pqrstuv", []byte("pqrstuv"), 5714438},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PartTwo(tt.input)
			if got != tt.expected {
				t.Errorf("PartTwo() = %d, want %d", got, tt.expected)
			}
		})
	}
}
