package daythree

import (
	"os"
	"strings"

	"advent-of-code/internal/registry"
)

func init() {
	registry.Register("2025D3", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) [][]int {
	content, _ := os.ReadFile(filepath)
	lines := strings.Split(string(content), "\n")
	var data [][]int
	for _, line := range lines {
		var bank []int
		for _, char := range line {
			bank = append(bank, int(char-'0'))
		}
		data = append(data, bank)
	}
	return data
}

func PartOne(data [][]int) int {
	total := 0
	for _, bank := range data {
		maxAfter := make([]int, len(bank))

		for idx := len(bank) - 2; idx >= 0; idx-- {
			maxAfter[idx] = max(maxAfter[idx+1], bank[idx+1])
		}

		maxJoltage := 0

		for idx := 0; idx < len(bank)-1; idx++ {
			joltage := bank[idx]*10 + maxAfter[idx]
			maxJoltage = max(maxJoltage, joltage)
		}
		total += maxJoltage
	}
	return total
}

func PartTwo(data [][]int) int {
	total := 0
	for _, bank := range data {
		var joltage int
		start := 0

		for position := range 12 {
			remaining := 12 - position - 1
			maxIdx := start
			maxValue := bank[start]

			for idx := start; idx < len(bank)-remaining; idx++ {
				if bank[idx] > maxValue {
					maxValue = bank[idx]
					maxIdx = idx
				}
			}

			joltage = joltage*10 + maxValue
			start = maxIdx + 1
		}
		total += joltage
	}
	return total
}
