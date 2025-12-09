package daythree

import (
	"advent-of-code/internal/registry"
	"fmt"
	"os"
	"strings"
)

func init() {
	registry.Register("2016D3", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) [][3]int {
	content, _ := os.ReadFile(filepath)
	lines := strings.Split(string(content), "\n")
	var result [][3]int
	for _, line := range lines {
		var a, b, c int
		fmt.Sscanf(line, "%d %d %d", &a, &b, &c)
		result = append(result, [3]int{a, b, c})

	}
	return result
}

func PartOne(data [][3]int) int {
	count := 0
	for _, triangle := range data {
		a, b, c := triangle[0], triangle[1], triangle[2]
		if a+b > c && a+c > b && b+c > a {
			count++
		}
	}
	return count
}

func PartTwo(data [][3]int) int {
	return 0
}
