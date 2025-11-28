package daythree

import "testing"

func TestPartOne(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{">", ">", 2},
		{"^>v<", "^>v<", 4},
		{"^v^v^v^v^v", "^v^v^v^v^v", 2},
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
