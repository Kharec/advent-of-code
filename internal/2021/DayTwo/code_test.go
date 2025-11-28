package daytwo

import "testing"

var testInput = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func TestPartOne(t *testing.T) {
	expected := 150
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 900
	got := PartTwo(testInput)
	if got != expected {
		t.Errorf("PartTwo() = %d, want %d", got, expected)
	}
}
