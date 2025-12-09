package daythree

import "testing"

func TestPartOne(t *testing.T) {
	input := [][3]int{{5, 10, 25}}
	expected := 0
	got := PartOne(input)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := [][3]int{
		{101, 301, 501},
		{102, 302, 502},
		{103, 303, 503},
		{201, 401, 601},
		{202, 402, 602},
		{203, 403, 603},
	}
	expected := 6
	got := PartTwo(input)
	if got != expected {
		t.Errorf("PartTwo() = %d, want %d", got, expected)
	}
}
