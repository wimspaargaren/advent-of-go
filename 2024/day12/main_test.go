package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wimspaargaren/aoc"
)

func TestSolution(t *testing.T) {
	input := aoc.MustReadFile("input.txt")
	grid := aoc.ParseGrid(input)

	totalCostPart1, totalCostPart2 := solve(grid)
	assert.Equal(t, 1456082, totalCostPart1)
	assert.Equal(t, 872382, totalCostPart2)
}
