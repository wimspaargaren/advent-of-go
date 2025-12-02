package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wimspaargaren/aoc"
)

func main() {
	input := aoc.MustReadFile("input.txt")
	inputs := strings.Split(input, ",")

	totalPart1 := 0
	totalPart2 := 0

	for i := range inputs {
		parts := strings.Split(inputs[i], "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		for i := start; i <= end; i++ {
			s := fmt.Sprintf("%d", i)
			if checkAnyRepeated(s) {
				totalPart2 += i
			}
			if checkPalinedrome(s) {
				totalPart1 += i
			}
		}

	}
	fmt.Println("Part 1:", totalPart1)
	fmt.Println("Part 2:", totalPart2)
}

func checkPalinedrome(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	first := s[:len(s)/2]
	second := s[len(s)/2:]
	return first == second
}

func checkAnyRepeated(s string) bool {
	for i := 1; i < len(s); i++ {
		sub := s[:i]
		for len(sub) < len(s) {
			sub += s[:i]
		}
		if sub == s {
			return true
		}
	}
	return false
}
