package main

import "testing"

func TestPartOne(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	expected := 514579
	got := PartOne(input)
	if got != expected {
		t.Errorf("PartOne(%v) = %d, want %d", input, got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	expected := 241861950
	got := PartTwo(input)
	if got != expected {
		t.Errorf("PartTwo(%v) = %d, want %d", input, got, expected)
	}
}
