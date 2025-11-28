package dayeight

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	argument  int
}

func init() {
	registry.Register("2020D8", ParseInput, PartOne, PartTwo)
}

func parseInstructions(data []string) []instruction {
	instructions := make([]instruction, 0, len(data))
	for _, line := range data {
		line = strings.TrimSpace(line)
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		arg, _ := strconv.Atoi(parts[1])
		instructions = append(instructions, instruction{operation: parts[0], argument: arg})
	}
	return instructions
}

func execute(instructions []instruction) (accumulator int, terminatedNormally bool) {
	visited := make(map[int]bool)
	accumulator = 0
	pc := 0

	for pc < len(instructions) {
		if visited[pc] {
			return accumulator, false
		}
		visited[pc] = true

		instruction := instructions[pc]
		switch instruction.operation {
		case "acc":
			accumulator += instruction.argument
			pc++
		case "jmp":
			pc += instruction.argument
		case "nop":
			pc++
		}
	}

	return accumulator, true
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	instructions := parseInstructions(data)
	accumulator, _ := execute(instructions)
	return accumulator
}

func PartTwo(data []string) int {
	instructions := parseInstructions(data)

	for idx := range instructions {
		if instructions[idx].operation != "jmp" && instructions[idx].operation != "nop" {
			continue
		}

		originalOperation := instructions[idx].operation
		if originalOperation == "jmp" {
			instructions[idx].operation = "nop"
		} else {
			instructions[idx].operation = "jmp"
		}

		accumulator, terminatedNormally := execute(instructions)
		instructions[idx].operation = originalOperation

		if terminatedNormally {
			return accumulator
		}
	}

	return 0
}
