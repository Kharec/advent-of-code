package daythree

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
)

func init() {
	registry.Register("2017D3", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) int {
	content, _ := os.ReadFile(filepath)
	data, _ := strconv.Atoi(string(content))
	return data
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func PartOne(data int) int {
	layer := 1

	for (2*layer+1)*(2*layer+1) < data {
		layer++
	}

	sideLength := 2*layer + 1
	offset := sideLength*sideLength - data
	side := offset / (sideLength - 1)
	position := offset % (sideLength - 1)

	coordinates := [4][2]int{
		{layer, layer - position},
		{layer - position, -layer},
		{-layer, -layer + position},
		{-layer + position, layer},
	}

	return abs(coordinates[side][0]) + abs(coordinates[side][1])
}

func PartTwo(data int) int {
	grid := make(map[[2]int]int)
	grid[[2]int{0, 0}] = 1

	directions := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	x, y := 1, 0
	directionIdx := 0
	sideLength := 1
	stepsInSide := 1

	for {
		sum := grid[[2]int{x - 1, y - 1}] + grid[[2]int{x - 1, y}] + grid[[2]int{x - 1, y + 1}] +
			grid[[2]int{x, y - 1}] + grid[[2]int{x, y + 1}] +
			grid[[2]int{x + 1, y - 1}] + grid[[2]int{x + 1, y}] + grid[[2]int{x + 1, y + 1}]

		if sum > data {
			return sum
		}

		grid[[2]int{x, y}] = sum

		if stepsInSide == sideLength {
			stepsInSide = 0
			if directionIdx%2 == 1 {
				sideLength++
			}
			directionIdx = (directionIdx + 1) % 4
		}

		x += directions[directionIdx][0]
		y += directions[directionIdx][1]
		stepsInSide++
	}
}
