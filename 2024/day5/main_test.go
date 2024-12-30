package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	part1, part2 := solve()
	assert.Equal(t, 5762, part1)
	assert.Equal(t, 4130, part2)
}
