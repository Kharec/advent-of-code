package dayfive

import "testing"

var testInput = []int{
	0,
	3,
	0,
	1,
	-3,
}

func TestPartOne(t *testing.T) {
	expected := 5
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}
