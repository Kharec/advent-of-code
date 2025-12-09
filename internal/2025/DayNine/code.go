package daynine

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2025D9", ParseInput, PartOne, PartTwo)
}

type Point struct {
	x, y int
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func parsePoints(data []string) []Point {
	points := make([]Point, 0, len(data))
	for _, line := range data {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, Point{x: x, y: y})
	}
	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func PartOne(data []string) int {
	points := parsePoints(data)
	maxArea := 0

	for i := range points {
		for j := 1; j < len(points); j++ {
			distanceX := abs(points[i].x - points[j].x)
			distanceY := abs(points[i].y - points[j].y)

			if distanceX > 0 && distanceY > 0 {
				area := (distanceX + 1) * (distanceY + 1)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func PartTwo(data []string) int {
	return 0
}
