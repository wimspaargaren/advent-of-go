package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wimspaargaren/aoc"
)

type Calc struct {
	Values       []string
	Operator     string
	ColumnLength int
}

func main() {
	input := aoc.MustReadFile("input.txt")
	lines := strings.Split(input, "\n")

	columnLengths := getColumnLengths(lines)

	calcs := getCalcs(lines, columnLengths)

	part1 := 0
	for _, calc := range calcs {
		part1 += getTotal(calc.Values, calc.Operator)
	}

	part2 := 0
	for _, calc := range calcs {
		newToCalc := make([]string, calc.ColumnLength)
		for _, val := range calc.Values {
			for i, c := range val {
				if c == ' ' {
					continue
				}
				newToCalc[i] += string(c)
			}
		}
		part2 += getTotal(newToCalc, calc.Operator)
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func getTotal(values []string, operator string) int {
	total := 0
	for _, val := range values {
		val = strings.Trim(val, " ")
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		switch operator {
		case "+":
			total += num
		case "*":
			if total == 0 {
				total = 1
			}
			total *= num
		default:
			panic("unknown operator")
		}
	}
	return total
}

func getCalcs(lines []string, columnLengths []int) []Calc {
	calcs := make([]Calc, len(columnLengths))
	for i, cLen := range columnLengths {
		calcs[i].ColumnLength = cLen
	}
	for lineIndex, line := range lines {
		line = line + " "
		if lineIndex == len(lines)-1 {
			start := 0
			for i, cLen := range columnLengths {
				end := start + cLen + 1
				calcs[i].Operator = strings.Trim(line[start:end], " ")
				start = end
			}
			break
		}
		start := 0
		for i, cLen := range columnLengths {
			end := start + cLen + 1
			calcs[i].Values = append(calcs[i].Values, line[start:end])
			start = end
		}
	}
	return calcs
}

func getColumnLengths(lines []string) []int {
	columnLengths := []int{}
	for iLine, line := range lines {
		if iLine == len(lines)-1 {
			break
		}

		numbers := strings.Fields(line)
		for i, number := range numbers {

			if iLine == 0 {
				columnLengths = append(columnLengths, len(number))
				continue
			} else {
				if len(number) > columnLengths[i] {
					columnLengths[i] = len(number)
				}
			}
		}
	}
	total := 0
	for _, clen := range columnLengths {
		total += clen + 1
	}
	if total != len(lines[0])+1 {
		panic("column length calculation is off")
	}
	return columnLengths
}
