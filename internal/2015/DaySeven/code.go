package dayseven

import (
	"maps"
	"os"
	"strconv"
	"strings"

	"advent-of-code/internal/registry"
)

func init() {
	registry.Register("2015D7", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) map[string]string {
	content, _ := os.ReadFile(filepath)
	instructions := make(map[string]string)
	for line := range strings.SplitSeq(string(content), "\n") {
		if parts := strings.Split(line, " -> "); len(parts) == 2 {
			instructions[parts[1]] = parts[0]
		}
	}
	return instructions
}

func evaluateWire(wire string, instructions map[string]string) uint16 {
	cache := make(map[string]uint16)
	var evaluate func(string) uint16
	evaluate = func(wire string) uint16 {
		if value, ok := cache[wire]; ok {
			return value
		}
		if num, err := strconv.Atoi(wire); err == nil {
			return uint16(num)
		}
		expression, exists := instructions[wire]
		if !exists {
			return 0
		}
		parts := strings.Fields(expression)
		var result uint16
		switch {
		case len(parts) == 1:
			result = evaluate(parts[0])
		case parts[0] == "NOT":
			result = ^evaluate(parts[1])
		default:
			left, right := evaluate(parts[0]), evaluate(parts[2])
			switch parts[1] {
			case "AND":
				result = left & right
			case "OR":
				result = left | right
			case "LSHIFT":
				result = left << right
			case "RSHIFT":
				result = left >> right
			}
		}
		cache[wire] = result
		return result
	}
	return evaluate(wire)
}

func PartOne(instructions map[string]string) int {
	return int(evaluateWire("a", instructions))
}

func PartTwo(instructions map[string]string) int {
	signalA := PartOne(instructions)
	instructionsCopy := make(map[string]string, len(instructions))
	maps.Copy(instructionsCopy, instructions)
	instructionsCopy["b"] = strconv.Itoa(signalA)
	return int(evaluateWire("a", instructionsCopy))
}
