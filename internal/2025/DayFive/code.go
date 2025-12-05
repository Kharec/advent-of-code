package dayfive

import (
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"advent-of-code/internal/registry"
)

type freshRange struct {
	start int
	end   int
}

func init() {
	registry.Register("2025D5", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(input []string) int {
	var freshRanges []freshRange
	var ingredientIDs []int
	separatorFound := false

	for _, line := range input {
		if line == "" {
			separatorFound = true
			continue
		}

		if !separatorFound {
			startStr, endStr, _ := strings.Cut(line, "-")
			start, _ := strconv.Atoi(startStr)
			end, _ := strconv.Atoi(endStr)
			freshRanges = append(freshRanges, freshRange{start: start, end: end})
		} else {
			id, _ := strconv.Atoi(line)
			ingredientIDs = append(ingredientIDs, id)
		}
	}

	slices.SortFunc(freshRanges, func(a, b freshRange) int {
		switch {
		case a.start < b.start:
			return -1
		case a.start > b.start:
			return 1
		default:
			return 0
		}
	})

	freshCount := 0
	for _, id := range ingredientIDs {
		idx := sort.Search(len(freshRanges), func(idx int) bool {
			return freshRanges[idx].start > id
		})

		for i := idx - 1; i >= 0 && freshRanges[i].start <= id; i-- {
			if id <= freshRanges[i].end {
				freshCount++
				break
			}
		}
	}

	return freshCount
}

func PartTwo(input []string) int {
	return 0
}
