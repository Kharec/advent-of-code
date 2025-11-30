package daytwo

import (
	"advent-of-code/internal/registry"
	"os"
	"strings"
)

func init() {
	registry.Register("2018D2", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	countWithTwo := 0
	countWithThree := 0

	for _, boxID := range data {
		charCounts := make(map[rune]int)
		for _, char := range boxID {
			charCounts[char]++
		}

		hasExactlyTwo := false
		hasExactlyThree := false

		for _, count := range charCounts {
			switch count {
			case 2:
				hasExactlyTwo = true
			case 3:
				hasExactlyThree = true
			}
		}

		if hasExactlyTwo {
			countWithTwo++
		}
		if hasExactlyThree {
			countWithThree++
		}
	}

	return countWithTwo * countWithThree
}
