package dayfive

import (
	"advent-of-code/internal/registry"
	"os"
	"strings"
)

func init() {
	registry.Register("2015D5", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	count := 0
	for _, line := range data {
		vowelCount := strings.Count(line, "a") + strings.Count(line, "e") + strings.Count(line, "i") + strings.Count(line, "o") + strings.Count(line, "u")
		if vowelCount < 3 {
			continue
		}

		hasDoubleLetter := false
		for i := 1; i < len(line); i++ {
			if line[i] == line[i-1] {
				hasDoubleLetter = true
				break
			}
		}
		if !hasDoubleLetter {
			continue
		}

		if strings.Contains(line, "ab") || strings.Contains(line, "cd") || strings.Contains(line, "pq") || strings.Contains(line, "xy") {
			continue
		}

		count++
	}
	return count
}

func PartTwo(data []string) int {
	count := 0
	for _, line := range data {
		if len(line) < 4 {
			continue
		}

		hasNonOverlappingPair := false
		for i := 0; i < len(line)-1; i++ {
			pair := line[i : i+2]
			if strings.Contains(line[i+2:], pair) {
				hasNonOverlappingPair = true
				break
			}
		}
		if !hasNonOverlappingPair {
			continue
		}

		hasRepeatingLetterWithGap := false
		for i := 0; i < len(line)-2; i++ {
			if line[i] == line[i+2] {
				hasRepeatingLetterWithGap = true
				break
			}
		}
		if hasRepeatingLetterWithGap {
			count++
		}
	}
	return count
}
