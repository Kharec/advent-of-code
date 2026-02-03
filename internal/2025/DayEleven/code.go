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
	graph := make(map[string][]string)

	for _, line := range data {
		device, outputsStr, _ := strings.Cut(line, ": ")
		graph[device] = strings.Fields(outputsStr)
	}

	var countPaths func(string) int
	countPaths = func(node string) int {
		if node == "out" {
			return 1
		}

		count := 0
		for _, neighbor := range graph[node] {
			count += countPaths(neighbor)
		}
		return count
	}

	return countPaths("you")
}

func PartTwo(data []string) int {
	return 0
}
