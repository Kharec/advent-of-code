package dayfour

import (
	"advent-of-code/internal/registry"
	"os"
	"strings"
)

func init() {
	registry.Register("2017D4", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	validCount := 0
	for _, passphrase := range data {
		words := strings.Fields(passphrase)
		seen := make(map[string]bool, len(words))
		hasDuplicate := false
		for _, word := range words {
			if seen[word] {
				hasDuplicate = true
				break
			}
			seen[word] = true
		}
		if !hasDuplicate {
			validCount++
		}
	}
	return validCount
}

func PartTwo(data []string) int {
	return 0
}
