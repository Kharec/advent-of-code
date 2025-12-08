package daytwo

import (
	"advent-of-code/internal/registry"
	"fmt"
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

func PartTwo(instructions []string) int {
	row, column := 2, 0

	keypad := [5][5]int{
		{0, 0, 1, 0, 0},
		{0, 2, 3, 4, 0},
		{5, 6, 7, 8, 9},
		{0, 10, 11, 12, 0},
		{0, 0, 13, 0, 0},
	}

	isValidPosition := func(r, c int) bool {
		return r >= 0 && r < 5 && c >= 0 && c < 5 && keypad[r][c] != 0
	}

	var codeBuilder strings.Builder

	for _, line := range instructions {
		for _, move := range line {
			newRow, newColumn := row, column
			switch move {
			case 'U':
				newRow--
			case 'D':
				newRow++
			case 'L':
				newColumn--
			case 'R':
				newColumn++
			}
			if isValidPosition(newRow, newColumn) {
				row, column = newRow, newColumn
			}
		}
		value := keypad[row][column]
		if value < 10 {
			codeBuilder.WriteByte(byte('0' + value))
		} else {
			codeBuilder.WriteByte(byte('A' + value - 10))
		}
	}

	code := codeBuilder.String()
	fmt.Println(code)
	return 0
}
