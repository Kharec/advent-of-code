package daytwo

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2015D2", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	total := 0
	for _, line := range data {
		parts := strings.Split(line, "x")
		length, _ := strconv.Atoi(parts[0])
		width, _ := strconv.Atoi(parts[1])
		height, _ := strconv.Atoi(parts[2])
		total += 2*length*width + 2*width*height + 2*height*length + min(length*width, width*height, height*length)
	}
	return total
}

func PartTwo(data []string) int {
	total := 0
	for _, line := range data {
		parts := strings.Split(line, "x")
		length, _ := strconv.Atoi(parts[0])
		width, _ := strconv.Atoi(parts[1])
		height, _ := strconv.Atoi(parts[2])

		smallest := min(length, width, height)
		middle := length + width + height - smallest - max(length, width, height)

		perimeter := 2 * (smallest + middle)
		volume := length * width * height

		total += perimeter + volume
	}
	return total
}
