package daytwo

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2020D2", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	lines := strings.Split(string(content), "\n")
	data := make([]string, 0, len(lines))
	return append(data, lines...)
}

func PartOne(data []string) int {
	valid := 0
	for _, line := range data {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ": ")
		policy := parts[0]
		password := parts[1]

		policyParts := strings.Split(policy, " ")
		rangeStr := policyParts[0]
		letter := policyParts[1]

		rangeParts := strings.Split(rangeStr, "-")
		min, _ := strconv.Atoi(rangeParts[0])
		max, _ := strconv.Atoi(rangeParts[1])

		count := strings.Count(password, letter)
		if count >= min && count <= max {
			valid++
		}
	}
	return valid
}

func PartTwo(data []string) int {
	valid := 0
	for _, line := range data {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ": ")
		policy := parts[0]
		password := parts[1]

		policyParts := strings.Split(policy, " ")
		rangeStr := policyParts[0]
		letter := policyParts[1]

		rangeParts := strings.Split(rangeStr, "-")
		firstPosition, _ := strconv.Atoi(rangeParts[0])
		secondPosition, _ := strconv.Atoi(rangeParts[1])

		atFirstPosition := len(password) >= firstPosition && password[firstPosition-1] == letter[0]
		atSecondPosition := len(password) >= secondPosition && password[secondPosition-1] == letter[0]

		if atFirstPosition != atSecondPosition {
			valid++
		}
	}
	return valid
}
