package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input1, _ := getInput()
	got := execCount(input1)
	assert.Equal(t, 165225049, got)
}

func TestPart2(t *testing.T) {
	_, input2 := getInput()
	got := execCount(input2)
	assert.Equal(t, 108830766, got)
}
