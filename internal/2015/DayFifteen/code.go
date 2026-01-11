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

func calculateCalories(ingredients []ingredient, amounts []int) int {
	var calories int
	for idx, ing := range ingredients {
		calories += ing.calories * amounts[idx]
	}
	return calories
}

func PartOne(data []string) int {
	var ingredients []ingredient
	for _, line := range data {
		ingredients = append(ingredients, parseIngredient(line))
	}
	amounts := make([]int, len(ingredients))
	var findMaxScore func(remaining int, index int) int
	findMaxScore = func(remaining int, index int) int {
		if index == len(ingredients)-1 {
			amounts[index] = remaining
			return calculateScore(ingredients, amounts)
		}
		maxScore := 0
		for idx := 0; idx <= remaining; idx++ {
			amounts[index] = idx
			score := findMaxScore(remaining-idx, index+1)
			if score > maxScore {
				maxScore = score
			}
		}
		return maxScore
	}
	return findMaxScore(100, 0)
}

func PartTwo(data []string) int {
	var ingredients []ingredient
	for _, line := range data {
		ingredients = append(ingredients, parseIngredient(line))
	}
	amounts := make([]int, len(ingredients))
	var findMaxScoreWithCalories func(remaining int, index int) int
	findMaxScoreWithCalories = func(remaining int, index int) int {
		if index == len(ingredients)-1 {
			amounts[index] = remaining
			if calculateCalories(ingredients, amounts) == 500 {
				return calculateScore(ingredients, amounts)
			}
			return 0
		}
		maxScore := 0
		for idx := 0; idx <= remaining; idx++ {
			amounts[index] = idx
			score := findMaxScoreWithCalories(remaining-idx, index+1)
			if score > maxScore {
				maxScore = score
			}
		}
		return maxScore
	}
	return findMaxScoreWithCalories(100, 0)
}
