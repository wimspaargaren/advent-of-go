package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wimspaargaren/aoc"
)

func TestPart1(t *testing.T) {
	input := aoc.MustReadFile("input.txt")
	inputPart1, _, _ := inputParser(input)

	resPart1 := moveToEmptySpace(inputPart1)

	assert.Equal(t, 6435922584968, countResult(resPart1))
}

func TestPart2(t *testing.T) {
	input := aoc.MustReadFile("input.txt")
	inputPart2, ids, idIndexMap := inputParser(input)

	resPart2 := moveToEmptySpaceIfFits(inputPart2, ids, idIndexMap)

	assert.Equal(t, 6469636832766, countResult(resPart2))
}
