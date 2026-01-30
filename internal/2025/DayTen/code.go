package dayten

import (
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"advent-of-code/internal/registry"
)

const (
	positiveInfinity = math.MaxFloat64
	negativeInfinity = -math.MaxFloat64
	epsilon          = 1e-9
)

var (
	bracketRegex     = regexp.MustCompile(`\[([.#]+)\]`)
	parenthesisRegex = regexp.MustCompile(`\(([^)]+)\)`)
	braceRegex       = regexp.MustCompile(`\{([^}]+)\}`)
)

func init() {
	registry.Register("2025D10", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func parseMachine(line string) ([]bool, [][]int) {
	bracketMatch := bracketRegex.FindStringSubmatch(line)
	targetString := bracketMatch[1]
	target := make([]bool, len(targetString))
	for idx, char := range targetString {
		target[idx] = char == '#'
	}

	parenthesisMatches := parenthesisRegex.FindAllStringSubmatch(line, -1)
	buttons := make([][]int, 0, len(parenthesisMatches))

	for _, match := range parenthesisMatches {
		buttonString := match[1]
		parts := strings.Split(buttonString, ",")
		button := make([]int, 0, len(parts))
		for _, part := range parts {
			light, _ := strconv.Atoi(part)
			button = append(button, light)
		}
		buttons = append(buttons, button)
	}

	return target, buttons
}

func solveIntegerLinearMinimization(constraints [][]float64, objective []float64) int {
	best := positiveInfinity
	stack := [][][]float64{constraints}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		numberOfRows := len(current)
		numberOfVariables := len(current[0]) - 1
		rightHandSideIdx := numberOfVariables + 1

		nonBasic := make([]int, numberOfVariables+1)
		for idx := range numberOfVariables {
			nonBasic[idx] = idx
		}
		nonBasic[numberOfVariables] = -1

		basic := make([]int, numberOfRows)
		for idx := range numberOfRows {
			basic[idx] = numberOfVariables + idx
		}

		tableau := make([][]float64, numberOfRows+2)
		for idx := range numberOfRows {
			tableau[idx] = make([]float64, numberOfVariables+2)
			copy(tableau[idx][:numberOfVariables+1], current[idx])
			tableau[idx][rightHandSideIdx] = -1.0
			tableau[idx][numberOfVariables], tableau[idx][rightHandSideIdx] = tableau[idx][rightHandSideIdx], tableau[idx][numberOfVariables]
		}
		tableau[numberOfRows] = make([]float64, numberOfVariables+2)
		copy(tableau[numberOfRows][:numberOfVariables], objective)
		tableau[numberOfRows+1] = make([]float64, numberOfVariables+2)
		tableau[numberOfRows+1][numberOfVariables] = 1.0

		pivot := func(row, column int) {
			k := 1.0 / tableau[row][column]
			for i := 0; i < numberOfRows+2; i++ {
				if i == row {
					continue
				}
				for j := 0; j < numberOfVariables+2; j++ {
					if j != column {
						tableau[i][j] -= tableau[row][j] * tableau[i][column] * k
					}
				}
			}
			for j := 0; j < numberOfVariables+2; j++ {
				tableau[row][j] *= k
			}
			for i := 0; i < numberOfRows+2; i++ {
				tableau[i][column] *= -k
			}
			tableau[row][column] = k
			basic[row], nonBasic[column] = nonBasic[column], basic[row]
		}

		findOptimalPivot := func(phase int) bool {
			for {
				pivotColumn := -1
				minimumValue := positiveInfinity
				minimumIndex := math.MaxInt

				for idx := 0; idx <= numberOfVariables; idx++ {
					if phase == 0 && nonBasic[idx] == -1 {
						continue
					}
					value := tableau[numberOfRows+phase][idx]
					if pivotColumn == -1 || value < minimumValue-epsilon || (math.Abs(value-minimumValue) <= epsilon && nonBasic[idx] < minimumIndex) {
						pivotColumn = idx
						minimumValue = value
						minimumIndex = nonBasic[idx]
					}
				}

				if tableau[numberOfRows+phase][pivotColumn] > -epsilon {
					return true
				}

				pivotRow := -1
				minimumRatio := positiveInfinity
				minimumBasicIndex := math.MaxInt

				for idx := range numberOfRows {
					if tableau[idx][pivotColumn] > epsilon {
						ratio := tableau[idx][rightHandSideIdx] / tableau[idx][pivotColumn]
						if pivotRow == -1 || ratio < minimumRatio-epsilon || (math.Abs(ratio-minimumRatio) <= epsilon && basic[idx] < minimumBasicIndex) {
							pivotRow = idx
							minimumRatio = ratio
							minimumBasicIndex = basic[idx]
						}
					}
				}

				if pivotRow == -1 {
					return false
				}

				pivot(pivotRow, pivotColumn)
			}
		}

		extractSolution := func() (float64, []float64) {
			solution := make([]float64, numberOfVariables)
			for idx := range numberOfRows {
				if basic[idx] >= 0 && basic[idx] < numberOfVariables {
					solution[basic[idx]] = tableau[idx][rightHandSideIdx]
				}
			}
			result := 0.0
			for idx := range numberOfVariables {
				result += objective[idx] * solution[idx]
			}
			return result, solution
		}

		var value float64
		var solution []float64

		mostNegativeRow := 0
		mostNegativeValue := tableau[0][rightHandSideIdx]
		for idx := 1; idx < numberOfRows; idx++ {
			if tableau[idx][rightHandSideIdx] < mostNegativeValue {
				mostNegativeValue = tableau[idx][rightHandSideIdx]
				mostNegativeRow = idx
			}
		}

		if mostNegativeValue < -epsilon {
			pivot(mostNegativeRow, numberOfVariables)
			if !findOptimalPivot(1) || tableau[numberOfRows+1][rightHandSideIdx] < -epsilon {
				value = negativeInfinity
				solution = nil
			} else {
				for i := range numberOfRows {
					if basic[i] == -1 {
						column := 0
						columnValue := tableau[i][0]
						columnIndex := nonBasic[0]
						for j := 1; j < numberOfVariables; j++ {
							if tableau[i][j] < columnValue-epsilon || (math.Abs(tableau[i][j]-columnValue) <= epsilon && nonBasic[j] < columnIndex) {
								column = j
								columnValue = tableau[i][j]
								columnIndex = nonBasic[j]
							}
						}
						pivot(i, column)
					}
				}
				if findOptimalPivot(0) {
					value, solution = extractSolution()
				} else {
					value = negativeInfinity
					solution = nil
				}
			}
		} else {
			if findOptimalPivot(0) {
				value, solution = extractSolution()
			} else {
				value = negativeInfinity
				solution = nil
			}
		}

		if value == negativeInfinity || value >= best-epsilon {
			continue
		}

		fractionalIdx := -1
		var fractionalValue float64
		for idx, value := range solution {
			if math.Abs(value-math.Round(value)) > epsilon {
				fractionalIdx = idx
				fractionalValue = value
				break
			}
		}

		if fractionalIdx == -1 {
			best = value
			continue
		}

		numberOfColumns := len(current[0])
		fractionalPart := fractionalValue - math.Floor(fractionalValue)

		ceilConstraint := make([]float64, numberOfColumns)
		ceilConstraint[fractionalIdx] = -1.0
		ceilConstraint[numberOfColumns-1] = -math.Ceil(fractionalValue)
		ceilBranch := make([][]float64, len(current)+1)
		copy(ceilBranch, current)
		ceilBranch[len(current)] = ceilConstraint

		floorConstraint := make([]float64, numberOfColumns)
		floorConstraint[fractionalIdx] = 1.0
		floorConstraint[numberOfColumns-1] = math.Floor(fractionalValue)
		floorBranch := make([][]float64, len(current)+1)
		copy(floorBranch, current)
		floorBranch[len(current)] = floorConstraint

		if fractionalPart > 0.5 {
			stack = append(stack, ceilBranch, floorBranch)
		} else {
			stack = append(stack, floorBranch, ceilBranch)
		}
	}

	return int(math.Round(best))
}

func findMinimumJoltagePresses(buttons [][]int, joltages []int) int {
	numberOfCounters, numberOfButtons := len(joltages), len(buttons)
	numberOfRows, numberOfColumns := 2*numberOfCounters+numberOfButtons, numberOfButtons+1
	constraints := make([][]float64, numberOfRows)
	for i := range constraints {
		constraints[i] = make([]float64, numberOfColumns)
	}
	for j := range numberOfButtons {
		constraints[numberOfRows-1-j][j] = -1.0
	}
	for buttonIdx, button := range buttons {
		for _, counterIdx := range button {
			constraints[counterIdx][buttonIdx] = 1.0
			constraints[counterIdx+numberOfCounters][buttonIdx] = -1.0
		}
	}
	for idx, value := range joltages {
		value := float64(value)
		constraints[idx][numberOfColumns-1] = value
		constraints[idx+numberOfCounters][numberOfColumns-1] = -value
	}
	objective := make([]float64, numberOfButtons)
	for idx := range objective {
		objective[idx] = 1.0
	}

	return solveIntegerLinearMinimization(constraints, objective)
}

func PartOne(data []string) int {
	total := 0
	for _, line := range data {
		target, buttons := parseMachine(line)
		numberOfLights := len(target)
		numberOfButtons := len(buttons)
		minimumPresses := numberOfButtons + 1

		for mask := 0; mask < (1 << numberOfButtons); mask++ {
			lights := make([]bool, numberOfLights)
			presses := 0

			for buttonIdx := range numberOfButtons {
				if mask&(1<<buttonIdx) != 0 {
					presses++
					for _, light := range buttons[buttonIdx] {
						if light < numberOfLights {
							lights[light] = !lights[light]
						}
					}
				}
			}

			matches := true
			for idx := range numberOfLights {
				if lights[idx] != target[idx] {
					matches = false
					break
				}
			}

			if matches && presses < minimumPresses {
				minimumPresses = presses
			}
		}
		total += minimumPresses
	}

	return total
}

func PartTwo(data []string) int {
	total := 0
	for _, line := range data {
		_, buttons := parseMachine(line)
		braceMatch := braceRegex.FindStringSubmatch(line)
		parts := strings.Split(braceMatch[1], ",")
		joltages := make([]int, len(parts))
		for idx, part := range parts {
			joltages[idx], _ = strconv.Atoi(part)
		}
		total += findMinimumJoltagePresses(buttons, joltages)
	}
	return total
}
