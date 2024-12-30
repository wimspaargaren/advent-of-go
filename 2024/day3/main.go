package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/wimspaargaren/aoc"
)

func main() {
	input1, input2 := getInput()

	fmt.Println("part1:", execCount(input1))
	fmt.Println("part2:", execCount(input2))
}

func getInput() (string, string) {
	input1 := aoc.MustReadFile("input.txt")
	doSplits := strings.Split(input1, "do()")
	input2 := ""
	for _, doSplit := range doSplits {
		dontSplit := strings.Split(doSplit, "don't()")
		input2 += dontSplit[0]
	}
	return input1, input2
}

func execCount(input string) int {
	reg := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	res := reg.FindAll([]byte(input), -1)
	total := 0
	for _, r := range res {
		submatches := reg.FindAllStringSubmatch(string(r), -1)
		x := aoc.MustParseInt(submatches[0][1])
		y := aoc.MustParseInt(submatches[0][2])
		total += x * y
	}
	return total
}
