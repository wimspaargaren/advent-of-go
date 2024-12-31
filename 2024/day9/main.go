package main

import (
	"fmt"

	"github.com/wimspaargaren/aoc"
)

func main() {
	input := aoc.MustReadFile("input.txt")
	inputPart1, ids, idIndexMap := inputParser(input)
	inputPart2 := aoc.CopySlice(inputPart1)

	resPart1 := moveToEmptySpace(inputPart1)
	resPart2 := moveToEmptySpaceIfFits(inputPart2, ids, idIndexMap)

	fmt.Println("part1:", countResult(resPart1))
	fmt.Println("part2:", countResult(resPart2))
}

func countResult(input []string) int {
	total := 0
	for i := 0; i < len(input); i++ {
		if input[i] == "." {
			continue
		}
		integer := aoc.MustParseInt(string(input[i]))
		total += integer * i
	}
	return total
}

func moveToEmptySpace(input []string) []string {
	start := 0
	end := len(input) - 1
	for start < end-1 {
		for input[start] != "." {
			start++
		}
		for input[end] == "." {
			end--
		}
		input[start], input[end] = input[end], input[start]
		start++
		end--
	}

	return input
}

func moveToEmptySpaceIfFits(input []string, ids int, idIndexMap map[int][]int) []string {
	for id := ids; id >= 0; id-- {
		indexes := idIndexMap[id]
		firstEmptySpaceIndex := getFirstEmptySpacesIndexForLength(input, indexes[0], len(indexes))
		if firstEmptySpaceIndex == -1 {
			continue
		}
		counter := 0
		for i := firstEmptySpaceIndex; i < firstEmptySpaceIndex+len(indexes); i++ {
			input[i] = input[indexes[counter]]
			counter++
		}
		for _, index := range indexes {
			input[index] = "."
		}
	}

	return input
}

func getFirstEmptySpacesIndexForLength(input []string, before, length int) int {
	for i := 0; i < before; i++ {
		if input[i] == "." {
			isEnough := true
			for j := i; j < i+length; j++ {
				if j > len(input)-1 {
					isEnough = false
					break
				}
				if input[j] != "." {
					isEnough = false
				}
			}
			if isEnough {
				return i
			}
		}
	}
	return -1
}

func inputParser(input string) ([]string, int, map[int][]int) {
	idIndexMap := map[int][]int{}
	resultString := []string{}
	idCount := 0
	resultIndexCounter := 0
	for i := 0; i < len(input); i++ {
		integer := aoc.MustParseInt(string(input[i]))
		char := "."
		if i%2 == 0 {
			char = fmt.Sprintf("%d", idCount)
			idCount++
		}
		ids := []int{}
		for j := 0; j < integer; j++ {
			resultString = append(resultString, char)
			ids = append(ids, resultIndexCounter)
			resultIndexCounter++
		}
		if i%2 == 0 {
			idIndexMap[idCount-1] = ids
		}
	}
	return resultString, idCount - 1, idIndexMap
}
