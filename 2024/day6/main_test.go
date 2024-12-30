package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wimspaargaren/aoc"
)

func TestExample(t *testing.T) {
	inputTest := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	expected := 41
	_, result := walkGuard(aoc.ParseGrid(inputTest))
	assert.Equal(t, expected, result)
}
func TestOther(t *testing.T) {
	inputTest := `..........
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	expected := 7
	_, result := walkGuard(aoc.ParseGrid(inputTest))
	assert.Equal(t, expected, result)
}

func TestReal(t *testing.T) {
	input := aoc.MustReadFile("input.txt")
	grid := aoc.ParseGrid(input)
	_, result := walkGuard(grid)
	assert.Equal(t, 4663, result)
}

func TestRealPart2(t *testing.T) {
	input := aoc.MustReadFile("input.txt")
	grid := aoc.ParseGrid(input)

	robotWalkCoords, _ := walkGuard(grid)

	result := totalLoopPositions(grid, robotWalkCoords)
	assert.Equal(t, 1530, result)
}
