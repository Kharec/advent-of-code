package daysix

import "testing"

var testInput = []string{
	"turn on 0,0 through 999,999",
	"toggle 0,0 through 999,0",
	"turn off 499,499 through 500,500",
}

func TestPartOne(t *testing.T) {
	expected := 998996
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}
