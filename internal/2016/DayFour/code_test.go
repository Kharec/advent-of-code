package dayfour

import (
	"testing"
)

var testInput = []string{
	"aaaaa-bbb-z-y-x-123[abxyz]",
	"a-b-c-d-e-f-g-h-987[abcde]",
	"not-a-real-room-404[oarel]",
	"totally-real-room-200[decoy]",
	"ijmockjgz-storage-5[gjoac]",
}

func TestPartOne(t *testing.T) {
	expected := 1519
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 5
	got := PartTwo(testInput)
	if got != expected {
		t.Errorf("PartTwo() = %d, want %d", got, expected)
	}
}
