package main

import (
	"fmt"

	"github.com/wimspaargaren/aoc"
)

func main() {
	recursiveStore := make(map[string]int)
	part1Visited := make(map[string]bool)
	input := aoc.MustReadFile("input.txt")
	grid := aoc.ParseGrid(input)
	part1, part2 := walkGrid(grid, recursiveStore, part1Visited, 1)
	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func walkGrid(grid [][]string, recursiveStore map[string]int, part1Visited map[string]bool, yIn int) (int, int) {
	part1 := 0
	part2 := 0
	for y := yIn; y < len(grid)-1; y++ {
		for x := 0; x < len(grid[y])-1; x++ {
			if shouldBeam(grid[y-1][x]) {
				switch grid[y][x] {
				case ".":
					grid[y][x] = "|"
				case "^":
					if !part1Visited[fmt.Sprintf("%d%d", y, x)] {
						part1++
						part1Visited[fmt.Sprintf("%d%d", y, x)] = true
					}
					if hasRight(x) {
						part1Res, part2Res := recurse(grid, x-1, y, recursiveStore, part1Visited)
						part1 += part1Res
						part2 += part2Res
					}
					if hasLeft(x, len(grid[y])) {
						part1Res, part2Res := recurse(grid, x+1, y, recursiveStore, part1Visited)
						part1 += part1Res
						part2 += part2Res
					}
					return part1, part2
				}
			}
		}
	}
	return part1, part2 + 1
}

func recurse(grid [][]string, x, y int, recursiveStore map[string]int, part1Visited map[string]bool) (int, int) {
	tile := grid[y][x]
	if tile != "." {
		return 0, 0
	}
	v, ok := recursiveStore[fmt.Sprintf("%d%d", y, x)]
	if ok {
		return 0, v
	}

	newGrid := copyGrid(grid)
	newGrid[y][x] = "|"
	part1Res, part2Res := walkGrid(newGrid, recursiveStore, part1Visited, y+1)
	recursiveStore[fmt.Sprintf("%d%d", y, x)] = part2Res
	return part1Res, part2Res
}

func shouldBeam(tile string) bool {
	return tile == "S" || tile == "|"
}

func hasRight(x int) bool {
	return x > 0
}

func hasLeft(x int, length int) bool {
	return x < length-1
}

func copyGrid(grid [][]string) [][]string {
	newGrid := make([][]string, len(grid))
	for i := range grid {
		newGrid[i] = make([]string, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}
