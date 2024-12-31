package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wimspaargaren/aoc"
)

func TestPart1(t *testing.T) {
	input := aoc.MustReadFile("input.txt")
	grid := aoc.ParseGrid(input)
	part1 := findTrailsInGrid(grid, true)
	assert.Equal(t, 496, part1)
}

func TestPart2(t *testing.T) {
	input := aoc.MustReadFile("input.txt")
	grid := aoc.ParseGrid(input)
	part2 := findTrailsInGrid(grid, false)
	assert.Equal(t, 1120, part2)
}
