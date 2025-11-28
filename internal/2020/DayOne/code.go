package dayone

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2020D1", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []int {
	content, _ := os.ReadFile(filepath)
	lines := strings.Fields(string(content))
	var data []int
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		data = append(data, num)
	}
	return data
}

func PartOne(data []int) int {
	seen := make(map[int]bool)
	for _, number := range data {
		complement := 2020 - number
		if seen[complement] {
			return number * complement
		}
		seen[number] = true
	}
	return 0
}

func PartTwo(data []int) int {
	seen := make(map[int]bool)
	for idx := range data {
		for _, number := range data[idx+1:] {
			complement := 2020 - data[idx] - number
			if seen[complement] {
				return complement * data[idx] * number
			}
		}
		seen[data[idx]] = true
	}
	return 0
}
