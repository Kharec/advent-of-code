package dayfour

import (
	"advent-of-code/internal/registry"
	"fmt"
	"os"
	"strings"
)

func init() {
	registry.Register("2022D4", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	count := 0
	for _, line := range data {
		parts := strings.Split(line, ",")

		var firstRangeStart, firstRangeEnd, secondRangeStart, secondRangeEnd int
		fmt.Sscanf(parts[0], "%d-%d", &firstRangeStart, &firstRangeEnd)
		fmt.Sscanf(parts[1], "%d-%d", &secondRangeStart, &secondRangeEnd)

		minStart := min(firstRangeStart, secondRangeStart)
		maxEnd := max(firstRangeEnd, secondRangeEnd)
		if (minStart == firstRangeStart && maxEnd == firstRangeEnd) || (minStart == secondRangeStart && maxEnd == secondRangeEnd) {
			count++
		}
	}

	return count
}

func PartTwo(data []string) int {
	return 0
}
