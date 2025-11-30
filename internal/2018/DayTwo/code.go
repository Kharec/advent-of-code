package daytwo

import (
	"advent-of-code/internal/registry"
	"fmt"
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

func PartTwo(data []string) int {
	for idx := range data {
		for otherIdx := idx + 1; otherIdx < len(data); otherIdx++ {
			if data[idx] == "" || data[otherIdx] == "" {
				continue
			}

			differenceCount := 0
			differingPosition := -1

			if len(data[idx]) != len(data[otherIdx]) {
				continue
			}

			for position := 0; position < len(data[idx]); position++ {
				if data[idx][position] != data[otherIdx][position] {
					differenceCount++
					differingPosition = position
				}
			}

			if differenceCount == 1 {
				common := data[idx][:differingPosition] + data[idx][differingPosition+1:]
				fmt.Println(common)
				return 0
			}
		}
	}
	return 0
}
