package main

import "testing"

func TestPartOne(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	expected := 2
	got := PartOne(input)
	if got != expected {
		t.Errorf("PartOne(%v) = %d, want %d", input, got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	expected := 1
	got := PartTwo(input)
	if got != expected {
		t.Errorf("PartTwo(%v) = %d, want %d", input, got, expected)
	}
}
