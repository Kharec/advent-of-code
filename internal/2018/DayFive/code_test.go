package dayfive

import "testing"

var testInput = "dabAcCaCBAcCcaDA"

func TestPartOne(t *testing.T) {
	expected := 10
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 4
	got := PartTwo(testInput)
	if got != expected {
		t.Errorf("PartTwo() = %d, want %d", got, expected)
	}
}
