package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkIntMin(b *testing.B) {
	input := `6 11 33023 4134 564 0 8922422 688775`
	parsed := parseInput(input)

	for i := 0; i < b.N; i++ {
		resMap := make(map[string]int)

		total := 0
		for i := 0; i < len(parsed); i++ {
			total += recurse(parsed[i], 75, resMap)
		}
	}
}

func TestSolution(t *testing.T) {
	input := `6 11 33023 4134 564 0 8922422 688775`
	resMap := make(map[string]int)

	parsed := parseInput(input)
	part1 := 0
	part2 := 0
	for i := 0; i < len(parsed); i++ {
		part1 += recurse(parsed[i], 25, resMap)
	}
	for i := 0; i < len(parsed); i++ {
		part2 += recurse(parsed[i], 75, resMap)
	}
	assert.Equal(t, 220999, part1)
	assert.Equal(t, 261936432123724, part2)

}
