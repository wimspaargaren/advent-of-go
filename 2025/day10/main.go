package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/wimspaargaren/aoc"
)

func main() {
	input := aoc.MustReadFile("input.txt")

	machines := []Machine{}
	for _, line := range strings.Split(input, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		machines = append(machines, parseMachine(line))
	}

	totalPresses := 0.0

	for _, m := range machines {
		totalPresses += solve(floatSlice(m.JoltageRequirements), m.Buttons)
	}

	fmt.Println("Part 2:", int(totalPresses))
}

func solve(joltageRequirements []float64, buttons [][]int) float64 {
	numRows := len(joltageRequirements)
	numCols := len(buttons)
	matrix := matrix(joltageRequirements, buttons)

	pivotCols := gaussianElimination(matrix, numRows, numCols)
	printMatrix(matrix)

	isPivot := make(map[int]bool)
	pivotRowMap := make(map[int]int)
	for r, c := range pivotCols {
		if c != -1 {
			isPivot[c] = true
			pivotRowMap[c] = r
		}
	}

	var freeCols []int
	for col := 0; col < numCols; col++ {
		if !isPivot[col] {
			freeCols = append(freeCols, col)
		}
	}

	minTotal := math.MaxFloat64

	maxSearchLimit := 0.0
	for _, t := range joltageRequirements {
		if t > maxSearchLimit {
			maxSearchLimit = t
		}
	}

	return search(matrix, maxSearchLimit, minTotal, numCols, numRows, isPivot, pivotRowMap, freeCols, 0, make(map[int]float64))
}

func matrix(joltageRequirements []float64, buttons [][]int) [][]float64 {
	numRows := len(joltageRequirements)
	numCols := len(buttons)

	matrix := make([][]float64, numRows)
	for i := range matrix {
		matrix[i] = make([]float64, numCols+1)
		matrix[i][numCols] = joltageRequirements[i]
	}

	for col, btnIndices := range buttons {
		for _, rowIdx := range btnIndices {
			if rowIdx < numRows {
				matrix[rowIdx][col] = 1.0
			}
		}
	}
	return matrix
}

func search(matrix [][]float64, maxSearchLimit, minTotal float64, numCols, numRows int, isPivot map[int]bool, pivotRowMap map[int]int, freeCols []int, idx int, currentFreeVals map[int]float64) float64 {
	if idx == len(freeCols) {
		currentSolution := make(map[int]float64)
		currentSum := 0.0

		for col, val := range currentFreeVals {
			currentSolution[col] = val
			currentSum += val
		}

		isValid := true

		for col := 0; col < numCols; col++ {
			if isPivot[col] {
				row := pivotRowMap[col]

				sumFree := 0.0
				for _, fCol := range freeCols {
					sumFree += matrix[row][fCol] * currentFreeVals[fCol]
				}

				val := matrix[row][numCols] - sumFree

				if val < -0.0001 {
					isValid = false
					break
				}

				nearest := math.Round(val)
				if math.Abs(val-nearest) > 0.0001 {
					isValid = false
					break
				}

				currentSolution[col] = nearest
				currentSum += nearest
			}
		}

		if isValid {
			if currentSum < minTotal {
				minTotal = currentSum
			}
		}
		return minTotal
	}

	fCol := freeCols[idx]

	for val := 0.0; val <= maxSearchLimit; val++ {
		if minTotal != math.MaxFloat64 && val > minTotal {
			break
		}

		currentFreeVals[fCol] = val
		temp := search(matrix, maxSearchLimit, minTotal, numCols, numRows, isPivot, pivotRowMap, freeCols, idx+1, currentFreeVals)
		if temp < minTotal {
			minTotal = temp
		}
	}
	return minTotal
}

func gaussianElimination(m [][]float64, rows, cols int) []int {
	pivotRow := 0
	pivotCols := make([]int, rows)
	for i := range pivotCols {
		pivotCols[i] = -1
	}

	for col := 0; col < cols && pivotRow < rows; col++ {
		selectedRow := pivotRow
		// search for row with max value in this column
		for i := pivotRow + 1; i < rows; i++ {
			if math.Abs(m[i][col]) > math.Abs(m[selectedRow][col]) {
				selectedRow = i
			}
		}

		// ignore column if all values are zero
		if math.Abs(m[selectedRow][col]) <= 0 {
			continue
		}

		// swap rows
		m[pivotRow], m[selectedRow] = m[selectedRow], m[pivotRow]
		pivotCols[pivotRow] = col

		div := m[pivotRow][col]
		// divide pivot row by largest value
		for j := col; j <= cols; j++ {
			m[pivotRow][j] /= div
		}

		// eliminate other rows
		for i := 0; i < rows; i++ {
			if i != pivotRow {
				mult := m[i][col]
				for j := col; j <= cols; j++ {
					m[i][j] -= mult * m[pivotRow][j]
				}
			}
		}
		pivotRow++
	}
	return pivotCols
}

func parseMachine(input string) Machine {
	startBrace := strings.Index(input, "{")
	endBrace := strings.Index(input, "}")
	targetStr := input[startBrace+1 : endBrace]

	joltages := []int{}
	for _, s := range strings.Split(targetStr, ",") {
		val, _ := strconv.Atoi(strings.TrimSpace(s))
		joltages = append(joltages, val)
	}

	buttons := [][]int{}
	tokens := strings.Split(input, " ")
	for _, t := range tokens {
		if strings.HasPrefix(t, "(") && strings.HasSuffix(t, ")") {
			content := t[1 : len(t)-1]
			btn := []int{}
			if len(content) > 0 {
				for _, numStr := range strings.Split(content, ",") {
					idx, _ := strconv.Atoi(strings.TrimSpace(numStr))
					btn = append(btn, idx)
				}
			}
			buttons = append(buttons, btn)
		}
	}

	return Machine{
		Buttons:             buttons,
		JoltageRequirements: joltages,
	}
}

type Machine struct {
	Buttons             [][]int
	JoltageRequirements []int
}

func floatSlice(ints []int) []float64 {
	floats := []float64{}
	for _, i := range ints {
		floats = append(floats, float64(i))
	}
	return floats
}

func printMatrix(m [][]float64) {
	fmt.Println("------")
	for _, row := range m {
		fmt.Println(row)
	}
}
