package main

import "testing"

var testInput = []string{
	"1-3 a: abcde",
	"1-3 b: cdefg",
	"2-9 c: ccccccccc",
}

func TestPartOne(t *testing.T) {
	expected := 2
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne(%v) = %d, want %d", testInput, got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 1
	got := PartTwo(testInput)
	if got != expected {
		t.Errorf("PartTwo(%v) = %d, want %d", testInput, got, expected)
	}
}
