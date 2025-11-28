package daythree

import (
	"advent-of-code/internal/registry"
	"os"
)

type coordinates struct {
	x, y int
}

func init() {
	registry.Register("2015D3", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) string {
	content, _ := os.ReadFile(filepath)
	return string(content)
}

func PartOne(data string) int {
	houses := make(map[coordinates]int)
	x, y := 0, 0
	houses[coordinates{x, y}] = 1
	for _, direction := range data {
		switch direction {
		case '>':
			x++
		case '<':
			x--
		case '^':
			y++
		case 'v':
			y--
		}
		houses[coordinates{x, y}]++
	}
	return len(houses)
}

func PartTwo(data string) int {
	return 0
}
