package dayfive

import "testing"

func TestPartOne(t *testing.T) {
	input := []string{
		"FBFBBFFRLR",
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	}
	expected := 820
	got := PartOne(input)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := []string{
		"FFFFFFFLLL",
		"FFFFFFFLLR",
		"FFFFFFFLRL",
		"FFFFFFFLRR",
		"FFFFFFFRLR",
		"FFFFFFFRRL",
		"FFFFFFFRRR",
	}
	expected := 4
	got := PartTwo(input)
	if got != expected {
		t.Errorf("PartTwo() = %d, want %d", got, expected)
	}
}
