# Advent of Code

Personal solutions for [Advent of Code](https://adventofcode.com/), all implemented in Go.

The goal is to practice programming and problem-solving habits while keeping each day's solution self-contained and easy to revisit later.

It uses pure Go, no external dependencies.

Also, it's a fresh start from 2025. I do some exercises from other years along the way, which explains why no year is complete and why everything is in "recent" Go.

Ultimately, my goal is to complete all the years of Advent of Code here.

## Requirements

- Go 1.25
- Puzzle input

## Repository Structure

```
├── cmd/
│   └── aoc/
│       └── main.go               # CLI entry point for running puzzles
└── internal/
    ├── 2020/
    │   ├── register.go           # Aggregates import for main.go
    │   ├── DayOne/
    │   │   ├── code.go           # Day 1 solution for 2020
    │   │   └── code_test.go      # Unit tests for Day 1 solution
    │   └── ...                   # Other days for 2020
    ├── 2021/
    │   └── ...                   # 2021 days code and tests
    ├── registry/
    │   └── registry.go           # Central map/registry for finding and running days
    └── data/
        ├── 2020/
        │   ├── DayOne/
        │   │   └── input.txt     # Puzzle input for Day 1 (2020)
        │   └── ...               # Inputs for other days
        └── ...                   # Inputs for other years
```

Each day's solution is organized into its own folder, named according to the day number (e.g., `DayOne`, `DayTwo`). Inside each folder, you'll find a `code.go` file with the Go implementation of the solution and a `code_test.go` file containing corresponding tests. The input data for each puzzle is located in the `internal/data` directory, mirroring the code structure for easy access.

Each year has its own `register.go` file to aggregate imports. To connect each solution to the CLI, the repository uses a central registry located in `internal/registry/registry.go`. This registry is essentially a map linking a day key (combining year and day number) to assign a "day runner" which is a struct that specifies the `ParseInput`, `PartOne` and `PartTwo` functions of the day's problem.

When running a solution, the CLI (`cmd/aoc/main.go`) looks up the appropriate runner from the registry based on your command-line input, uses the parsing function to load the input data, and then runs the desired part (or both parts) using the registered solution functions.

### Starting a new day

```bash
make 2020D1
```

If the `ADVENTOFCODE_SESSION` environment variable is set with your Advent of Code session cookie, the input will be downloaded from the AOC website right into the corresponding `internal/data` directory, else it will create an empty `input.txt` file.

Then, it will create the following files:

```bash
internal/2020/DayOne/code.go
internal/2020/DayOne/code_test.go
```

### Handling string answers

Occasionally, a puzzle requires a string as the answer. However, my solution framework expects PartOne/PartTwo functions to return integers.

To work around this, I print the string result (see [2018D2P2](https://git.kharec.info/Kharec/advent-of-code/tree/main/internal/2018/DayTwo/code.go)) and return a dummy integer to satisfy the type requirements.

Meanwhile, the corresponding [test](https://git.kharec.info/Kharec/advent-of-code/tree/main/internal/2018/DayTwo/code_test.go) captures the output and verifies that the correct string is produced.

It's not the most elegant solution, but since only a handful of days across all Advent of Code years require a string result, this compromise keeps the rest of the code simple.

## Usage

Build the CLI tool:

```bash
make
```

Run a day's solution:

```bash
bin/aoc 2020D1      # Run both parts
bin/aoc 2020D1P1    # Run only part one
bin/aoc 2020D1P2    # Run only part two
```

Example output:

```bash
$ bin/aoc 2020D1
197451
138233720

$ bin/aoc 2020D1P1
197451

$ bin/aoc 2020D1P2
138233720
```

## Tests

In the Advent of Code, every day the logic is explained with a sample result, like this:

```
For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines are 12, 38, 15, and 77.

Adding these together produces 142.
```

I'm using these examples each day to validate my logic.

Run all tests:

```bash
make test
```

Or run tests for a specific day:

```bash
go test ./internal/2020/DayOne/...
```

## Notes

Happy coding, and good luck on the puzzles!
