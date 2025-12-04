package aoc

import (
	"strings"
)

func ParseGrid(input string) [][]string {
	var grid [][]string
	rowCount := -1
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		row := []string{}
		for _, char := range line {
			row = append(row, string(char))
		}
		if rowCount == -1 {
			rowCount = len(row)
		}
		if len(row) != rowCount {
			panic("Rows are not of equal length")
		}
		grid = append(grid, row)
	}

	return grid
}

func AdjacentPositionsForGrid(grid [][]string, x, y int) []string {
	adjacent := []string{}

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[ny]) {
				adjacent = append(adjacent, grid[ny][nx])
			}
		}
	}

	return adjacent
}

func IsNearEdge(grid [][]string, x, y int) bool {
	if x == 0 || y == 0 {
		return true
	}

	if x == len(grid[0])-1 || y == len(grid)-1 {
		return true
	}

	return false
}

func ValAt(grid [][]string, x, y int) string {
	return grid[y][x]
}

func ValOkAt(grid [][]string, x, y int) (string, bool) {
	if !IsInGrid(grid, x, y) {
		return "", false
	}
	return ValAt(grid, x, y), true
}

func IsInGrid(grid [][]string, x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}

	if x >= len(grid) || y >= len(grid[0]) {
		return false
	}

	return true
}

func ToTopLeft(x, y int) (int, int) {
	return x - 1, y - 1
}

func ToTopRight(x, y int) (int, int) {
	return x + 1, y - 1
}

func ToBottomLeft(x, y int) (int, int) {
	return x - 1, y + 1
}

func ToBottomRight(x, y int) (int, int) {
	return x + 1, y + 1
}

func ToLeft(x, y int) (int, int) {
	return x - 1, y
}

func ToRight(x, y int) (int, int) {
	return x + 1, y
}

func ToUp(x, y int) (int, int) {
	return x, y - 1
}

func ToDown(x, y int) (int, int) {
	return x, y + 1
}

func NegativeSlopeDiagonal(grid [][]string) [][]string {
	diagonals := [][]string{}
	width := len(grid[0])
	height := len(grid)
	for k := 0; k < width+height-1; k++ {
		var diagonal []string
		for x := 0; x <= k; x++ {
			y := k - x
			if y < height && x < width {
				diagonal = append(diagonal, grid[y][x])
			}
		}
		if len(diagonal) > 0 {
			diagonals = append(diagonals, diagonal)
		}
	}
	return diagonals
}

func PositiveSlopeDiagonal(grid [][]string) [][]string {
	diagonals := [][]string{}
	width := len(grid[0])
	height := len(grid)
	for k := 0; k < width+height-1; k++ {
		var diagonal []string
		for x := 0; x <= k; x++ {
			y := k - x
			if y < height && x < width {
				diagonal = append(diagonal, grid[height-y-1][x])
			}
		}
		if len(diagonal) > 0 {
			diagonals = append(diagonals, diagonal)
		}
	}
	return diagonals
}

func AllDiagonals(grid [][]string) [][]string {
	negDiags := NegativeSlopeDiagonal(grid)
	posDiags := PositiveSlopeDiagonal(grid)
	return append(negDiags, posDiags...)
}

func Verticals(grid [][]string) [][]string {
	verticals := [][]string{}
	for i := 0; i < len(grid[0]); i++ {
		vertical := []string{}
		for j := 0; j < len(grid); j++ {
			vertical = append(vertical, grid[j][i])
		}
		verticals = append(verticals, vertical)
	}
	return verticals
}
