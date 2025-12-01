package daytwo

import (
	"advent-of-code/internal/registry"
	"os"
	"strings"
)

func init() {
	registry.Register("2022D2", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(strings.Trim(string(content), "\n"), "\n")
}

func PartOne(data []string) int {
	totalScore := 0
	for _, line := range data {
		opponent := line[0]
		me := line[2]

		shapeScore := int(me - 'X' + 1)

		var outcomeScore int
		if (opponent == 'A' && me == 'Y') || (opponent == 'B' && me == 'Z') || (opponent == 'C' && me == 'X') {
			outcomeScore = 6
		} else if (opponent == 'A' && me == 'X') || (opponent == 'B' && me == 'Y') || (opponent == 'C' && me == 'Z') {
			outcomeScore = 3
		} else {
			outcomeScore = 0
		}

		totalScore += shapeScore + outcomeScore
	}
	return totalScore
}

func PartTwo(data []string) int {
	totalScore := 0
	for _, line := range data {
		opponent := line[0]
		outcome := line[2]

		var me byte
		switch outcome {
		case 'Y':
			me = 'X' + (opponent - 'A')
		case 'Z':
			me = 'X' + ((opponent-'A')+1)%3
		default:
			me = 'X' + ((opponent-'A')+2)%3
		}

		shapeScore := int(me - 'X' + 1)
		outcomeScore := int(outcome-'X') * 3

		totalScore += shapeScore + outcomeScore
	}
	return totalScore
}
