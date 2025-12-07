package dayone

import (
	"os"
	"strconv"
	"strings"

	"advent-of-code/internal/registry"
)

func init() {
	registry.Register("2016D1", ParseInput, PartOne, PartTwo)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	parts := strings.Split(string(content), ",")
	for i, p := range parts {
		parts[i] = strings.TrimSpace(p)
	}
	return parts
}

func PartOne(data []string) int {
	x, y := 0, 0
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	currentDirection := 0

	for _, instruction := range data {
		turn := instruction[0]
		distance, _ := strconv.Atoi(instruction[1:])

		switch turn {
		case 'R':
			currentDirection = (currentDirection + 1) % 4
		case 'L':
			currentDirection = (currentDirection + 3) % 4
		}

		x += directions[currentDirection][0] * distance
		y += directions[currentDirection][1] * distance
	}

	return abs(x) + abs(y)
}

func PartTwo(data []string) int {
	return 0
}
