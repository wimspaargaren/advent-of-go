package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wimspaargaren/aoc"
)

type Rotation struct {
	Dir string
	N   int
}

func main() {
	lines := strings.Split(aoc.MustReadFile("input.txt"), "\n")
	rotations := []Rotation{}
	for _, line := range lines {
		rotations = append(rotations, parseInput(line))
	}
	start := 50
	exactZero := 0
	totalZero := 0
	// start is between 0 and 99
	for _, rot := range rotations {
		if rot.Dir == "R" {
			for i := 0; i < rot.N; i++ {
				start++
				if start == 100 {
					totalZero++
					start = 0
				}
			}
		} else {
			for i := 0; i < rot.N; i++ {
				start--
				if start == 0 {
					totalZero++
				}
				if start == -1 {
					start = 99
				}
			}
		}
		if start == 0 {
			exactZero++
		}
	}
	fmt.Println("part 1: ", exactZero)
	fmt.Println("part 2: ", totalZero)
}

func parseInput(line string) Rotation {
	dir := string(line[0])
	n, err := strconv.Atoi(line[1:])
	if err != nil {
		panic(err)
	}

	return Rotation{
		Dir: dir,
		N:   n,
	}
}
