package main

import (
	"fmt"

	"github.com/wimspaargaren/aoc"
)

func main() {
	input := aoc.MustReadFile("input.txt")
	grid := aoc.ParseGrid(input)

	robotWalkCoords, result := walkGuard(grid)
	fmt.Println("part1:", result)
	fmt.Println("part2:", totalLoopPositions(grid, robotWalkCoords))
}

func totalLoopPositions(grid [][]string, visits map[GridCoord]bool) int {
	total := 0
	guardX, guardY := getGuardPos(grid)

	for k := range visits {
		if grid[k.y][k.x] == "^" {
			continue
		}

		grid[k.y][k.x] = "#"
		hasLoop, _ := walk(grid, guardX, guardY, "up", true)
		if hasLoop {
			total++
		}
		grid[k.y][k.x] = "."
	}

	return total
}

func walkGuard(grid [][]string) (map[GridCoord]bool, int) {
	guardX, guardY := getGuardPos(grid)
	facing := "up"
	_, visits := walk(grid, guardX, guardY, facing, false)
	distinctCoords := map[GridCoord]bool{}
	for k := range visits {
		distinctCoords[GridCoord{x: k.x, y: k.y}] = true
	}
	return distinctCoords, len(distinctCoords)
}

type GridCoord struct {
	x, y int
}

type TilePos struct {
	x, y   int
	facing string
}

func walk(grid [][]string, guardX, guardY int, facing string, returnOnLoop bool) (bool, map[TilePos]bool) {
	visitsMap := map[TilePos]bool{}

	for isInMap(guardX, guardY, grid) {
		tilePos := TilePos{
			x:      guardX,
			y:      guardY,
			facing: facing,
		}
		if visitsMap[tilePos] && returnOnLoop {
			return true, visitsMap
		}
		visitsMap[tilePos] = true

		tempX := guardX
		tempY := guardY
		switch facing {
		case "up":
			tempX, tempY = aoc.ToUp(guardX, guardY)
		case "right":
			tempX, tempY = aoc.ToRight(guardX, guardY)
		case "down":
			tempX, tempY = aoc.ToDown(guardX, guardY)
		case "left":
			tempX, tempY = aoc.ToLeft(guardX, guardY)
		}
		if !aoc.IsInGrid(grid, tempX, tempY) {
			return false, visitsMap
		}
		if aoc.ValAt(grid, tempX, tempY) == "#" {
			state := changeState(grid, facing, guardX, guardY)

			switch state {
			case "right":
				facing = "right"
				guardX++
			case "down":
				facing = "down"
				guardY++
			case "left":
				facing = "left"
				guardX--
			case "up":
				facing = "up"
				guardY--
			}
		} else {
			guardX = tempX
			guardY = tempY
		}
	}
	return false, visitsMap
}

func changeState(mapGuard [][]string, facing string, x, y int) string {
	switch facing {
	case "up":
		if mapGuard[y][x+1] == "#" {
			return changeState(mapGuard, "right", x, y)
		}
		return "right"
	case "right":
		if mapGuard[y+1][x] == "#" {
			return changeState(mapGuard, "down", x, y)
		}
		return "down"
	case "down":
		if mapGuard[y][x-1] == "#" {
			return changeState(mapGuard, "left", x, y)
		}
		return "left"
	case "left":
		if mapGuard[y-1][x] == "#" {
			return changeState(mapGuard, "up", x, y)
		}
		return "up"
	}
	return ""
}

func isInMap(x, y int, grid [][]string) bool {
	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
		return true
	}
	return false
}

func getGuardPos(grid [][]string) (int, int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "^" {
				return j, i
			}
		}
	}
	return -1, -1
}
