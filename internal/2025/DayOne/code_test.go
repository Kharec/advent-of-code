package dayone

import "testing"

var testInput = []string{
	"L68",
	"L30",
	"R48",
	"L5",
	"R60",
	"L55",
	"L1",
	"L99",
	"R14",
	"L82",
}

func TestPartOne(t *testing.T) {
	expected := 3
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 6
	got := PartTwo(testInput)
	if got != expected {
		t.Errorf("PartTwo() = %d, want %d", got, expected)
	}
}
