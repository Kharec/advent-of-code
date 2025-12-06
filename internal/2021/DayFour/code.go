package dayfour

import (
	"os"
	"strconv"
	"strings"

	"advent-of-code/internal/registry"
)

const boardSize = 5

func init() {
	registry.Register("2021D4", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

type board struct {
	numbers      [boardSize][boardSize]int
	positionMap  map[int][2]int
	rowCounts    [boardSize]int
	columnCounts [boardSize]int
	marked       map[int]bool
	won          bool
}

func newBoard() board {
	return board{
		positionMap: make(map[int][2]int, boardSize*boardSize),
		marked:      make(map[int]bool),
	}
}

func parseBoards(lines []string) []board {
	var boards []board
	current := newBoard()
	row := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			if row == boardSize {
				boards = append(boards, current)
				current = newBoard()
				row = 0
			}
			continue
		}

		fields := strings.Fields(line)
		for column, field := range fields {
			number, _ := strconv.Atoi(field)
			current.numbers[row][column] = number
			current.positionMap[number] = [2]int{row, column}
		}
		row++
	}

	if row == boardSize {
		boards = append(boards, current)
	}

	return boards
}

func (b *board) mark(number int) bool {
	if b.won {
		return false
	}

	position, exists := b.positionMap[number]
	if !exists {
		return false
	}

	b.marked[number] = true
	row, column := position[0], position[1]
	b.rowCounts[row]++
	b.columnCounts[column]++

	if b.rowCounts[row] == boardSize || b.columnCounts[column] == boardSize {
		b.won = true
		return true
	}

	return false
}

func (b *board) sumUnmarked() int {
	sum := 0
	for row := range boardSize {
		for column := range boardSize {
			number := b.numbers[row][column]
			if !b.marked[number] {
				sum += number
			}
		}
	}
	return sum
}

func parseNumbers(line string) []int {
	numbersStr := strings.Split(line, ",")
	numbers := make([]int, 0, len(numbersStr))
	for _, numStr := range numbersStr {
		num, _ := strconv.Atoi(numStr)
		numbers = append(numbers, num)
	}
	return numbers
}

func PartOne(lines []string) int {
	numbers := parseNumbers(lines[0])
	boards := parseBoards(lines[1:])

	for _, number := range numbers {
		for idx := range boards {
			if boards[idx].mark(number) {
				return boards[idx].sumUnmarked() * number
			}
		}
	}

	return 0
}

func PartTwo(lines []string) int {
	return 0
}
