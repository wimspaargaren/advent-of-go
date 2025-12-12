package main

import (
	"fmt"
	"strings"

	"github.com/wimspaargaren/aoc"
)

func main() {
	input := aoc.MustReadFile("input.txt")
	graph := parseInput(input)
	part1 := part1("you", "out", graph, map[string]int{})
	fmt.Println("Part 1:", part1)
	start := "svr"
	end := "out"
	part2 := pathsTo(start, end, false, false, graph, map[node]int{})
	fmt.Println("Part 2:", part2)
}

type node struct {
	val    string
	hasFFT bool
	hasDAC bool
}

func part1(current string, target string, graph map[string][]string, visited map[string]int) int {
	if current == target {
		return 1
	}

	visitedCount, ok := visited[current]
	if ok {
		return visitedCount
	}

	next := graph[current]
	total := 0
	for _, n := range next {
		total += part1(n, target, graph, visited)
	}
	visited[current] = total
	return total
}

func pathsTo(current string, target string, hasFFt, hasDAC bool, graph map[string][]string, visited map[node]int) int {
	if current == "fft" {
		hasFFt = true
	}
	if current == "dac" {
		hasDAC = true
	}

	if current == target {
		if hasDAC && hasFFt {
			return 1
		}
		return 0
	}

	visitedCount, ok := visited[node{val: current, hasFFT: hasFFt, hasDAC: hasDAC}]
	if ok {
		return visitedCount
	}

	next := graph[current]
	total := 0
	for _, n := range next {
		total += pathsTo(n, target, hasFFt, hasDAC, graph, visited)
	}
	visited[node{val: current, hasFFT: hasFFt, hasDAC: hasDAC}] = total
	return total
}

func parseInput(input string) map[string][]string {
	res := map[string][]string{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		key := parts[0]
		values := strings.Split(parts[1], " ")
		res[key] = values
	}
	return res
}
