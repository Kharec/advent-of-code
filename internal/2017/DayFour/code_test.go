package dayfour

import "testing"

func TestPartOne(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "aa bb cc dd ee is valid",
			input:    []string{"aa bb cc dd ee"},
			expected: 1,
		},
		{
			name:     "aa bb cc dd aa is not valid",
			input:    []string{"aa bb cc dd aa"},
			expected: 0,
		},
		{
			name:     "aa bb cc dd aaa is valid",
			input:    []string{"aa bb cc dd aaa"},
			expected: 1,
		},
		{
			name:     "multiple passphrases",
			input:    []string{"aa bb cc dd ee", "aa bb cc dd aa", "aa bb cc dd aaa"},
			expected: 2,
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
