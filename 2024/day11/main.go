package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wimspaargaren/aoc"
)

func main() {
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

	fmt.Println("part1", part1)
	fmt.Println("part2", part2)
}

func recurse(parsed, n int, resMap map[string]int) int {
	if n == 0 {
		return 1
	}
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(parsed))
	sb.WriteString(":")
	sb.WriteString(strconv.Itoa(n))
	key := sb.String()

	res, ok := resMap[key]
	if ok {
		return res
	}
	rulesApplied := applyRules(parsed)

	total := 0
	for i := 0; i < len(rulesApplied); i++ {
		total += recurse(rulesApplied[i], n-1, resMap)
	}

	resMap[key] = total
	return total
}

func applyRules(n int) []int {
	if n == 0 {
		return []int{1}
	} else if isEven(n) {
		first, second := splitHalf(n)
		return []int{first, second}
	} else {
		return []int{n * 2024}
	}
}

func splitHalf(n int) (int, int) {
	word := strconv.Itoa(n)
	half := len(word) / 2
	firstHalf := aoc.MustParseInt(word[:half])
	secondHalf := aoc.MustParseInt(word[half:])
	return firstHalf, secondHalf
}

func isEven(n int) bool {
	return len(strconv.Itoa(n))%2 == 0
}

func parseInput(input string) []int {
	split := strings.Split(input, " ")
	var res []int
	for _, s := range split {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res = append(res, i)
	}
	return res
}
