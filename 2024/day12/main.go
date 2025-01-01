package main

import (
	"fmt"

	"github.com/wimspaargaren/aoc"
)

type Fence struct {
	Tile  string
	Count int
	X, Y  int
}

func main() {
	input := aoc.MustReadFile("input.txt")
	grid := aoc.ParseGrid(input)

	totalCostPart1, totalCostPart2 := solve(grid)
	fmt.Println("part1:", totalCostPart1)
	fmt.Println("part2:", totalCostPart2)
}

func solve(grid [][]string) (int, int) {
	resultStore := make(map[string][]Fence)

	for y := 0; y < len(grid[0]); y++ {
		for x := 0; x < len(grid); x++ {
			tile := grid[y][x]
			fenceCount := checkFenceTile(grid, tile, x, y)
			resultStore[tile] = append(resultStore[tile], Fence{Count: fenceCount, X: x, Y: y, Tile: tile})
		}
	}

	totalCostPart1 := 0
	totalCostPart2 := 0
	for _, v := range resultStore {
		groupedResults := groupResult(v)
		for _, group := range groupedResults {
			fences := 0
			for _, fence := range group {
				fences += fence.Count
			}
			totalCostPart1 += fences * len(group)

			sides := sidesForGroup(grid, group)
			totalCostPart2 += sides * len(group)
		}
	}
	return totalCostPart1, totalCostPart2
}

func sidesForGroup(grid [][]string, group []Fence) int {
	xMap := map[int][]Fence{}
	yMap := map[int][]Fence{}
	for _, fence := range group {
		xGroup2, yGroup2 := getXAndYFenceGroups(grid, fence.Tile, fence.X, fence.Y)
		xMap[fence.Y] = append(xMap[fence.Y], xGroup2...)
		yMap[fence.X] = append(yMap[fence.X], yGroup2...)
	}

	totalSides := 0
	for _, xGroup := range xMap {
		totalSides += len(groupResult(xGroup))
	}
	for _, yGroup := range yMap {
		totalSides += len(groupResult(yGroup))
	}

	return totalSides
}

func groupResult(fences []Fence) [][]Fence {
	if len(fences) == 0 {
		return [][]Fence{}
	}
	groups := [][]Fence{
		{fences[0]},
	}
	for i := 1; i < len(fences); i++ {
		notInGroup := true
		for j := 0; j < len(groups); j++ {
			if isInGroup(fences[i], groups[j]) {
				notInGroup = false
				groups[j] = append(groups[j], fences[i])
				break
			}
		}
		if notInGroup {
			groups = append(groups, []Fence{fences[i]})
		}
	}

	return combineGroups(groups)
}

func combineGroups(groups [][]Fence) [][]Fence {
	combined := [][]Fence{}
	addedGroups := []int{}
	for i := 0; i < len(groups); i++ {
		if aoc.Contains(addedGroups, i) {
			continue
		}
		currentGroup := groups[i]
		for j := 0; j < len(groups); j++ {
			if i == j || aoc.Contains(addedGroups, j) {
				continue
			}
			if isAdjacentGroups(groups[i], groups[j]) {
				addedGroups = append(addedGroups, i, j)
				currentGroup = append(currentGroup, groups[j]...)

			}
		}
		combined = append(combined, currentGroup)
	}

	if len(combined) != len(groups) {
		return combineGroups(combined)
	}
	return combined
}

func isAdjacentGroups(group1, group2 []Fence) bool {
	for _, fenceGroup1 := range group1 {
		if isInGroup(fenceGroup1, group2) {
			return true
		}
	}
	return false
}

func isInGroup(fence Fence, group []Fence) bool {
	for _, f := range group {
		if isAdjacentFence(fence, f) {
			return true
		}
	}
	return false
}

func isAdjacentFence(fence1, fence2 Fence) bool {
	if fence1.X == fence2.X-1 && fence1.Y == fence2.Y {
		return true
	}
	if fence1.X == fence2.X+1 && fence1.Y == fence2.Y {
		return true
	}
	if fence1.Y == fence2.Y-1 && fence1.X == fence2.X {
		return true
	}
	if fence1.Y == fence2.Y+1 && fence1.X == fence2.X {
		return true
	}
	return false
}

func getXAndYFenceGroups(grid [][]string, tile string, x, y int) ([]Fence, []Fence) {
	xGroup := []Fence{}
	yGroup := []Fence{}
	if needsFence(grid, tile, x-1, y) {
		yGroup = append(yGroup, Fence{X: x - 1, Y: y})
	}
	if needsFence(grid, tile, x+1, y) {
		yGroup = append(yGroup, Fence{X: x + 1, Y: y})
	}
	if needsFence(grid, tile, x, y-1) {
		xGroup = append(xGroup, Fence{X: x, Y: y - 1})
	}
	if needsFence(grid, tile, x, y+1) {
		xGroup = append(xGroup, Fence{X: x, Y: y + 1})
	}
	return xGroup, yGroup
}

func checkFenceTile(grid [][]string, tile string, x, y int) int {
	fenceCount := 0
	if needsFence(grid, tile, x-1, y) {
		fenceCount++
	}
	if needsFence(grid, tile, x+1, y) {
		fenceCount++
	}
	if needsFence(grid, tile, x, y-1) {
		fenceCount++
	}
	if needsFence(grid, tile, x, y+1) {
		fenceCount++
	}
	return fenceCount
}

func needsFence(grid [][]string, tile string, x, y int) bool {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return true
	}
	return grid[y][x] != tile
}
