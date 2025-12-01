package dayone

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2025D1", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	position := 50
	count := 0

	for _, rotation := range data {
		direction := rotation[0]
		distance, _ := strconv.Atoi(rotation[1:])

		if direction == 'L' {
			position = (position - distance + 100) % 100
		} else {
			position = (position + distance) % 100
		}

		if position == 0 {
			count++
		}
	}

	return count
}

func PartTwo(data []string) int {
	position := 50
	count := 0

	for _, rotation := range data {
		direction := rotation[0]
		distance, _ := strconv.Atoi(rotation[1:])

		for range distance {
			if direction == 'L' {
				position = (position - 1 + 100) % 100
			} else {
				position = (position + 1) % 100
			}

			if position == 0 {
				count++
			}
		}
	}

	return count
}
