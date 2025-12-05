package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/wimspaargaren/aoc"
)

type Range struct {
	start int
	end   int
}

func (r Range) Contains(value int) bool {
	return value >= r.start && value <= r.end
}

func main() {
	input := aoc.MustReadFile("input.txt")

	lines := strings.Split(string(input), "\n")
	isRanges := true
	ranges := []Range{}
	toCheck := []int{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			isRanges = false
			continue
		}

		if isRanges {
			parts := strings.Split(line, "-")
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			end, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			ranges = append(ranges, Range{start: start, end: end})
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			toCheck = append(toCheck, value)
		}
	}

	freshCount := 0
	for _, check := range toCheck {
		for _, r := range ranges {
			if r.Contains(check) {
				freshCount++
				break
			}
		}
	}

	mergedRanges := mergeRanges(ranges)
	total := countRanges(mergedRanges)

	fmt.Println("part 1:", freshCount)
	fmt.Println("part 2:", total)
}

func countRanges(ranges []Range) int {
	count := 0
	for _, r := range ranges {
		count += r.end - r.start + 1
	}
	return count
}

func mergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return []Range{}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	merged := []Range{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		current := ranges[i]
		lastMerged := merged[len(merged)-1]

		if current.start <= lastMerged.end {
			if current.end >= lastMerged.start && current.end > lastMerged.end {
				merged[len(merged)-1].end = current.end
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}
