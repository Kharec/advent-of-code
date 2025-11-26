package dayone

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2021D1", ParseInput, PartOne, PartTwo)
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
	count := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			count++
		}
	}
	return count
}

func PartTwo(data []int) int {
	count := 0
	for i := 3; i < len(data); i++ {
		if data[i] > data[i-3] {
			count++
		}
	}
	return count
}

