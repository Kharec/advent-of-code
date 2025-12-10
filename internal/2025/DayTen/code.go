package dayten

import (
	"advent-of-code/internal/registry"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2025D10", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func parseMachine(line string) ([]bool, [][]int) {
	bracketRegex := regexp.MustCompile(`\[([.#]+)\]`)
	parenthesisRegex := regexp.MustCompile(`\(([^)]+)\)`)

	bracketMatch := bracketRegex.FindStringSubmatch(line)
	targetString := bracketMatch[1]
	target := make([]bool, len(targetString))
	for idx, char := range targetString {
		target[idx] = char == '#'
	}

	parenthesisMatches := parenthesisRegex.FindAllStringSubmatch(line, -1)
	buttons := make([][]int, 0, len(parenthesisMatches))

	for _, match := range parenthesisMatches {
		buttonString := match[1]
		parts := strings.Split(buttonString, ",")
		button := make([]int, 0, len(parts))
		for _, part := range parts {
			light, _ := strconv.Atoi(part)
			button = append(button, light)
		}
		buttons = append(buttons, button)
	}

	return target, buttons
}

func PartOne(data []string) int {
	total := 0
	for _, line := range data {
		target, buttons := parseMachine(line)
		numberOfLights := len(target)
		numberOfButtons := len(buttons)
		minimumPresses := numberOfButtons + 1

		for mask := 0; mask < (1 << numberOfButtons); mask++ {
			lights := make([]bool, numberOfLights)
			presses := 0

			for buttonIdx := range numberOfButtons {
				if mask&(1<<buttonIdx) != 0 {
					presses++
					for _, light := range buttons[buttonIdx] {
						if light < numberOfLights {
							lights[light] = !lights[light]
						}
					}
				}
			}

			matches := true
			for idx := range numberOfLights {
				if lights[idx] != target[idx] {
					matches = false
					break
				}
			}

			if matches && presses < minimumPresses {
				minimumPresses = presses
			}
		}
		total += minimumPresses
	}

	return total
}

func PartTwo(data []string) int {
	return 0
}
