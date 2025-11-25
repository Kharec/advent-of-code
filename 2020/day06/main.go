package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parseInput(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
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

func main() {
	input := parseInput("input.txt")
	fmt.Println("Part 1:", PartOne(input))
	fmt.Println("Part 2:", PartTwo(input))
}
