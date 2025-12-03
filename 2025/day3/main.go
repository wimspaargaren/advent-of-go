package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wimspaargaren/aoc"
)

func main() {
	input := aoc.MustReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	part1 := 0
	part2 := 0
	for _, line := range lines {
		numbers := []int{}
		for _, char := range line {
			n, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, n)
		}

		part1 += largest("", 0, 2, numbers)
		part2 += largest("", 0, 12, numbers)
	}
	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func largest(total string, startIndex int, totalLength int, numbers []int) int {
	index := 0
	maxNumber := 0
	for i := startIndex; i < len(numbers)-(totalLength-len(total)-1); i++ {
		if numbers[i] > maxNumber {
			maxNumber = numbers[i]
			index = i + 1
		}
	}
	total += strconv.Itoa(maxNumber)
	if len(total) == totalLength {
		n, err := strconv.Atoi(total)
		if err != nil {
			panic(err)
		}
		return n
	}

	return largest(total, index, totalLength, numbers)
}
