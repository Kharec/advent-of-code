package dayone

import "testing"

var testInput = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

func TestPartOne(t *testing.T) {
	expected := 7
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 5
	got := PartTwo(testInput)
	if got != expected {
		t.Errorf("PartTwo() = %d, want %d", got, expected)
	}
}

