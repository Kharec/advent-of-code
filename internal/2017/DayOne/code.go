package dayone

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
)

func init() {
	registry.Register("2017D1", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []int {
	content, _ := os.ReadFile(filepath)
	digits := string(content)
	result := make([]int, 0, len(digits))
	for _, char := range digits {
		digit, _ := strconv.Atoi(string(char))
		result = append(result, digit)
	}
	return result
}

func PartOne(data []int) int {
	sum := 0
	for idx := range data {
		nextIndex := (idx + 1) % len(data)
		if data[idx] == data[nextIndex] {
			sum += data[idx]
		}
	}
	return sum
}

func PartTwo(data []int) int {
	return 0
}
