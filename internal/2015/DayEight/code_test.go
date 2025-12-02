package dayeight

import "testing"

var testInput = []string{
	`""`,
	`"abc"`,
	`"aaa\"aaa"`,
	`"\x27"`,
}

func TestPartOne(t *testing.T) {
	expected := 12
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}
