package main

import (
	"fmt"

	"github.com/wimspaargaren/aoc"
)

func main() {
	input := aoc.MustReadFile("input.txt")
	grid := aoc.ParseGrid(input)
	part1 := findTrailsInGrid(grid, true)
	fmt.Println("part1:", part1)
	part2 := findTrailsInGrid(grid, false)
	fmt.Println("part2:", part2)
}

func findTrailsInGrid(grid [][]string, uniqueTrails bool) int {
	total := 0
	for y, row := range grid {
		for x, val := range row {
			if val != "0" {
				continue
			}
			visits := map[string]int{}
			paths := findTrails([]Coord{{X: x, Y: y}}, grid, visits, uniqueTrails)
			total += paths
		}
	}
	return total
}

func findTrails(coords []Coord, grid [][]string, visits map[string]int, uniqueTrails bool) int {
	total := 0
	lastElement := coords[len(coords)-1]
	if lastElement.Val == 9 {
		visits[fmt.Sprintf("%d,%d", lastElement.X, lastElement.Y)]++
		return 1
	}
	newCoordsReachable := getReachableTiles(grid, lastElement.X, lastElement.Y)
	if len(newCoordsReachable) == 0 {
		return 0
	}

	for _, c := range newCoordsReachable {
		_, ok := visits[fmt.Sprintf("%d,%d", c.X, c.Y)]
		if ok && uniqueTrails {
			continue
		}
		newCoords := aoc.CopySlice(coords)
		newCoords = append(newCoords, c)
		total += findTrails(newCoords, grid, visits, uniqueTrails)
	}
	return total
}

type Coord struct {
	X, Y int
	Val  int
}

func getReachableTiles(grid [][]string, x, y int) []Coord {
	coords := []Coord{}
	curval := aoc.MustParseInt(grid[y][x])
	nextVal := curval + 1
	if aoc.IsInGrid(grid, x+1, y) && aoc.MustParseInt(grid[y][x+1]) == nextVal {
		coords = append(coords, Coord{X: x + 1, Y: y, Val: aoc.MustParseInt(grid[y][x+1])})
	}
	if aoc.IsInGrid(grid, x-1, y) && aoc.MustParseInt(grid[y][x-1]) == nextVal {
		coords = append(coords, Coord{X: x - 1, Y: y, Val: aoc.MustParseInt(grid[y][x-1])})
	}
	if aoc.IsInGrid(grid, x, y+1) && aoc.MustParseInt(grid[y+1][x]) == nextVal {
		coords = append(coords, Coord{X: x, Y: y + 1, Val: aoc.MustParseInt(grid[y+1][x])})
	}
	if aoc.IsInGrid(grid, x, y-1) && aoc.MustParseInt(grid[y-1][x]) == nextVal {
		coords = append(coords, Coord{X: x, Y: y - 1, Val: aoc.MustParseInt(grid[y-1][x])})
	}
	return coords
}
