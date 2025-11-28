package dayseven

import "testing"

var testInput = []string{
	"light red bags contain 1 bright white bag, 2 muted yellow bags.",
	"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
	"bright white bags contain 1 shiny gold bag.",
	"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
	"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
	"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
	"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
	"faded blue bags contain no other bags.",
	"dotted black bags contain no other bags.",
}

func TestPartOne(t *testing.T) {
	expected := 4
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	secondTestInput := []string{
		"shiny gold bags contain 2 dark red bags.",
		"dark red bags contain 2 dark orange bags.",
		"dark orange bags contain 2 dark yellow bags.",
		"dark yellow bags contain 2 dark green bags.",
		"dark green bags contain 2 dark blue bags.",
		"dark blue bags contain 2 dark violet bags.",
		"dark violet bags contain no other bags.",
	}

	t.Run("First Example", func(t *testing.T) {
		expected := 32
		got := PartTwo(testInput)
		if got != expected {
			t.Errorf("PartTwo() = %d, want %d", got, expected)
		}
	})

	t.Run("Second Example", func(t *testing.T) {
		expected := 126
		got := PartTwo(secondTestInput)
		if got != expected {
			t.Errorf("PartTwo() = %d, want %d", got, expected)
		}
	})
}
