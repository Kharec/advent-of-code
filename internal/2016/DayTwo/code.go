package daytwo

import (
	"advent-of-code/internal/registry"
	"os"
	"strings"
)

func init() {
	registry.Register("2016D2", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(instructions []string) int {
	row, column := 1, 1

	keypad := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	code := 0

	for _, line := range instructions {
		for _, move := range line {
			switch move {
			case 'U':
				if row > 0 {
					row--
				}
			case 'D':
				if row < 2 {
					row++
				}
			case 'L':
				if column > 0 {
					column--
				}
			case 'R':
				if column < 2 {
					column++
				}
			}
		}
		code = code*10 + keypad[row][column]
	}

	return code
}

func PartTwo(data []string) int {
	return 0
}
