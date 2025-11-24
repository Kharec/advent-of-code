package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(file string) []int {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	lines := strings.Fields(string(content))
	var data []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v", err)
		}
		data = append(data, num)
	}
	return data
}

func PartOne(data []int) int {
	seen := make(map[int]bool)
	for _, number := range data {
		complement := 2020 - number
		if seen[complement] {
			return number * complement
		}
		seen[number] = true
	}
	return 0
}

func PartTwo(data []int) int {
	seen := make(map[int]bool)
	for idx := range data {
		for _, number := range data[idx+1:] {
			complement := 2020 - data[idx] - number
			if seen[complement] {
				return complement * data[idx] * number
			}
		}
		seen[data[idx]] = true
	}
	return 0
}

func main() {
	data := parseInput("input.txt")
	fmt.Println("Part 1:", PartOne(data))
	fmt.Println("Part 2:", PartTwo(data))
}
