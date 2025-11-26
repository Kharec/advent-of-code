package dayone

import (
	"advent-of-code/internal/registry"
	"os"
	"slices"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2022D1", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(input []string) int {
	maxCalories, currentElf := 0, 0

	for _, line := range input {
		if line == "" {
			maxCalories = max(maxCalories, currentElf)
			currentElf = 0
			continue
		}
		val, _ := strconv.Atoi(line)
		currentElf += val
	}

	return max(maxCalories, currentElf)
}

func PartTwo(input []string) int {
	var totals []int
	currentElf := 0

	for _, line := range input {
		if line == "" {
			totals = append(totals, currentElf)
			currentElf = 0
			continue
		}

		val, _ := strconv.Atoi(line)
		currentElf += val
	}
	totals = append(totals, currentElf)

	slices.Sort(totals)
	slices.Reverse(totals)

	sum := 0
	for _, val := range totals[:min(3, len(totals))] {
		sum += val
	}
	return sum
}
