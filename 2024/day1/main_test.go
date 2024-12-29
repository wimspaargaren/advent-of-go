package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolutions(t *testing.T) {
	part1, part2 := solve()
	assert.Equal(t, 1110981, part1)
	assert.Equal(t, 24869388, part2)
}
