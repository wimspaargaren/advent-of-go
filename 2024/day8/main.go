package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/wimspaargaren/aoc"
)

func main() {
	input := aoc.MustReadFile("input.txt")
	parsedMap, pointMap := parseMap(input)
	part1, part2 := processPointMap(parsedMap, pointMap)
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func processPointMap(parsedMap [][]string, pointMap map[string][]Point) (int, int) {
	allCombinations := []Combination{}
	for _, v := range pointMap {
		combinations := getAllCombinations(v)
		allCombinations = append(allCombinations, combinations...)
	}
	part1Map := map[string]bool{}
	part2Map := map[string]bool{}
	for y, row := range parsedMap {
		for x := range row {
			for _, c := range allCombinations {
				// get distance from (x, y) to c.Point1 and c.Point2
				distanceToPoint1 := distance(Point{X: x, Y: y}, c.Point1)
				distanceToPoint2 := distance(Point{X: x, Y: y}, c.Point2)

				if areCollinear(c.Point1, c.Point2, Point{X: x, Y: y}) {
					// and distance between c.Point1 and c.Point2 is double, count x, y as antinode
					if distanceToPoint1 == 2*distanceToPoint2 {
						part1Map[fmt.Sprintf("%d,%d", x, y)] = true
					}
					// part 2 all colinear are antinodes
					part2Map[fmt.Sprintf("%d,%d", x, y)] = true
				}
			}
		}
	}
	return len(part1Map), len(part2Map)
}

func areCollinear(p1, p2, p3 Point) bool {
	return (p2.Y-p1.Y)*(p3.X-p2.X) == (p3.Y-p2.Y)*(p2.X-p1.X)
}

func distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2))
}

type Combination struct {
	Point1 Point
	Point2 Point
}

func (c Combination) PositiveDistance() (int, int) {
	return int(math.Abs(float64(c.Point1.X - c.Point2.X))),
		int(math.Abs(float64(c.Point1.Y - c.Point2.Y)))
}

func getAllCombinations(points []Point) []Combination {
	result := []Combination{}
	for i, p1 := range points {
		for j, p2 := range points {
			if i != j {
				result = append(result, Combination{Point1: p1, Point2: p2})
			}
		}
	}
	return result
}

func isInMap(parsedMap [][]string, x, y int) bool {
	return y >= 0 && y < len(parsedMap) && x >= 0 && x < len(parsedMap[y])
}

func printMap(m [][]string) {
	var res string
	for _, row := range m {
		res += strings.Join(row, "") + "\n"
	}
	fmt.Println(res)
}

type Point struct {
	X, Y int
}

func parseMap(input string) ([][]string, map[string][]Point) {
	var res [][]string
	pointMap := map[string][]Point{}
	for y, line := range strings.Split(input, "\n") {
		xs := strings.Split(line, "")
		row := []string{}
		for x, xV := range xs {
			if xV != "." {
				pointMap[xV] = append(pointMap[xV], Point{X: x, Y: y})
			}
			row = append(row, xV)
		}

		res = append(res, row)
	}
	return res, pointMap
}
