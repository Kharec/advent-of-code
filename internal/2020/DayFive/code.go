package dayfive

import (
	"advent-of-code/internal/registry"
	"os"
	"strings"
)

func init() {
	registry.Register("2020D5", ParseInput, PartOne, PartTwo)
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

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	maxSeatID := 0
	for _, pass := range data {
		seatID := calculateSeatID(pass)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	return maxSeatID
}

func PartTwo(data []string) int {
	seatIDs := make(map[int]bool)
	minSeatID := 1000
	maxSeatID := 0

	for _, pass := range data {
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
