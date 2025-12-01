package dayseven

import "testing"

var testInput = []string{
	"123 -> x",
	"456 -> y",
	"x AND y -> d",
	"x OR y -> e",
	"x LSHIFT 2 -> f",
	"y RSHIFT 2 -> g",
	"NOT x -> h",
	"NOT y -> i",
	"d -> a",
}

func TestPartOne(t *testing.T) {
	expected := 72
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}
