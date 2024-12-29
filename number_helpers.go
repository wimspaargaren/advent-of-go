package aoc

import (
	"math"
	"strconv"
)

func Abs(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func MustParseInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}
