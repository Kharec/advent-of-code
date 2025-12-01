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
