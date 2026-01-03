package dayfive

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2017D5", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []int {
	content, _ := os.ReadFile(filepath)
	var data []int
	for line := range strings.SplitSeq(string(content), "\n") {
		num, _ := strconv.Atoi(line)
		data = append(data, num)
	}
	return data
}

func PartOne(data []int) int {
	instructions := make([]int, len(data))
	copy(instructions, data)

	position := 0
	steps := 0

	for position >= 0 && position < len(instructions) {
		jump := instructions[position]
		instructions[position]++
		position += jump
		steps++
	}

	return steps
}

func PartTwo(data []int) int {
	return 0
}
