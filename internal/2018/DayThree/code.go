package daythree

import (
	"advent-of-code/internal/registry"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2018D3", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	coverage := make(map[string]int)
	re := regexp.MustCompile(`#\d+ @ (\d+),(\d+): (\d+)x(\d+)`)

	for _, line := range data {
		matches := re.FindStringSubmatch(line)

		left, _ := strconv.Atoi(matches[1])
		top, _ := strconv.Atoi(matches[2])
		width, _ := strconv.Atoi(matches[3])
		height, _ := strconv.Atoi(matches[4])

		for x := left; x < left+width; x++ {
			for y := top; y < top+height; y++ {
				key := fmt.Sprintf("%d,%d", x, y)
				coverage[key]++
			}
		}
	}

	overlapCount := 0
	for _, count := range coverage {
		if count >= 2 {
			overlapCount++
		}
	}

	return overlapCount
}

func PartTwo(data []string) int {
	coverage := make(map[string]int)
	re := regexp.MustCompile(`#\d+ @ (\d+),(\d+): (\d+)x(\d+)`)

	for _, line := range data {
		matches := re.FindStringSubmatch(line)

		left, _ := strconv.Atoi(matches[1])
		top, _ := strconv.Atoi(matches[2])
		width, _ := strconv.Atoi(matches[3])
		height, _ := strconv.Atoi(matches[4])

		for x := left; x < left+width; x++ {
			for y := top; y < top+height; y++ {
				key := fmt.Sprintf("%d,%d", x, y)
				coverage[key]++
			}
		}
	}

	reWithID := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	for _, line := range data {
		matches := reWithID.FindStringSubmatch(line)
		if len(matches) != 6 {
			continue
		}

		id, _ := strconv.Atoi(matches[1])
		left, _ := strconv.Atoi(matches[2])
		top, _ := strconv.Atoi(matches[3])
		width, _ := strconv.Atoi(matches[4])
		height, _ := strconv.Atoi(matches[5])

		overlaps := false
		for x := left; x < left+width && !overlaps; x++ {
			for y := top; y < top+height; y++ {
				key := fmt.Sprintf("%d,%d", x, y)
				if coverage[key] > 1 {
					overlaps = true
					break
				}
			}
		}

		if !overlaps {
			return id
		}
	}

	return 0
}
