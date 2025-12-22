package daythree

import "testing"

func TestPartOne(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "square 1 is 0 steps",
			input:    1,
			expected: 0,
		},
		{
			name:     "square 12 is 3 steps",
			input:    12,
			expected: 3,
		},
		{
			name:     "square 23 is 2 steps",
			input:    23,
			expected: 2,
		},
		{
			name:     "square 1024 is 31 steps",
			input:    1024,
			expected: 31,
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
