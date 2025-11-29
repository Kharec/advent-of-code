package dayfive

import "testing"

func TestPartOne(t *testing.T) {
	input := []string{
		"ugknbfddgicrmopn",
		"aaa",
		"jchzalrnumimnmhp",
		"haegwjzuvuyypxyu",
		"dvszwmarrgswjxmb",
	}
	expected := 2
	got := PartOne(input)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := []string{
		"qjhvhtzxzqqjkmpb",
		"xxyxx",
		"uurcxstgmygtbstg",
		"ieodomkazucvgmuy",
	}
	expected := 2
	got := PartTwo(input)
	if got != expected {
		t.Errorf("PartTwo() = %d, want %d", got, expected)
	}
}
