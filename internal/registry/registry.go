package registry

import (
	"fmt"
	"os"
)

type Runner interface {
	Run(input string, part string)
}

type DayRunner[T any] struct {
	ParseInput func(string) T
	PartOne    func(T) int
	PartTwo    func(T) int
}

func (r *DayRunner[T]) Run(input string, part string) {
	data := r.ParseInput(input)
	switch part {
	case "":
		fmt.Println(r.PartOne(data))
		fmt.Println(r.PartTwo(data))
	case "1":
		fmt.Println(r.PartOne(data))
	case "2":
		fmt.Println(r.PartTwo(data))
	default:
		fmt.Fprintf(os.Stderr, "Invalid part: %s\n", part)
		os.Exit(1)
	}
}

var Days = make(map[string]any)

func Register[T any](key string, parseInput func(string) T, partOne func(T) int, partTwo func(T) int) {
	Days[key] = &DayRunner[T]{
		ParseInput: parseInput,
		PartOne:    partOne,
		PartTwo:    partTwo,
	}
}
