package dayseven

import (
	"advent-of-code/internal/registry"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	rulePattern      = regexp.MustCompile(`^(\w+ \w+) bags contain (.+)\.$`)
	containedPattern = regexp.MustCompile(`(\d+) (\w+ \w+) bags?`)
)

type bagCount struct {
	count int
	bag   string
}

func init() {
	registry.Register("2020D7", ParseInput, PartOne, PartTwo)
}

func buildForwardGraph(data []string) map[string][]bagCount {
	forwardGraph := make(map[string][]bagCount)

	for _, line := range data {
		matches := rulePattern.FindStringSubmatch(line)
		if len(matches) != 3 {
			continue
		}

		containerBag := matches[1]
		containedString := matches[2]

		if containedString == "no other bags" {
			forwardGraph[containerBag] = []bagCount{}
			continue
		}

		containedMatches := containedPattern.FindAllStringSubmatch(containedString, -1)
		for _, match := range containedMatches {
			if len(match) >= 3 {
				count, _ := strconv.Atoi(match[1])
				containedBag := match[2]
				forwardGraph[containerBag] = append(forwardGraph[containerBag], bagCount{count: count, bag: containedBag})
			}
		}
	}

	return forwardGraph
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	forwardGraph := buildForwardGraph(data)
	reverseGraph := make(map[string][]string)

	for container, contained := range forwardGraph {
		for _, child := range contained {
			reverseGraph[child.bag] = append(reverseGraph[child.bag], container)
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
