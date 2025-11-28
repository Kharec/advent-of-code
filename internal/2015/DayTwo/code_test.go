package daytwo

import "testing"

func TestPartOne(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"2x3x4", "2x3x4", 58},
		{"1x1x10", "1x1x10", 43},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PartOne([]string{tt.input})
			if got != tt.expected {
				t.Errorf("PartOne(%q) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}
