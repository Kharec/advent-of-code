# Advent of Code

Personal solutions for [Advent of Code](https://adventofcode.com/), all implemented in Go.

The goal is to practice programming and problem-solving habits while keeping each day's solution self-contained and easy to revisit later.

It uses pure Go, no external dependencies.

Also, it's a fresh start from 2025. I do some exercises from other years along the way, which explains why no year is complete and why everything is in "recent" Go.

## Requirements

- Go 1.25
- Puzzle input

## Repository Structure

```
├── cmd/
│   └── aoc/
│       └── main.go                # CLI entry point that runs solutions
└── internal/
    ├── 2020/
    │   ├── DayOne/
    │   │   ├── code.go            # Go solution for Day 1 (2020)
    │   │   └── code_test.go       # Unit tests for Day 1
    │   └── ...
    ├── 2021/
    │   └── ...                    # Additional years and days
    ├── registry/
    │   └── registry.go            # Central registry for day runners
    └── data/
        ├── 2020/
        │   ├── DayOne/
        │   │   └── input.txt      # Puzzle input for Day 1 (2020)
        │   └── ...
        └── ...
```

Each day's solution is organized into its own folder, named according to the day number (e.g., `DayOne`, `DayTwo`). Inside each folder, you'll find a `code.go` file with the Go implementation of the solution and a `code_test.go` file containing corresponding tests. The input data for each puzzle is located in the `internal/data` directory, mirroring the code structure for easy access.

To connect each solution to the CLI, the repository uses a central registry located in `internal/registry/registry.go`. This registry is essentially a map linking a day key (combining year and day number) to assign a "day runner" which is a struct that specifies the `ParseInput`, `PartOne` and `PartTwo` functions of the day's problem.

When running a solution, the CLI (`cmd/aoc/main.go`) looks up the appropriate runner from the registry based on your command-line input, uses the parsing function to load the input data, and then runs the desired part (or both parts) using the registered solution functions.

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
