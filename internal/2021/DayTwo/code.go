package daytwo

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2021D2", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	horizontal := 0
	depth := 0
	for _, line := range data {
		parts := strings.Split(line, " ")
		direction := parts[0]
		amount, _ := strconv.Atoi(parts[1])
		switch direction {
		case "forward":
			horizontal += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		}
	}
	return horizontal * depth
}

func PartTwo(data []string) int {
	horizontal := 0
	depth := 0
	aim := 0

	for _, line := range data {
		parts := strings.Split(line, " ")
		direction := parts[0]
		amount, _ := strconv.Atoi(parts[1])

		switch direction {
		case "forward":
			horizontal += amount
			depth += aim * amount
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}
	return horizontal * depth
}
