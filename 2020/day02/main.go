package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	lines := strings.Split(string(content), "\n")
	var data []string
	for _, line := range lines {
		data = append(data, line)
	}
	return data
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

func main() {
	data := parseInput("input.txt")
	fmt.Println("Part 1:", PartOne(data))
	fmt.Println("Part 2:", PartTwo(data))
}
