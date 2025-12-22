package daytwo

import "testing"

var testInput = [][]int{
	{5, 1, 9, 5},
	{7, 5, 3},
	{2, 4, 6, 8},
}

func TestPartOne(t *testing.T) {
	got := PartOne(testInput)
	expected := 18
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}
