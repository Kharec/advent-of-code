package dayfive

import (
	"advent-of-code/internal/registry"
	"os"
	"unicode"
)

func init() {
	registry.Register("2018D5", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) string {
	content, _ := os.ReadFile(filepath)
	return string(content)
}

func reactPolymer(data string) int {
	stack := []rune{}

	for _, char := range data {
		if len(stack) == 0 {
			stack = append(stack, char)
			continue
		}

		top := stack[len(stack)-1]

		if unicode.ToLower(top) == unicode.ToLower(char) && top != char {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack)
}

func PartOne(data string) int {
	return reactPolymer(data)
}

func PartTwo(data string) int {
	unitTypes := make(map[rune]bool)
	for _, char := range data {
		unitTypes[unicode.ToLower(char)] = true
	}

	shortestPolymerLength := len(data)
	for unitType := range unitTypes {
		filtered := make([]rune, 0, len(data))
		for _, char := range data {
			if unicode.ToLower(char) != unitType {
				filtered = append(filtered, char)
			}
		}
		length := reactPolymer(string(filtered))
		if length < shortestPolymerLength {
			shortestPolymerLength = length
		}
	}

	return shortestPolymerLength
}
