package daythree

import (
	"advent-of-code/internal/registry"
	"os"
	"strings"
)

func init() {
	registry.Register("2022D3", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	totalPriority := 0
	for _, rucksack := range data {
		compartmentSize := len(rucksack) / 2
		firstCompartment := rucksack[:compartmentSize]
		secondCompartment := rucksack[compartmentSize:]

		firstCompartmentItems := make(map[rune]bool)
		for _, item := range firstCompartment {
			firstCompartmentItems[item] = true
		}

		var commonItem rune
		for _, item := range secondCompartment {
			if firstCompartmentItems[item] {
				commonItem = item
				break
			}
		}

		if commonItem >= 'a' && commonItem <= 'z' {
			totalPriority += int(commonItem - 'a' + 1)
		} else if commonItem >= 'A' && commonItem <= 'Z' {
			totalPriority += int(commonItem - 'A' + 27)
		}
	}

	return totalPriority
}

func PartTwo(data []string) int {
	return 0
}
