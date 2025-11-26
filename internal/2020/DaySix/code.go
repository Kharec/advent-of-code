package daysix

import (
	"advent-of-code/internal/registry"
	"os"
	"strings"
)

func init() {
	registry.Register("2020D6", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

func PartOne(input []string) int {
	total := 0
	groupAnswers := make(map[rune]bool)

	for idx, line := range input {
		if line != "" {
			for _, char := range line {
				if char >= 'a' && char <= 'z' {
					groupAnswers[char] = true
				}
			}
		}

		if line == "" || idx == len(input)-1 {
			total += len(groupAnswers)
			groupAnswers = make(map[rune]bool)
		}
	}

	total += len(groupAnswers)

	return total
}

func PartTwo(input []string) int {
	total := 0
	groupAnswers := make(map[rune]int)
	groupSize := 0

	for idx, line := range input {
		if line != "" {
			groupSize++
			for _, char := range line {
				if char >= 'a' && char <= 'z' {
					groupAnswers[char]++
				}
			}
		}

		if line == "" || idx == len(input)-1 {
			for _, count := range groupAnswers {
				if count == groupSize {
					total++
				}
			}
			groupAnswers = make(map[rune]int)
			groupSize = 0
		}
	}

	return total
}

