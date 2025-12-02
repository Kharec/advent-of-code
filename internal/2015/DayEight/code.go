package dayeight

import (
	"os"
	"strings"

	"advent-of-code/internal/registry"
)

func init() {
	registry.Register("2015D8", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}
