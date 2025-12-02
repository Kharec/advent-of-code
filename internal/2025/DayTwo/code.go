package daytwo

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2025D2", ParseInput, PartOne, PartTwo)
}

func isInvalidID(id int) bool {
	sID := strconv.Itoa(id)
	if len(sID)%2 != 0 {
		return false
	}
	half := len(sID) / 2
	return sID[:half] == sID[half:]
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
				if isInvalidID(id) {
					sum += id
				}
			}
		}
	}
	return sum
}

func PartTwo(input []string) int {
	return 0
}
