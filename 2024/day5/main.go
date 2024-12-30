package main

import (
	"fmt"
	"strings"

	"github.com/wimspaargaren/aoc"
)

func main() {
	part1, part2 := solve()
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func solve() (int, int) {
	input := aoc.MustReadFile("input.txt")
	tuples, rows := parseInput(input)

	total := 0
	total2 := 0

	for _, row := range rows {
		isOk := isOk(row, tuples)
		if !isOk {
			correctly := orderCorrectly(row, tuples)
			total2 += correctly[int(len(correctly)/2)]
		}
		if isOk {
			total += row[int(len(row)/2)]
		}
	}
	return total, total2
}

func isOk(row []int, tuples map[string]bool) bool {
	for i := 0; i < len(row); i++ {
		for j := i + 1; j < len(row); j++ {
			if !tuples[fmt.Sprintf("%d|%d", row[i], row[j])] {
				return false
			}
		}
	}
	return true
}

func orderCorrectly(row []int, tuples map[string]bool) []int {
	res := []int{}
	for i := 0; i < len(row); i++ {
		for j := i + 1; j < len(row); j++ {
			if !tuples[fmt.Sprintf("%d|%d", row[i], row[j])] {
				res = append(res, row[j], row[i])
				for k := i; k < len(row); k++ {
					if k == i || k == j {
						continue
					}
					res = append(res, row[k])
				}
				return orderCorrectly(res, tuples)
			}
		}
		res = append(res, row[i])
	}
	return res
}

func parseInput(input string) (map[string]bool, [][]int) {
	splitted := strings.Split(input, "\n")
	tuples := map[string]bool{}
	for _, s := range splitted {
		if strings.Contains(s, "|") {
			tuples[s] = true
		}
	}

	rows := [][]int{}
	for _, line := range splitted {
		if strings.Contains(line, ",") {
			row := []int{}
			pages := strings.Split(line, ",")
			for i := 0; i < len(pages); i++ {
				row = append(row, aoc.MustParseInt(pages[i]))
			}
			rows = append(rows, row)
		}
	}
	return tuples, rows
}
