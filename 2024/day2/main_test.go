package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolutions(t *testing.T) {
	part1, part2 := solve()
	assert.Equal(t, 407, part1)
	assert.Equal(t, 459, part2)
}
