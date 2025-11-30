package daytwo

import "testing"

var testInput = []string{
	"abcdef",
	"bababc",
	"abbcde",
	"abcccd",
	"aabcdd",
	"abcdee",
	"ababab",
}

func TestPartOne(t *testing.T) {
	expected := 12
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}
