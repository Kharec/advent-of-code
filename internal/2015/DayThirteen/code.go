package daythirteen

import (
	"advent-of-code/internal/registry"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var inputPattern = regexp.MustCompile(`(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+)\.?`)

func init() {
	registry.Register("2015D13", ParseInput, PartOne, PartTwo)
}

func buildHappinessMap(data []string) map[string]map[string]int {
	happinessMap := make(map[string]map[string]int)

	for _, line := range data {
		matches := inputPattern.FindStringSubmatch(line)
		person := matches[1]
		operation := matches[2]
		value, _ := strconv.Atoi(matches[3])
		neighbor := matches[4]

		if happinessMap[person] == nil {
			happinessMap[person] = make(map[string]int)
		}

		if operation == "gain" {
			happinessMap[person][neighbor] = value
		} else {
			happinessMap[person][neighbor] = -value
		}
	}

	return happinessMap
}

func getAllPeople(happinessMap map[string]map[string]int) []string {
	people := make([]string, 0, len(happinessMap))
	for person := range happinessMap {
		people = append(people, person)
	}
	return people
}

func generatePermutations(items []string) [][]string {
	if len(items) == 0 {
		return [][]string{{}}
	}

	var result [][]string
	for idx, item := range items {
		remaining := make([]string, len(items)-1)
		copy(remaining[:idx], items[:idx])
		copy(remaining[idx:], items[idx+1:])

		for _, permutation := range generatePermutations(remaining) {
			result = append(result, append([]string{item}, permutation...))
		}
	}
	return result
}

func calculateTotalHappiness(arrangement []string, happinessMap map[string]map[string]int) int {
	totalHappiness := 0
	arrangementLength := len(arrangement)

	for idx := range arrangementLength {
		currentPerson := arrangement[idx]
		leftNeighbor := arrangement[(idx-1+arrangementLength)%arrangementLength]
		rightNeighbor := arrangement[(idx+1)%arrangementLength]

		totalHappiness += happinessMap[currentPerson][leftNeighbor]
		totalHappiness += happinessMap[currentPerson][rightNeighbor]
	}

	return totalHappiness
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	happinessMap := buildHappinessMap(data)
	allPeople := getAllPeople(happinessMap)

	fixedPerson := allPeople[0]
	remainingPeople := allPeople[1:]
	permutations := generatePermutations(remainingPeople)

	maxHappiness := math.MinInt
	arrangement := make([]string, len(allPeople))
	arrangement[0] = fixedPerson

	for _, perm := range permutations {
		copy(arrangement[1:], perm)
		totalHappiness := calculateTotalHappiness(arrangement, happinessMap)
		if totalHappiness > maxHappiness {
			maxHappiness = totalHappiness
		}
	}

	return maxHappiness
}

func PartTwo(data []string) int {
	happinessMap := buildHappinessMap(data)
	allPeople := getAllPeople(happinessMap)

	me := "Me"
	happinessMap[me] = make(map[string]int)
	for _, person := range allPeople {
		happinessMap[person][me] = 0
		happinessMap[me][person] = 0
	}

	allPeople = append(allPeople, me)
	fixedPerson := allPeople[0]
	remainingPeople := allPeople[1:]
	permutations := generatePermutations(remainingPeople)

	maxTotalChange := math.MinInt
	arrangement := make([]string, len(allPeople))
	arrangement[0] = fixedPerson

	for _, permutation := range permutations {
		copy(arrangement[1:], permutation)
		totalHappiness := calculateTotalHappiness(arrangement, happinessMap)
		if totalHappiness > maxTotalChange {
			maxTotalChange = totalHappiness
		}
	}

	return maxTotalChange
}
