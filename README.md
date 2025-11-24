# Advent of Code

Personal solutions for [Advent of Code](https://adventofcode.com/), all implemented in Go.

The goal is to practice programming and problem-solving habits while keeping each day's solution self-contained and easy to revisit later.

It uses pure Go, no external dependencies.

Also, tt's a fresh start from 2025. So I was able to do some exercises from other years along the way, which explains why no year is complete and why everything is in "recent" Go.

## Requirements

- Go 1.25
- Puzzle input

## Repository Structure

```
yyyy/
├── dayXX/
│   ├── main.go      # The main code to run to print solutions.
│   ├── main_test.go # The test file that validates the logic.
│   └── input.txt    # The day's puzzle input.
└── ...
```

Each day's code can be run with:

```bash
cd yyyy/dayXX
go run main.go
```

Expected output:

```bash
Part 1: <answer>
Part 2: <answer>
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

I'm using these examples to validate my logic.

If you want to run those, you can use:

```bash
cd dayXX
go test
```

## Notes

Happy coding, and good luck on the puzzles!
