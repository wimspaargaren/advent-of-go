package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wimspaargaren/aoc"
)

func TestPart1(t *testing.T) {
	input := aoc.MustReadFile("./input.txt")
	got := getXMAS(input)
	assert.Equal(t, 2567, got)
}

func TestPart2(t *testing.T) {
	input := aoc.MustReadFile("./input.txt")
	got := XMAS(aoc.ParseGrid(input))
	assert.Equal(t, 2029, got)
}
