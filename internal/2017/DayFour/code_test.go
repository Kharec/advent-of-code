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

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "abcde fghij is valid",
			input:    []string{"abcde fghij"},
			expected: 1,
		},
		{
			name:     "abcde xyz ecdab is not valid",
			input:    []string{"abcde xyz ecdab"},
			expected: 0,
		},
		{
			name:     "a ab abc abd abf abj is valid",
			input:    []string{"a ab abc abd abf abj"},
			expected: 1,
		},
		{
			name:     "iiii oiii ooii oooi oooo is valid",
			input:    []string{"iiii oiii ooii oooi oooo"},
			expected: 1,
		},
		{
			name:     "oiii ioii iioi iiio is not valid",
			input:    []string{"oiii ioii iioi iiio"},
			expected: 0,
		},
		{
			name:     "multiple passphrases",
			input:    []string{"abcde fghij", "abcde xyz ecdab", "a ab abc abd abf abj", "oiii ioii iioi iiio"},
			expected: 2,
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
