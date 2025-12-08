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

func PartOne(data []string) int {
	sum := 0
	for _, line := range data {
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

func PartTwo(data []string) int {
	totalSum := 0
	for _, line := range data {
		for rangeSpec := range strings.SplitSeq(line, ",") {
			rangeStartStr, rangeEndStr, _ := strings.Cut(rangeSpec, "-")
			rangeStart, _ := strconv.Atoi(rangeStartStr)
			rangeEnd, _ := strconv.Atoi(rangeEndStr)

			seenNumbers := make(map[int]bool)
			maxDigits := len(rangeEndStr)

			var generatePatterns func(string, int, int)
			generatePatterns = func(currentPattern string, targetPatternLength, currentDepth int) {
				if currentDepth == targetPatternLength {
					repeated := currentPattern + currentPattern
					for len(repeated) <= maxDigits {
						number, err := strconv.Atoi(repeated)
						if err != nil || number > rangeEnd {
							break
						}
						if number >= rangeStart && !seenNumbers[number] {
							totalSum += number
							seenNumbers[number] = true
						}
						repeated += currentPattern
					}
					return
				}

				minDigit := 0
				if currentDepth == 0 {
					minDigit = 1
				}

				for digit := minDigit; digit <= 9; digit++ {
					extendedPattern := currentPattern + strconv.Itoa(digit)
					generatePatterns(extendedPattern, targetPatternLength, currentDepth+1)
				}
			}

			for patternLength := 1; patternLength <= maxDigits/2; patternLength++ {
				generatePatterns("", patternLength, 0)
			}
		}
	}
	return totalSum
}
