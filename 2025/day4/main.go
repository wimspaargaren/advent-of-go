package main

import (
	"fmt"

	"github.com/wimspaargaren/aoc"
)

func main() {
	input := aoc.MustReadFile("input.txt")
	grid := aoc.ParseGrid(input)
	part1 := totalRollsToBeRemoved(grid, false)

	rollsRemovedTotal := 0
	for {
		rollsAllowedToRemove := totalRollsToBeRemoved(grid, true)
		if rollsAllowedToRemove == 0 {
			break
		}
		rollsRemovedTotal += rollsAllowedToRemove
	}
	fmt.Println("part 1: ", part1)
	fmt.Println("part 2: ", rollsRemovedTotal)
}

func totalRollsToBeRemoved(grid [][]string, withRemoval bool) int {
	rollsAllowedToRemove := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == "." {
				continue

			}

			adjacentRolls := 0
			for _, a := range aoc.AdjacentPositionsForGrid(grid, x, y) {
				if a == "@" {
					adjacentRolls++
				}
			}
			if adjacentRolls < 4 {
				rollsAllowedToRemove++
				if withRemoval {
					grid[y][x] = "."
				}
			}

		}
	}
	return rollsAllowedToRemove
}
