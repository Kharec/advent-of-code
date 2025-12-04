package daytwelve

import (
	"encoding/json"
	"os"
	"regexp"
	"strconv"
	"strings"

	"advent-of-code/internal/registry"
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
	var sumNumbers func(any) int
	sumNumbers = func(v any) int {
		switch value := v.(type) {
		case float64:
			return int(value)
		case []any:
			sum := 0
			for _, item := range value {
				sum += sumNumbers(item)
			}
			return sum
		case map[string]any:
			sum := 0
			for _, item := range value {
				if str, ok := item.(string); ok && str == "red" {
					return 0
				}
				sum += sumNumbers(item)
			}
			return sum
		default:
			return 0
		}
	}

	sum := 0
	for _, line := range data {
		var value any
		_ = json.Unmarshal([]byte(line), &value)
		sum += sumNumbers(value)
	}
	return sum
}
