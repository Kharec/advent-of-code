package daythree

import (
	"advent-of-code/internal/registry"
	"os"
	"strings"
)

func init() {
	registry.Register("2020D3", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(input []string) int {
	trees := 0
	column := 0
	for row := range input {
		if len(input[row]) == 0 {
			continue
		}
		if input[row][column%len(input[row])] == '#' {
			trees++
		}
		column += 3
	}
	return trees
}

func PartTwo(input []string) int {
	result := 1
	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, slope := range slopes {
		trees := 0
		column := 0
		for row := 0; row < len(input); row += slope[1] {
			if len(input[row]) == 0 {
				continue
			}
			if input[row][column%len(input[row])] == '#' {
				trees++
			}
			column += slope[0]
		}
		result *= trees
	}
	return result
}

