package dayone

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "Example 1",
			input:    []string{"R2", "L3"},
			expected: 5,
		},
		{
			name:     "Example 2",
			input:    []string{"R2", "R2", "R2"},
			expected: 2,
		},
		{
			name:     "Example 3",
			input:    []string{"R5", "L5", "R5", "R3"},
			expected: 12,
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
