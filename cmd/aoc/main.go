package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"advent-of-code/internal/registry"

	// 2015
	_ "advent-of-code/internal/2015/DayFive"
	_ "advent-of-code/internal/2015/DayFour"
	_ "advent-of-code/internal/2015/DayOne"
	_ "advent-of-code/internal/2015/DaySeven"
	_ "advent-of-code/internal/2015/DaySix"
	_ "advent-of-code/internal/2015/DayThree"
	_ "advent-of-code/internal/2015/DayTwo"

	// 2018
	_ "advent-of-code/internal/2018/DayOne"
	_ "advent-of-code/internal/2018/DayThree"
	_ "advent-of-code/internal/2018/DayTwo"

	// 2020
	_ "advent-of-code/internal/2020/DayEight"
	_ "advent-of-code/internal/2020/DayFour"
	_ "advent-of-code/internal/2020/DayOne"
	_ "advent-of-code/internal/2020/DaySeven"
	_ "advent-of-code/internal/2020/DaySix"
	_ "advent-of-code/internal/2020/DayThree"
	_ "advent-of-code/internal/2020/DayTwo"

	// 2021
	_ "advent-of-code/internal/2021/DayOne"
	_ "advent-of-code/internal/2021/DayThree"
	_ "advent-of-code/internal/2021/DayTwo"

	// 2022
	_ "advent-of-code/internal/2022/DayOne"
	_ "advent-of-code/internal/2022/DayTwo"

	// 2025
	_ "advent-of-code/internal/2025/DayOne"
	_ "advent-of-code/internal/2025/DayTwo"
)

func capitalize(day string) string {
	dayNames := map[string]string{
		"1": "One", "2": "Two", "3": "Three", "4": "Four", "5": "Five",
		"6": "Six", "7": "Seven", "8": "Eight", "9": "Nine", "10": "Ten",
		"11": "Eleven", "12": "Twelve", "13": "Thirteen", "14": "Fourteen", "15": "Fifteen",
		"16": "Sixteen", "17": "Seventeen", "18": "Eighteen", "19": "Nineteen", "20": "Twenty",
		"21": "TwentyOne", "22": "TwentyTwo", "23": "TwentyThree", "24": "TwentyFour", "25": "TwentyFive",
	}
	return dayNames[day]
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <YEAR>D<DAY>[P<PART>]\n", os.Args[0])
		os.Exit(1)
	}

	re := regexp.MustCompile(`^(\d{4})D(\d+)(?:P(\d+))?$`)
	matches := re.FindStringSubmatch(os.Args[1])
	if matches == nil {
		fmt.Fprintf(os.Stderr, "Invalid format: %s\n", os.Args[1])
		os.Exit(1)
	}

	year, dayNum, part := matches[1], matches[2], matches[3]
	dayKey := fmt.Sprintf("%sD%s", year, dayNum)
	dayName := fmt.Sprintf("Day%s", capitalize(dayNum))
	input := filepath.Join("internal", "data", year, dayName, "input.txt")

	runner, ok := registry.Days[dayKey].(registry.Runner)
	if !ok {
		log.Fatalf("Day not found: %s\n", dayKey)
	}

	runner.Run(input, part)
}
