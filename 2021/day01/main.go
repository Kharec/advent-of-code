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
	var data []int
	for line := range strings.SplitSeq(string(content), "\n") {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v", err)
		}
		data = append(data, num)
	}
	return data
}

func PartOne(data []int) int {
	count := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			count++
		}
	}
	return count
}

func PartTwo(data []int) int {
	count := 0
	for i := 3; i < len(data); i++ {
		if data[i] > data[i-3] {
			count++
		}
	}
	return count
}

func main() {
	data := parseInput("input.txt")
	fmt.Println("Part 1:", PartOne(data))
	fmt.Println("Part 2:", PartTwo(data))
}
