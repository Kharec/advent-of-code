package daytwo

import (
	"os"
	"strconv"
	"strings"

	"advent-of-code/internal/registry"
)

func init() {
	registry.Register("2025D2", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

func PartOne(input []string) int {
	sum := 0
	for _, line := range input {
		for newRange := range strings.SplitSeq(line, ",") {
			parts := strings.Split(newRange, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			for id := start; id <= end; id++ {
				sID := strconv.Itoa(id)
				if len(sID)%2 == 0 {
					half := len(sID) / 2
					if sID[:half] == sID[half:] {
						sum += id
					}
				}
			}
		}
	}
	return sum
}

func PartTwo(input []string) int {
	return 0
}
