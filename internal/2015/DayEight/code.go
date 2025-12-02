package dayeight

import (
	"os"
	"strings"

	"advent-of-code/internal/registry"
)

func init() {
	registry.Register("2015D8", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	result := 0
	for _, line := range data {
		codeLength := len(line)
		memoryLength := 0
		for idx := 1; idx < len(line)-1; idx++ {
			if line[idx] == '\\' && idx+1 < len(line)-1 {
				switch line[idx+1] {
				case '\\', '"':
					memoryLength++
					idx++
				case 'x':
					if idx+3 < len(line)-1 {
						memoryLength++
						idx += 3
					}
				}
			} else {
				memoryLength++
			}
		}
		result += codeLength - memoryLength
	}
	return result
}

func PartTwo(data []string) int {
	return 0
}
