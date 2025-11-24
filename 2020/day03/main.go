package main

import (
	"fmt"
	"log"
	"os"
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

func PartOne(input []string) int {
	trees := 0
	column := 0
	for row := range input {
		if len(input[row]) == 0 {
			continue
		}
		if input[row][column%len(input[row])] == '#' {
			trees++
		}
		column += 3
	}
	return trees
}

func PartTwo(input []string) int {
	result := 1
	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, slope := range slopes {
		trees := 0
		column := 0
		for row := 0; row < len(input); row += slope[1] {
			if len(input[row]) == 0 {
				continue
			}
			if input[row][column%len(input[row])] == '#' {
				trees++
			}
			column += slope[0]
		}
		result *= trees
	}
	return result
}

func main() {
	data := parseInput("input.txt")
	fmt.Println("Part 1:", PartOne(data))
	fmt.Println("Part 2:", PartTwo(data))
}
