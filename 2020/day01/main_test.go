package main

import "testing"

var testInput = []int{1721, 979, 366, 299, 675, 1456}

func TestPartOne(t *testing.T) {
	expected := 514579
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne(%v) = %d, want %d", testInput, got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 241861950
	got := PartTwo(testInput)
	if got != expected {
		t.Errorf("PartTwo(%v) = %d, want %d", testInput, got, expected)
	}
}
