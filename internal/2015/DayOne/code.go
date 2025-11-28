package dayone

import (
	"advent-of-code/internal/registry"
	"os"
)

func init() {
	registry.Register("2015D1", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) string {
	content, _ := os.ReadFile(filepath)
	return string(content)
}

func PartOne(data string) int {
	floor := 0
	for _, char := range data {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return floor
}

func PartTwo(data string) int {
	floor := 0
	for idx, char := range data {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			return idx + 1
		}
	}
	return 0
}
