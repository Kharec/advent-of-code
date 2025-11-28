package dayone

import "testing"

func TestPartOne(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"(())", "(())", 0},
		{"()()", "()()", 0},
		{"(((", "(((", 3},
		{"(()(()(", "(()(()(", 3},
		{"))(((((", "))(((((", 3},
		{"())", "())", -1},
		{"))(", "))(", -1},
		{")))", ")))", -3},
		{")())())", ")())())", -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PartOne(tt.input)
			if got != tt.expected {
				t.Errorf("PartOne(%q) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{")", ")", 1},
		{"()())", "()())", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PartTwo(tt.input)
			if got != tt.expected {
				t.Errorf("PartTwo(%q) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}
