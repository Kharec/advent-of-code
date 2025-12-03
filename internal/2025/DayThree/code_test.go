package daythree

import "testing"

var testInput = [][]int{
	{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
	{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
	{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
	{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
}

func TestPartOne(t *testing.T) {
	expected := 357
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 3121910778619
	got := PartTwo(testInput)
	if got != expected {
		t.Errorf("PartTwo() = %d, want %d", got, expected)
	}
}
