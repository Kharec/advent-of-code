package dayfour

import (
	"os"
	"strings"

	"advent-of-code/internal/registry"
)

type position struct {
	row, column int
}

func init() {
	registry.Register("2025D4", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	isRoll := make(map[position]bool)

	for row := range data {
		for column := 0; column < len(data[row]); column++ {
			if data[row][column] == '@' {
				isRoll[position{row, column}] = true
			}
		}
	}

	accessibleRolls := 0
	directions := []position{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for pos := range isRoll {
		adjacentRolls := 0
		for _, direction := range directions {
			neighbor := position{pos.row + direction.row, pos.column + direction.column}
			if isRoll[neighbor] {
				adjacentRolls++
			}
		}

		if adjacentRolls < 4 {
			accessibleRolls++
		}
	}

	return accessibleRolls
}

func PartTwo(data []string) int {
	return 0
}
