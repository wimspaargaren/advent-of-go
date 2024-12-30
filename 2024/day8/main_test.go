package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wimspaargaren/aoc"
)

func TestSolution(t *testing.T) {
	input := aoc.MustReadFile("input.txt")
	parsedMap, pointMap := parseMap(input)
	part1, part2 := processPointMap(parsedMap, pointMap)
	assert.Equal(t, 252, part1)
	assert.Equal(t, 839, part2)
}
