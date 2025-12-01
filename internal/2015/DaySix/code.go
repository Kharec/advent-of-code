package daysix

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2015D6", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

func PartOne(data []string) int {
	grid := make([]bool, 1000*1000)

	for _, line := range data {
		var op string
		var x1, y1, x2, y2 int

		switch {
		case strings.HasPrefix(line, "turn on "):
			op = "on"
			line = strings.TrimPrefix(line, "turn on ")
		case strings.HasPrefix(line, "turn off "):
			op = "off"
			line = strings.TrimPrefix(line, "turn off ")
		case strings.HasPrefix(line, "toggle "):
			op = "toggle"
			line = strings.TrimPrefix(line, "toggle ")
		default:
			continue
		}

		parts := strings.Split(line, " through ")
		firstCoordinates := strings.Split(parts[0], ",")
		secondCoordinates := strings.Split(parts[1], ",")

		x1, _ = strconv.Atoi(firstCoordinates[0])
		y1, _ = strconv.Atoi(firstCoordinates[1])
		x2, _ = strconv.Atoi(secondCoordinates[0])
		y2, _ = strconv.Atoi(secondCoordinates[1])

		for y := y1; y <= y2; y++ {
			for x := x1; x <= x2; x++ {
				idx := y*1000 + x
				switch op {
				case "on":
					grid[idx] = true
				case "off":
					grid[idx] = false
				case "toggle":
					grid[idx] = !grid[idx]
				}
			}
		}
	}

	count := 0
	for _, lit := range grid {
		if lit {
			count++
		}
	}
	return count
}

func PartTwo(data []string) int {
	grid := make([]int, 1000*1000)

	for _, line := range data {
		var op string
		var x1, y1, x2, y2 int

		switch {
		case strings.HasPrefix(line, "turn on "):
			op = "on"
			line = strings.TrimPrefix(line, "turn on ")
		case strings.HasPrefix(line, "turn off "):
			op = "off"
			line = strings.TrimPrefix(line, "turn off ")
		case strings.HasPrefix(line, "toggle "):
			op = "toggle"
			line = strings.TrimPrefix(line, "toggle ")
		default:
			continue
		}

		parts := strings.Split(line, " through ")
		firstCoordinates := strings.Split(parts[0], ",")
		secondCoordinates := strings.Split(parts[1], ",")

		x1, _ = strconv.Atoi(firstCoordinates[0])
		y1, _ = strconv.Atoi(firstCoordinates[1])
		x2, _ = strconv.Atoi(secondCoordinates[0])
		y2, _ = strconv.Atoi(secondCoordinates[1])

		for y := y1; y <= y2; y++ {
			for x := x1; x <= x2; x++ {
				idx := y*1000 + x
				switch op {
				case "on":
					grid[idx]++
				case "off":
					if grid[idx] > 0 {
						grid[idx]--
					}
				case "toggle":
					grid[idx] += 2
				}
			}
		}
	}

	total := 0
	for _, brightness := range grid {
		total += brightness
	}
	return total
}
