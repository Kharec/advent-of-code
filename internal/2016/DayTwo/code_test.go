package daytwo

import (
	"testing"
)

var testInput = []string{"ULL", "RRDDD", "LURDL", "UUUUD"}

func TestPartOne(t *testing.T) {
	expected := 1985
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}
