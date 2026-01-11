package dayfifteen

import (
	"advent-of-code/internal/registry"
	"fmt"
	"os"
	"strings"
)

func init() {
	registry.Register("2015D15", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

type ingredient struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func parseIngredient(line string) ingredient {
	parts := strings.Split(line, ": ")
	props := strings.Split(parts[1], ", ")
	var cap, dur, flav, tex, cal int

	for _, prop := range props {
		var name string
		var value int
		fmt.Sscanf(prop, "%s %d", &name, &value)
		switch name {
		case "capacity":
			cap = value
		case "durability":
			dur = value
		case "flavor":
			flav = value
		case "texture":
			tex = value
		case "calories":
			cal = value
		}
	}

	return ingredient{
		capacity:   cap,
		durability: dur,
		flavor:     flav,
		texture:    tex,
		calories:   cal,
	}
}

func calculateScore(ingredients []ingredient, amounts []int) int {
	var capacity, durability, flavor, texture int
	for idx, ing := range ingredients {
		capacity += ing.capacity * amounts[idx]
		durability += ing.durability * amounts[idx]
		flavor += ing.flavor * amounts[idx]
		texture += ing.texture * amounts[idx]
	}
	if capacity < 0 {
		capacity = 0
	}
	if durability < 0 {
		durability = 0
	}
	if flavor < 0 {
		flavor = 0
	}
	if texture < 0 {
		texture = 0
	}
	return capacity * durability * flavor * texture
}

func findMaxScore(ingredients []ingredient, remaining int, amounts []int, index int) int {
	if index == len(ingredients)-1 {
		amounts[index] = remaining
		return calculateScore(ingredients, amounts)
	}
	maxScore := 0
	for i := 0; i <= remaining; i++ {
		amounts[index] = i
		score := findMaxScore(ingredients, remaining-i, amounts, index+1)
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

func PartOne(data []string) int {
	var ingredients []ingredient
	for _, line := range data {
		ingredients = append(ingredients, parseIngredient(line))
	}
	amounts := make([]int, len(ingredients))
	return findMaxScore(ingredients, 100, amounts, 0)
}

func PartTwo(data []string) int {
	return 0
}
