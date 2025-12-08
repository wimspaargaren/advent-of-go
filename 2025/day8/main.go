package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/wimspaargaren/aoc"
)

type Point struct {
	x, y, z int
}

// Distance calculates the squared Euclidean distance between three points.
func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	dz := p.z - other.z
	return math.Sqrt(math.Pow(float64(dx), 2) + math.Pow(float64(dy), 2) + math.Pow(float64(dz), 2))
}

func main() {
	input := aoc.MustReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(input), "\n")

	points := getPoints(lines)
	pairs := getPairs(points)

	groups := []Group{}
	individualPoints := map[Point]bool{}

	part1 := 0
	part2 := 0

	for i := range pairs {
		pair := pairs[i]
		individualPoints[pair.a] = true
		individualPoints[pair.b] = true
		if len(points) == len(individualPoints) {
			part2 = pair.a.x * pair.b.x
			break
		}
		indexA := -1
		indexB := -1
		for i, group := range groups {
			_, aExists := group.points[pair.a]
			if aExists {
				indexA = i
			}
			_, bExists := group.points[pair.b]
			if bExists {
				indexB = i
			}
		}
		if indexA == -1 && indexB == -1 {
			newGroup := Group{points: map[Point]bool{}}
			newGroup.AddPoint(pair.a)
			newGroup.AddPoint(pair.b)
			groups = append(groups, newGroup)
		} else if indexA != -1 && indexB == -1 {
			groups[indexA].AddPoint(pair.a)
			groups[indexA].AddPoint(pair.b)
		} else if indexA == -1 && indexB != -1 {
			groups[indexB].AddPoint(pair.a)
			groups[indexB].AddPoint(pair.b)
		} else if indexA != indexB {
			for point := range groups[indexB].points {
				groups[indexA].AddPoint(point)
			}
			groups = append(groups[:indexB], groups[indexB+1:]...)
		}
		if i == 999 {
			part1 = getPart1(groups)
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func getPart1(groups []Group) int {
	slices.SortFunc(groups, func(a, b Group) int {
		if len(a.points) < len(b.points) {
			return 1
		} else if len(a.points) > len(b.points) {
			return -1
		}
		return 0
	})
	res := len(groups[0].points)
	for i := 1; i < 3; i++ {
		res *= len(groups[i].points)
	}
	return res
}

type Group struct {
	points map[Point]bool
}

func (g *Group) AddPoint(p Point) {
	g.points[p] = true
}

type Pair struct {
	a, b     Point
	distance float64
}

func getPairs(points []Point) []Pair {
	pairs := []Pair{}
	added := map[string]bool{}
	for i := range points {
		for j := range points {
			if i == j {
				continue
			}
			key1 := fmt.Sprintf("%d-%d", i, j)
			key2 := fmt.Sprintf("%d-%d", j, i)
			if added[key1] || added[key2] {
				continue
			}
			added[key1] = true
			added[key2] = true
			pairs = append(pairs, Pair{a: points[i], b: points[j], distance: points[i].Distance(points[j])})
		}
	}
	slices.SortFunc(pairs, func(a, b Pair) int {
		if a.distance < b.distance {
			return -1
		} else if a.distance > b.distance {
			return 1
		}
		return 0
	})
	return pairs
}

func getPoints(lines []string) []Point {
	points := []Point{}
	for _, line := range lines {
		splitted := strings.Split(line, ",")
		x, err := strconv.Atoi(splitted[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(splitted[1])
		if err != nil {
			panic(err)
		}
		z, err := strconv.Atoi(splitted[2])
		if err != nil {
			panic(err)
		}
		points = append(points, Point{x: x, y: y, z: z})
	}
	return points
}
