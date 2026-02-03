package dayeleven

import "testing"

var testInput = []string{
	"aaa: you hhh",
	"you: bbb ccc",
	"bbb: ddd eee",
	"ccc: ddd eee fff",
	"ddd: ggg",
	"eee: out",
	"fff: out",
	"ggg: out",
	"hhh: ccc fff iii",
	"iii: out",
}

func TestPartOne(t *testing.T) {
	expected := 5
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}
