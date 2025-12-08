package daythree

import (
	"advent-of-code/internal/registry"
	"os"
	"strings"
)

func init() {
	registry.Register("2022D3", ParseInput, PartOne, PartTwo)
}

func buildItemSet(items string) map[rune]bool {
	itemSet := make(map[rune]bool)
	for _, item := range items {
		itemSet[item] = true
	}
	return itemSet
}

func findCommonItem(itemSet map[rune]bool, items string) rune {
	for _, item := range items {
		if itemSet[item] {
			return item
		}
	}
	return 0
}

func findIntersection(itemSet map[rune]bool, items string) map[rune]bool {
	intersection := make(map[rune]bool)
	for _, item := range items {
		if itemSet[item] {
			intersection[item] = true
		}
	}
	return intersection
}

func getItemPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	}
	if item >= 'A' && item <= 'Z' {
		return int(item - 'A' + 27)
	}
	return 0
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

		firstCompartmentItems := buildItemSet(firstCompartment)
		commonItem := findCommonItem(firstCompartmentItems, secondCompartment)
		totalPriority += getItemPriority(commonItem)
	}

	return totalPriority
}

func PartTwo(data []string) int {
	totalPriority := 0
	for i := 0; i < len(data); i += 3 {
		if i+2 >= len(data) {
			break
		}
		firstRucksack := data[i]
		secondRucksack := data[i+1]
		thirdRucksack := data[i+2]

		firstRucksackItems := buildItemSet(firstRucksack)
		commonInFirstTwo := findIntersection(firstRucksackItems, secondRucksack)
		badge := findCommonItem(commonInFirstTwo, thirdRucksack)
		totalPriority += getItemPriority(badge)
	}

	return totalPriority
}
