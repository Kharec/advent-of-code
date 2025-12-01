package daytwo

import "testing"

var testInput = []string{
	"A Y",
	"B X",
	"C Z",
}

func TestPartOne(t *testing.T) {
	expected := 15
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 12
	got := PartTwo(testInput)
	if got != expected {
		t.Errorf("PartTwo() = %d, want %d", got, expected)
	}
}
