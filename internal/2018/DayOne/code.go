package dayone

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2018D1", ParseInput, PartOne, PartTwo)
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
	sum := 0
	for _, num := range data {
		sum += num
	}
	return sum
}

func PartTwo(data []int) int {
	seen := make(map[int]bool)
	sum := 0
	seen[0] = true
	for {
		for _, num := range data {
			sum += num
			if seen[sum] {
				return sum
			}
			seen[sum] = true
		}
	}
}
