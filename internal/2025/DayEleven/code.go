package dayeleven

import (
	"os"
	"slices"
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

func buildGraph(data []string) map[string][]string {
	graph := make(map[string][]string)

	for _, line := range data {
		device, outputsStr, _ := strings.Cut(line, ": ")
		graph[device] = strings.Fields(outputsStr)
	}

	return graph
}

func hasCycle(graph map[string][]string) bool {
	const (
		unvisited = iota
		visiting
		visited
	)

	state := make(map[string]uint8)

	var dfs func(node string) bool
	dfs = func(node string) bool {
		switch state[node] {
		case visiting:
			return true
		case visited:
			return false
		}

		state[node] = visiting
		if slices.ContainsFunc(graph[node], dfs) {
			return true
		}
		state[node] = visited
		return false
	}

	for node := range graph {
		if state[node] == unvisited {
			if dfs(node) {
				return true
			}
		}
	}

	return false
}

func countPaths(graph map[string][]string, start, end string, requiredNodes map[string]bool) int {
	requiredIndex := make(map[string]uint)
	var targetMask uint
	for node := range requiredNodes {
		requiredIndex[node] = uint(len(requiredIndex))
		targetMask |= 1 << requiredIndex[node]
	}

	if hasCycle(graph) {
		var dfs func(node string, visited map[string]bool, seenMask uint) int
		dfs = func(node string, visited map[string]bool, seenMask uint) int {
			if node == end {
				if seenMask != targetMask {
					return 0
				}
				return 1
			}

			if idx, ok := requiredIndex[node]; ok {
				seenMask |= 1 << idx
			}

			visited[node] = true
			defer delete(visited, node)

			count := 0
			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					count += dfs(neighbor, visited, seenMask)
				}
			}
			return count
		}

		return dfs(start, make(map[string]bool), 0)
	}

	type key struct {
		node string
		mask uint
	}
	memo := make(map[key]int)

	var dfs func(node string, seenMask uint) int
	dfs = func(node string, seenMask uint) int {
		if node == end {
			if seenMask != targetMask {
				return 0
			}
			return 1
		}

		if idx, ok := requiredIndex[node]; ok {
			seenMask |= 1 << idx
		}

		state := key{node: node, mask: seenMask}
		if cached, ok := memo[state]; ok {
			return cached
		}

		count := 0
		for _, neighbor := range graph[node] {
			count += dfs(neighbor, seenMask)
		}

		memo[state] = count
		return count
	}

	return dfs(start, 0)
}

func PartOne(data []string) int {
	graph := buildGraph(data)
	return countPaths(graph, "you", "out", nil)
}

func PartTwo(data []string) int {
	graph := buildGraph(data)
	requiredNodes := map[string]bool{
		"dac": true,
		"fft": true,
	}
	return countPaths(graph, "svr", "out", requiredNodes)
}
