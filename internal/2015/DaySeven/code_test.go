package dayseven

import (
	"testing"
)

var testInput = map[string]string{
	"x": "123",
	"y": "456",
	"d": "x AND y",
	"e": "x OR y",
	"f": "x LSHIFT 2",
	"g": "y RSHIFT 2",
	"h": "NOT x",
	"i": "NOT y",
	"a": "d",
}

func TestPartOne(t *testing.T) {
	expected := 72
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	instructions := map[string]string{
		"x": "10",
		"y": "20",
		"z": "x AND y",
		"b": "z",
		"w": "b LSHIFT 1",
		"v": "NOT b",
		"u": "w OR v",
		"a": "u",
	}

	partOneResult := PartOne(instructions)
	bValue := uint16(partOneResult)
	w := bValue << 1
	v := ^bValue
	u := w | v
	expected := int(u)

	got := PartTwo(instructions)
	if got != expected {
		t.Errorf("PartTwo() = %d, want %d (PartOne result: %d)", got, expected, partOneResult)
	}
}
