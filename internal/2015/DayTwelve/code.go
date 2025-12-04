package daytwelve

import (
	"advent-of-code/internal/registry"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2015D12", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	re := regexp.MustCompile(`-?\d+`)
	sum := 0
	for _, line := range data {
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			number, _ := strconv.Atoi(match)
			sum += number
		}
	}
	return sum
}

func PartTwo(data []string) int {
	return 0
}
