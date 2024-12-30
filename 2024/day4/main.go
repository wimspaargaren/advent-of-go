package main

import (
	"fmt"
	"strings"

	"github.com/wimspaargaren/aoc"
)

func main() {
	input := aoc.MustReadFile("./input.txt")

	fmt.Println("part1:", getXMAS(input))
	fmt.Println("part2:", XMAS(aoc.ParseGrid(input)))
}

func getXMAS(input string) int {
	lines := strings.Split(input, "\n")
	grid := aoc.ParseGrid(input)
	diagonals := aoc.AllDiagonals(grid)
	for _, v := range diagonals {
		lines = append(lines, strings.Join(v, ""))
	}

	verticals := aoc.Verticals(grid)
	for _, v := range verticals {
		lines = append(lines, strings.Join(v, ""))
	}

	total := 0
	for _, line := range lines {
		total += strings.Count(line, "XMAS")
		total += strings.Count(aoc.ReverseString(line), "XMAS")
	}
	return total
}

func XMAS(grid [][]string) int {
	total := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[y][x] != "A" {
				continue
			}
			topLeft, ok := aoc.ValOkAt(grid, x-1, y-1)
			if !ok {
				continue
			}
			bottomRight, ok := aoc.ValOkAt(grid, x+1, y+1)
			if !ok {
				continue
			}
			topRight, ok := aoc.ValOkAt(grid, x+1, y-1)
			if !ok {
				continue
			}
			bottomLeft, ok := aoc.ValOkAt(grid, x-1, y+1)
			if !ok {
				continue
			}
			if ((topLeft == "M" && bottomRight == "S") || (topLeft == "S" && bottomRight == "M")) &&
				((topRight == "M" && bottomLeft == "S") || (topRight == "S" && bottomLeft == "M")) {
				total++
			}
		}
	}
	return total
}
