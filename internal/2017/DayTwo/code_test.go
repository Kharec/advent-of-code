package daytwo

import "testing"

func TestPartOne(t *testing.T) {
	input := [][]int{
		{5, 1, 9, 5},
		{7, 5, 3},
		{2, 4, 6, 8},
	}
	got := PartOne(input)
	expected := 18
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := [][]int{
		{5, 9, 2, 8},
		{9, 4, 7, 3},
		{3, 8, 6, 5},
	}
	got := PartTwo(input)
	expected := 9
	if got != expected {
		t.Errorf("PartTwo() = %d, want %d", got, expected)
	}
}
