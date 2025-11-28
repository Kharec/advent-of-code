package dayseven

import (
	"advent-of-code/internal/registry"
	"os"
	"regexp"
	"strings"
)

func init() {
	registry.Register("2020D7", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	reverseGraph := make(map[string][]string)
	rulePattern := regexp.MustCompile(`^(\w+ \w+) bags contain (.+)\.$`)
	containedPattern := regexp.MustCompile(`(\d+) (\w+ \w+) bags?`)

	for _, line := range data {
		matches := rulePattern.FindStringSubmatch(line)
		if len(matches) != 3 {
			continue
		}

		containerBag := matches[1]
		containedString := matches[2]

		if containedString == "no other bags" {
			continue
		}

		containedMatches := containedPattern.FindAllStringSubmatch(containedString, -1)
		for _, match := range containedMatches {
			if len(match) >= 3 {
				containedBag := match[2]
				reverseGraph[containedBag] = append(reverseGraph[containedBag], containerBag)
			}
		}
	}

	visited := make(map[string]bool)
	queue := reverseGraph["shiny gold"]

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}
		visited[current] = true

		for _, parent := range reverseGraph[current] {
			if !visited[parent] {
				queue = append(queue, parent)
			}
		}
	}

	return len(visited)
}

func PartTwo(data []string) int {
	return 0
}
