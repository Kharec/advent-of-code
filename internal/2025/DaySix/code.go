package daysix

import (
	"os"
	"strconv"
	"strings"

	"advent-of-code/internal/registry"
)

func init() {
	registry.Register("2025D6", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	lines := strings.Split(strings.TrimRight(string(content), "\n"), "\n")
	return lines
}

func transpose(lines []string) []string {
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	transposed := make([]string, maxWidth)
	for column := 0; column < maxWidth; column++ {
		var builder strings.Builder
		for _, line := range lines {
			if column < len(line) {
				builder.WriteByte(line[column])
			} else {
				builder.WriteByte(' ')
			}
		}
		transposed[column] = builder.String()
	}

	return transposed
}

func isSpaceColumn(column string) bool {
	return strings.TrimSpace(column) == ""
}

func extractRowFromColumns(transposed []string, row, startColumn, endColumn int) string {
	var builder strings.Builder
	for column := startColumn; column <= endColumn && column < len(transposed); column++ {
		if row < len(transposed[column]) {
			builder.WriteByte(transposed[column][row])
		} else {
			builder.WriteByte(' ')
		}
	}
	return builder.String()
}

func extractNumbers(transposed []string, startColumn, endColumn, operationRow int) []int {
	numbers := make([]int, 0)
	for row := range operationRow {
		rowStr := extractRowFromColumns(transposed, row, startColumn, endColumn)
		trimmed := strings.TrimSpace(rowStr)
		fields := strings.FieldsSeq(trimmed)
		for field := range fields {
			if number, err := strconv.Atoi(field); err == nil {
				numbers = append(numbers, number)
			}
		}
	}
	return numbers
}

func extractOperation(segment string) byte {
	if idx := strings.IndexAny(segment, "+*"); idx != -1 {
		return segment[idx]
	}
	return 0
}

func applyOperation(numbers []int, operation byte) int {
	result := numbers[0]
	if operation == '+' {
		for idx := 1; idx < len(numbers); idx++ {
			result += numbers[idx]
		}
	} else {
		for idx := 1; idx < len(numbers); idx++ {
			result *= numbers[idx]
		}
	}
	return result
}

func PartOne(input []string) int {
	operationRow := len(input) - 1
	transposed := transpose(input)

	total := 0
	column := 0
	for column < len(transposed) {
		if isSpaceColumn(transposed[column]) {
			column++
			continue
		}

		startColumn := column
		for column < len(transposed) && !isSpaceColumn(transposed[column]) {
			column++
		}
		endColumn := column - 1

		numbers := extractNumbers(transposed, startColumn, endColumn, operationRow)
		operationRowStr := extractRowFromColumns(transposed, operationRow, startColumn, endColumn)
		operation := extractOperation(operationRowStr)

		if len(numbers) > 0 && operation != 0 {
			result := applyOperation(numbers, operation)
			total += result
		}
	}

	return total
}

func PartTwo(input []string) int {
	return 0
}
