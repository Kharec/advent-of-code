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
	return strings.Split(string(content), "\n")
}

func PartOne(input []string) int {
	maxSeatID := 0
	for _, pass := range input {
		rowStr := pass[:7]
		columnStr := pass[7:10]

		row := 0
		for idx, char := range rowStr {
			if char == 'B' {
				row |= 1 << (6 - idx)
			}
		}

		column := 0
		for idx, char := range columnStr {
			if char == 'R' {
				column |= 1 << (2 - idx)
			}
		}

		seatID := row*8 + column
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	return maxSeatID
}

func main() {
	input := parseInput("input.txt")
	fmt.Println("Part 1:", PartOne(input))
}
