package daynine

import (
	"advent-of-code/internal/registry"
	"os"
	"slices"
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

func minMax(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func compressCoordinates(points []Point) ([]int, []int, map[int]int, map[int]int) {
	xs := make(map[int]bool)
	ys := make(map[int]bool)
	for _, point := range points {
		xs[point.x] = true
		ys[point.y] = true
	}
	sortedX := make([]int, 0, len(xs))
	sortedY := make([]int, 0, len(ys))
	for x := range xs {
		sortedX = append(sortedX, x)
	}
	for y := range ys {
		sortedY = append(sortedY, y)
	}
	slices.Sort(sortedX)
	slices.Sort(sortedY)

	mapX := make(map[int]int, len(sortedX))
	mapY := make(map[int]int, len(sortedY))
	for idx, x := range sortedX {
		mapX[x] = 2 * idx
	}
	for idx, y := range sortedY {
		mapY[y] = 2 * idx
	}
	return sortedX, sortedY, mapX, mapY
}

func drawBoundary(grid [][]bool, points []Point, mapX, mapY map[int]int) {
	for idx := range points {
		current, next := points[idx], points[(idx+1)%len(points)]
		cx1, cy1 := mapX[current.x], mapY[current.y]
		cx2, cy2 := mapX[next.x], mapY[next.y]

		if cx1 == cx2 {
			minY, maxY := minMax(cy1, cy2)
			for y := minY; y <= maxY; y++ {
				grid[cx1][y] = true
			}
		} else {
			minX, maxX := minMax(cx1, cx2)
			for x := minX; x <= maxX; x++ {
				grid[x][cy1] = true
			}
		}
	}
}

func scanLineFill(grid [][]bool, points []Point, mapX, mapY map[int]int, _, height int) {
	type edge struct {
		x, yMin, yMax int
	}
	edges := make([]edge, 0, len(points))
	for idx := range points {
		current, next := points[idx], points[(idx+1)%len(points)]
		if current.x == next.x {
			minY, maxY := minMax(mapY[current.y], mapY[next.y])
			edges = append(edges, edge{mapX[current.x], minY, maxY})
		}
	}

	for y := range height {
		crossings := make([]int, 0, len(edges))
		for _, e := range edges {
			if e.yMin <= y && y < e.yMax {
				crossings = append(crossings, e.x)
			}
		}
		slices.Sort(crossings)
		for i := 0; i < len(crossings)-1; i += 2 {
			for x := crossings[i]; x <= crossings[i+1]; x++ {
				grid[x][y] = true
			}
		}
	}
}

func forceEmptyGaps(grid [][]bool, sortedX, sortedY []int, width, height int) {
	for idx := 0; idx < len(sortedX)-1; idx++ {
		if sortedX[idx+1] == sortedX[idx]+1 {
			cx := 2*idx + 1
			for y := range height {
				grid[cx][y] = true
			}
		}
	}
	for idx := 0; idx < len(sortedY)-1; idx++ {
		if sortedY[idx+1] == sortedY[idx]+1 {
			cy := 2*idx + 1
			for x := range width {
				grid[x][cy] = true
			}
		}
	}
}

func buildPrefixSum(grid [][]bool, width, height int) [][]int {
	sum := make([][]int, width+1)
	for idx := range sum {
		sum[idx] = make([]int, height+1)
	}
	for i := range width {
		for j := range height {
			value := 0
			if grid[i][j] {
				value = 1
			}
			sum[i+1][j+1] = value + sum[i][j+1] + sum[i+1][j] - sum[i][j]
		}
	}
	return sum
}

func getSum(sum [][]int, x1, y1, x2, y2 int) int {
	return sum[x2+1][y2+1] - sum[x1][y2+1] - sum[x2+1][y1] + sum[x1][y1]
}

func PartOne(data []string) int {
	points := parsePoints(data)
	maxArea := 0

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			dx := abs(points[i].x - points[j].x)
			dy := abs(points[i].y - points[j].y)

			if area := (dx + 1) * (dy + 1); area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func PartTwo(data []string) int {
	points := parsePoints(data)
	sortedX, sortedY, mapX, mapY := compressCoordinates(points)

	width, height := 2*len(sortedX), 2*len(sortedY)
	grid := make([][]bool, width)
	for idx := range grid {
		grid[idx] = make([]bool, height)
	}

	drawBoundary(grid, points, mapX, mapY)
	scanLineFill(grid, points, mapX, mapY, width, height)
	forceEmptyGaps(grid, sortedX, sortedY, width, height)

	sum := buildPrefixSum(grid, width, height)

	compressedPoints := make([]Point, len(points))
	for idx, point := range points {
		compressedPoints[idx] = Point{mapX[point.x], mapY[point.y]}
	}

	maxArea := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]
			dx := abs(p1.x - p2.x)
			dy := abs(p1.y - p2.y)

			if dx == 0 || dy == 0 {
				continue
			}

			area := (dx + 1) * (dy + 1)
			if area <= maxArea {
				continue
			}

			cx1, cy1 := compressedPoints[i].x, compressedPoints[i].y
			cx2, cy2 := compressedPoints[j].x, compressedPoints[j].y
			minCX, maxCX := minMax(cx1, cx2)
			minCY, maxCY := minMax(cy1, cy2)

			totalCells := (maxCX - minCX + 1) * (maxCY - minCY + 1)
			if getSum(sum, minCX, minCY, maxCX, maxCY) == totalCells {
				maxArea = area
			}
		}
	}

	return maxArea
}
