package daytwelve

import "testing"

func TestPartOne(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"[1,2,3]", "[1,2,3]", 6},
		{"{\"a\":2,\"b\":4}", "{\"a\":2,\"b\":4}", 6},
		{"[[[3]]]", "[[[3]]]", 3},
		{"{\"a\":{\"b\":4},\"c\":-1}", "{\"a\":{\"b\":4},\"c\":-1}", 3},
		{"{\"a\":[-1,1]}", "{\"a\":[-1,1]}", 0},
		{"[-1,{\"a\":1}]", "[-1,{\"a\":1}]", 0},
		{"[]", "[]", 0},
		{"{}", "{}", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PartOne([]string{tt.input})
			if got != tt.expected {
				t.Errorf("PartOne() = %d, want %d", got, tt.expected)
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
		{"[1,2,3]", "[1,2,3]", 6},
		{"[1,{\"c\":\"red\",\"b\":2},3]", "[1,{\"c\":\"red\",\"b\":2},3]", 4},
		{"{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}", "{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}", 0},
		{"[1,\"red\",5]", "[1,\"red\",5]", 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PartTwo([]string{tt.input})
			if got != tt.expected {
				t.Errorf("PartTwo() = %d, want %d", got, tt.expected)
			}
		})
	}
}
