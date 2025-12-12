package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wimspaargaren/aoc"
)

func main() {
	input := aoc.MustReadFile("input.txt")
	_ = input
	grids, _ := parseInput(input)
	part1 := 0
	for _, grid := range grids {
		check := 0
		for _, presentID := range grid.Presents {
			check += presentID
		}
		if check*8 > grid.Width*grid.Height {
			continue
		}
		part1++
	}
	fmt.Println("Part 1:", part1)
}

type Present struct {
	Width    int
	Height   int
	Grid     [][]string
	Variants [][][]string
	Area     int
}

type Grid struct {
	Width    int
	Height   int
	Presents []int
}

func parseInput(input string) ([]Grid, []Present) {
	grids, presents := []Grid{}, []Present{}

	isGrids := false
	present := Present{}
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "x") {
			splitted := strings.Split(line, " ")
			dimension := splitted[0]
			dimensions := strings.Split(dimension[0:len(dimension)-1], "x")
			width, err := strconv.Atoi(dimensions[0])
			if err != nil {
				panic(err)
			}
			height, err := strconv.Atoi(dimensions[1])
			if err != nil {
				panic(err)
			}
			presentIDs := []int{}
			for _, idStr := range splitted[1:] {
				id, err := strconv.Atoi(idStr)
				if err != nil {
					panic(err)
				}
				presentIDs = append(presentIDs, id)
			}
			grid := Grid{
				Width:    width,
				Height:   height,
				Presents: presentIDs,
			}
			grids = append(grids, grid)
			isGrids = true
			continue
		}
		if line == "" {
			presents = append(presents, present)
		}
		if strings.HasSuffix(line, ":") {
			present = Present{}
		} else if !isGrids {
			presentRow := strings.Split(line, "")
			present.Grid = append(present.Grid, presentRow)
			present.Height++
			present.Width = len(presentRow)
		}

	}
	for i, present := range presents {
		// create variants (rotations) & flips
		variants := [][][]string{}
		variants = append(variants, present.Grid)
		degree90 := rotate90(present.Grid)
		variants = append(variants, degree90)
		degree180 := rotate90(degree90)
		variants = append(variants, degree180)
		degree270 := rotate90(degree180)
		variants = append(variants, degree270)
		flipped := flip(present.Grid)
		variants = append(variants, flipped)
		flipped90 := rotate90(flipped)
		variants = append(variants, flipped90)
		flipped180 := rotate90(flipped90)
		variants = append(variants, flipped180)
		flipped270 := rotate90(flipped180)
		variants = append(variants, flipped270)
		unique := map[string]bool{}

		for _, v := range variants {
			_, exists := unique[gridToString(v)]
			if !exists {
				unique[gridToString(v)] = true
				presents[i].Variants = append(presents[i].Variants, v)
			}
		}

		// area is always the same
		for _, row := range present.Grid {
			for _, cell := range row {
				if cell == "#" {
					presents[i].Area++
				}
			}
		}
	}

	return grids, presents
}

func rotate90(grid [][]string) [][]string {
	height := len(grid)
	width := len(grid[0])
	newGrid := make([][]string, width)
	for i := range newGrid {
		newGrid[i] = make([]string, height)
	}
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			newGrid[c][height-1-r] = grid[r][c]
		}
	}
	return newGrid
}

func flip(grid [][]string) [][]string {
	height := len(grid)
	width := len(grid[0])
	newGrid := make([][]string, height)
	for i := range newGrid {
		newGrid[i] = make([]string, width)
	}
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			newGrid[r][width-1-c] = grid[r][c]
		}
	}
	return newGrid
}

func gridToString(grid [][]string) string {
	res := ""
	for _, row := range grid {
		res += strings.Join(row, "") + "\n"
	}
	return res
}

func printGrid(grid [][]string) {
	fmt.Println(gridToString(grid))
}
