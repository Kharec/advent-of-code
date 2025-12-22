package daytwo

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2017D2", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) [][]int {
	content, _ := os.ReadFile(filepath)
	lines := strings.Split(string(content), "\n")
	result := make([][]int, 0, len(lines))

	for _, line := range lines {
		fields := strings.Fields(line)
		row := make([]int, 0, len(fields))
		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			row = append(row, num)
		}
		result = append(result, row)
	}

	return result
}

func PartOne(data [][]int) int {
	sum := 0
	for _, row := range data {
		min, max := row[0], row[0]
		for _, value := range row[1:] {
			if value < min {
				min = value
			}
			if value > max {
				max = value
			}
		}
		sum += max - min
	}
	return sum
}

func PartTwo(data [][]int) int {
	sum := 0
	for _, row := range data {
		found := false
		for idx, first := range row {
			for _, second := range row[idx+1:] {
				max, min := first, second
				if first < second {
					max, min = second, first
				}
				if min != 0 && max%min == 0 {
					sum += max / min
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	return sum
}
