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

func main() {
	data := parseInput("input.txt")
	fmt.Println("Part 1:", PartOne(data))
	fmt.Println("Part 2:", PartTwo(data))
}
