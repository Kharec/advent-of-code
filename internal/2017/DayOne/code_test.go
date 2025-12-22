package dayone

import "testing"

func TestPartOne(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "1122 produces 3",
			input:    []int{1, 1, 2, 2},
			expected: 3,
		},
		{
			name:     "1111 produces 4",
			input:    []int{1, 1, 1, 1},
			expected: 4,
		},
		{
			name:     "1234 produces 0",
			input:    []int{1, 2, 3, 4},
			expected: 0,
		},
		{
			name:     "91212129 produces 9",
			input:    []int{9, 1, 2, 1, 2, 1, 2, 9},
			expected: 9,
		},
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
		input    []int
		expected int
	}{
		{
			name:     "1212 produces 6",
			input:    []int{1, 2, 1, 2},
			expected: 6,
		},
		{
			name:     "1221 produces 0",
			input:    []int{1, 2, 2, 1},
			expected: 0,
		},
		{
			name:     "123425 produces 4",
			input:    []int{1, 2, 3, 4, 2, 5},
			expected: 4,
		},
		{
			name:     "123123 produces 12",
			input:    []int{1, 2, 3, 1, 2, 3},
			expected: 12,
		},
		{
			name:     "12131415 produces 4",
			input:    []int{1, 2, 1, 3, 1, 4, 1, 5},
			expected: 4,
		},
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
