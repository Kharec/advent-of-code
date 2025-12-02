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

func isInvalid(id int, exactTwo bool) bool {
	sID := strconv.Itoa(id)
	if exactTwo {
		half := len(sID) / 2
		return sID[:half] == sID[half:]
	}
	for patternLength := 1; patternLength <= len(sID)/2; patternLength++ {
		repetitions := len(sID) / patternLength
		if repetitions < 2 || len(sID)%patternLength != 0 {
			continue
		}
		if strings.Repeat(sID[:patternLength], repetitions) == sID {
			return true
		}
	}
	return false
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
				if isInvalid(id, true) {
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
