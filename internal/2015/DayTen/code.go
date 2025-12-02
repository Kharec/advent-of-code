package dayten

import (
	"os"
	"strconv"
	"strings"

	"advent-of-code/internal/registry"
)

func init() {
	registry.Register("2015D10", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) string {
	content, _ := os.ReadFile(filepath)
	return string(content)
}

func applyLookAndSay(data string, iterations int) int {
	for range iterations {
		var result strings.Builder
		idx := 0
		for idx < len(data) {
			digit := data[idx]
			count := 1
			for idx+count < len(data) && data[idx+count] == digit {
				count++
			}
			result.WriteString(strconv.Itoa(count))
			result.WriteByte(digit)
			idx += count
		}
		data = result.String()
	}
	return len(data)
}

func PartOne(data string) int {
	return applyLookAndSay(data, 40)
}

func PartTwo(data string) int {
	return applyLookAndSay(data, 50)
}
