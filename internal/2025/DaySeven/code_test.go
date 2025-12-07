package dayseven

import "testing"

var testInput = []string{
	".......S.......",
	"...............",
	".......^.......",
	"...............",
	"......^.^......",
	"...............",
	".....^.^.^.....",
	"...............",
	"....^.^...^....",
	"...............",
	"...^.^...^.^...",
	"...............",
	"..^...^.....^..",
	"...............",
	".^.^.^.^.^...^.",
	"...............",
}

func TestPartOne(t *testing.T) {
	expected := 21
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 40
	got := PartTwo(testInput)
	if got != expected {
		t.Errorf("PartTwo() = %d, want %d", got, expected)
	}
}
