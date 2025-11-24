package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parseInput(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

func PartOne(data []string) int {
	bitlen := len(data[0])
	ones := make([]int, bitlen)

	for _, line := range data {
		for i := range bitlen {
			if line[i] == '1' {
				ones[i]++
			}
		}
	}

	var gamma int
	total := len(data)
	for idx := range bitlen {
		gamma <<= 1
		if ones[idx] > total-ones[idx] {
			gamma |= 1
		}
	}

	mask := (1 << bitlen) - 1
	epsilon := mask ^ gamma
	return gamma * epsilon
}

func PartTwo(data []string) int {
	bitlen := len(data[0])

	oxygenIndices := make([]int, len(data))
	for idx := range oxygenIndices {
		oxygenIndices[idx] = idx
	}

	co2Indices := make([]int, len(data))
	for idx := range co2Indices {
		co2Indices[idx] = idx
	}

	for indice := 0; indice < bitlen && len(oxygenIndices) > 1; indice++ {
		ones := 0
		for _, idx := range oxygenIndices {
			if data[idx][indice] == '1' {
				ones++
			}
		}
		target := byte('1')
		if ones*2 < len(oxygenIndices) {
			target = '0'
		}
		writeIdx := 0
		for readIdx := 0; readIdx < len(oxygenIndices); readIdx++ {
			if data[oxygenIndices[readIdx]][indice] == target {
				oxygenIndices[writeIdx] = oxygenIndices[readIdx]
				writeIdx++
			}
		}
		oxygenIndices = oxygenIndices[:writeIdx]
	}

	for idx := 0; idx < bitlen && len(co2Indices) > 1; idx++ {
		ones := 0
		for _, indice := range co2Indices {
			if data[indice][idx] == '1' {
				ones++
			}
		}
		target := byte('0')
		if ones*2 < len(co2Indices) {
			target = '1'
		}
		writeIdx := 0
		for readIdx := 0; readIdx < len(co2Indices); readIdx++ {
			if data[co2Indices[readIdx]][idx] == target {
				co2Indices[writeIdx] = co2Indices[readIdx]
				writeIdx++
			}
		}
		co2Indices = co2Indices[:writeIdx]
	}

	oxygenRating := 0
	for _, char := range data[oxygenIndices[0]] {
		oxygenRating <<= 1
		if char == '1' {
			oxygenRating |= 1
		}
	}

	co2Rating := 0
	for _, char := range data[co2Indices[0]] {
		co2Rating <<= 1
		if char == '1' {
			co2Rating |= 1
		}
	}

	return oxygenRating * co2Rating
}

func main() {
	data := parseInput("input.txt")
	fmt.Println("Part 1:", PartOne(data))
	fmt.Println("Part 2:", PartTwo(data))
}
