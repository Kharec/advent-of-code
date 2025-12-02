package daynine

import (
	"os"
	"strconv"
	"strings"

	"advent-of-code/internal/registry"
)

func init() {
	registry.Register("2015D9", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	distances := make(map[string]map[string]int)

	for _, line := range data {
		parts := strings.Split(line, " = ")
		distance, _ := strconv.Atoi(parts[1])
		route := strings.Split(parts[0], " to ")
		from, to := route[0], route[1]

		if distances[from] == nil {
			distances[from] = make(map[string]int)
		}
		if distances[to] == nil {
			distances[to] = make(map[string]int)
		}

		distances[from][to] = distance
		distances[to][from] = distance
	}

	cities := make([]string, 0, len(distances))
	for city := range distances {
		cities = append(cities, city)
	}

	var generatePermutations func([]string) [][]string
	generatePermutations = func(cities []string) [][]string {
		if len(cities) == 0 {
			return [][]string{{}}
		}
		var result [][]string
		for idx, city := range cities {
			remaining := make([]string, len(cities)-1)
			copy(remaining[:idx], cities[:idx])
			copy(remaining[idx:], cities[idx+1:])

			for _, permutations := range generatePermutations(remaining) {
				result = append(result, append([]string{city}, permutations...))
			}
		}
		return result
	}

	minimalDistance := int(^uint(0) >> 1)
	for _, route := range generatePermutations(cities) {
		total := 0
		for idx := 0; idx < len(route)-1; idx++ {
			total += distances[route[idx]][route[idx+1]]
		}
		if total < minimalDistance {
			minimalDistance = total
		}
	}

	return minimalDistance
}

func PartTwo(data []string) int {
	return 0
}
