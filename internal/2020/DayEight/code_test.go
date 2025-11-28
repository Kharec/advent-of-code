package dayeight

import "testing"

var testInput = []string{
	"nop +0",
	"acc +1",
	"jmp +4",
	"acc +3",
	"jmp -3",
	"acc -99",
	"acc +1",
	"jmp -4",
	"acc +6",
}

func TestPartOne(t *testing.T) {
	expected := 5
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}
