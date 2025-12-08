package dayseven

import (
	"os"
	"strings"

	"advent-of-code/internal/registry"
)

func init() {
	registry.Register("2025D7", ParseInput, PartOne, PartTwo)
}

type Position struct {
	row    int
	column int
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func findStart(data []string) (int, int) {
	for row := range data {
		if column := strings.IndexByte(data[row], 'S'); column != -1 {
			return row, column
		}
	}
	return -1, -1
}

func PartOne(data []string) int {
	rows, columns := len(data), len(data[0])
	startRow, startColumn := findStart(data)

	queue := []Position{{startRow, startColumn}}
	activated := make(map[Position]bool)
	splits := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		row := current.row + 1
		column := current.column

		for row < rows && data[row][column] != '^' {
			row++
		}

		if row < rows {
			splitPosition := Position{row, column}
			if !activated[splitPosition] {
				activated[splitPosition] = true
				splits++

				if column > 0 {
					queue = append(queue, Position{row, column - 1})
				}
				if column < columns-1 {
					queue = append(queue, Position{row, column + 1})
				}
			}
		}
	}

	return splits
}

func PartTwo(input []string) int {
	rows, columns := len(input), len(input[0])
	startRow, startColumn := findStart(input)

	cache := make(map[Position]int)

	var countTimelines func(row, column int) int
	countTimelines = func(row, column int) int {
		position := Position{row, column}
		if count, exists := cache[position]; exists {
			return count
		}

		r, c := row+1, column
		for r < rows && input[r][c] != '^' {
			r++
		}

		if r == rows {
			return 1
		}

		total := 0
		if c > 0 {
			total += countTimelines(r, c-1)
		}
		if c < columns-1 {
			total += countTimelines(r, c+1)
		}

		cache[position] = total
		return total
	}

	return countTimelines(startRow, startColumn)
}
