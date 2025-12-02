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

func buildDistanceMap(data []string) map[string]map[string]int {
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

	return distances
}

func getCities(distances map[string]map[string]int) []string {
	cities := make([]string, 0, len(distances))
	for city := range distances {
		cities = append(cities, city)
	}
	return cities
}

func generatePermutations(cities []string) [][]string {
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

func calculateRouteDistance(route []string, distances map[string]map[string]int) int {
	total := 0
	for idx := 0; idx < len(route)-1; idx++ {
		total += distances[route[idx]][route[idx+1]]
	}
	return total
}

func PartOne(data []string) int {
	distances := buildDistanceMap(data)
	cities := getCities(distances)
	permutations := generatePermutations(cities)

	minimalDistance := int(^uint(0) >> 1)
	for _, route := range permutations {
		total := calculateRouteDistance(route, distances)
		if total < minimalDistance {
			minimalDistance = total
		}
	}

	return minimalDistance
}

func PartTwo(data []string) int {
	distances := buildDistanceMap(data)
	cities := getCities(distances)
	permutations := generatePermutations(cities)

	maximalDistance := 0
	for _, route := range permutations {
		total := calculateRouteDistance(route, distances)
		if total > maximalDistance {
			maximalDistance = total
		}
	}

	return maximalDistance
}
