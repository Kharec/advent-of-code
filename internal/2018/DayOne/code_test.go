package dayone

import "testing"

var testInput = []int{
	+1,
	-2,
	3,
	1,
}

func TestPartOne(t *testing.T) {
	expected := 3
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 2
	got := PartTwo(testInput)
	if got != expected {
		t.Errorf("PartTwo() = %d, want %d", got, expected)
	}
}
