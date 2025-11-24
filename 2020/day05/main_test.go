package main

import "testing"

var testInput = []string{
	"FBFBBFFRLR",
	"BFFFBBFRRR",
	"FFFBBBFRRR",
	"BBFFBBFRLL",
}

func TestPartOne(t *testing.T) {
	expected := 820
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}
