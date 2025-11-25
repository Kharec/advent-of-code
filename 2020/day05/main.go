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

func calculateSeatID(pass string) int {
	if len(pass) < 10 {
		return -1
	}
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

	return row*8 + column
}

func PartOne(input []string) int {
	maxSeatID := 0
	for _, pass := range input {
		seatID := calculateSeatID(pass)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	return maxSeatID
}

func PartTwo(input []string) int {
	seatIDs := make(map[int]bool)
	minSeatID := 1000
	maxSeatID := 0

	for _, pass := range input {
		if len(pass) < 10 {
			continue
		}
		seatID := calculateSeatID(pass)
		seatIDs[seatID] = true
		if seatID < minSeatID {
			minSeatID = seatID
		}
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	for seatID := minSeatID + 1; seatID < maxSeatID; seatID++ {
		if !seatIDs[seatID] && seatIDs[seatID-1] && seatIDs[seatID+1] {
			return seatID
		}
	}

	return 0
}

func main() {
	input := parseInput("input.txt")
	fmt.Println("Part 1:", PartOne(input))
	fmt.Println("Part 2:", PartTwo(input))
}
