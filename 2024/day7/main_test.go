package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	totalPart1, totalPart2 := solve()
	assert.Equal(t, 12553187650171, totalPart1)
	assert.Equal(t, 96779702119491, totalPart2)
}
