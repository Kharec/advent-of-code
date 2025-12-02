package daynine

import "testing"

var testInput = []string{
	"London to Dublin = 464",
	"London to Belfast = 518",
	"Dublin to Belfast = 141",
}

func TestPartOne(t *testing.T) {
	expected := 605
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 982
	got := PartTwo(testInput)
	if got != expected {
		t.Errorf("PartTwo() = %d, want %d", got, expected)
	}
}
