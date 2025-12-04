package dayfour

import "testing"

var testInput = []string{
	"..@@.@@@@.",
	"@@@.@.@.@@",
	"@@@@@.@.@@",
	"@.@@@@..@.",
	"@@.@@@@.@@",
	".@@@@@@@.@",
	".@.@.@.@@@",
	"@.@@@.@@@@",
	".@@@@@@@@.",
	"@.@.@@@.@.",
}

func TestPartOne(t *testing.T) {
	expected := 13
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}
