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
	_ = input
	lines := strings.Split(strings.TrimSpace(input), "\n")

	points := getPoints(lines)
	pairs := getPairs(points)

	groups := []Group{}

	for i := 0; i < 1000; i++ {
		pair := pairs[i]
		indexA := -1
		indexB := -1
		for i, group := range groups {
			_, aExists := group.points[pair.a]
			if aExists {
				if indexA != -1 {
					panic("WAT")
				}
				indexA = i
			}
			_, bExists := group.points[pair.b]
			if bExists {
				if indexB != -1 {
					panic("WAT")
				}
				indexB = i
			}
		}
		if indexA == -1 && indexB == -1 {
			newGroup := Group{points: map[Point]bool{}}
			newGroup.AddPoint(pair.a)
			newGroup.AddPoint(pair.b)
			groups = append(groups, newGroup)
			continue
		}
		if indexA != -1 && indexB == -1 {
			groups[indexA].AddPoint(pair.a)
			groups[indexA].AddPoint(pair.b)
			continue
		}
		if indexA == -1 && indexB != -1 {
			groups[indexB].AddPoint(pair.a)
			groups[indexB].AddPoint(pair.b)
			continue
		}
		if indexA != indexB {
			for point := range groups[indexB].points {
				groups[indexA].AddPoint(point)
			}
			groups = append(groups[:indexB], groups[indexB+1:]...)
		}
	}

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
	fmt.Println("Result:", res)
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
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
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
