package dayeleven

import (
	"os"
	"strings"

	"advent-of-code/internal/registry"
)

func init() {
	registry.Register("2025D11", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	return 0
}

func PartTwo(data []string) int {
	return 0
}
