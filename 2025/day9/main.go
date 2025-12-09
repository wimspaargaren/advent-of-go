package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wimspaargaren/aoc"
)

func main() {
	input := aoc.MustReadFile("input.txt")
	points := getPoints(strings.Split(strings.TrimSpace(input), "\n"))
	pairs, segments := getPairsAndSegments(points)
	part1 := 0
	part2 := 0
	for _, pair := range pairs {
		area := pair.Area()
		if area > part1 {
			part1 = area
		}
		if area <= part2 {
			continue
		}
		intersect := false
		minX, maxX := pair.minMaxX()
		minY, maxY := pair.minMaxY()

		for _, segment := range segments {
			if pointIntersect(minX, maxX, minY, maxY, segment) {
				intersect = true
				break
			}
		}
		if intersect {
			continue
		}

		if area > part2 {
			part2 = area
		}

	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func pointIntersect(minX, maxX, minY, maxY int, segment Segment) bool {
	segMinX, segMaxX := segment.minMaxX()
	segMinY, segMaxY := segment.minMaxY()
	if segment.start.x == segment.end.x {
		if segment.start.x <= minX || segment.start.x >= maxX {
			return false
		}
		overlapStart := max(minY, segMinY)
		overlapEnd := min(maxY, segMaxY)
		return overlapStart < overlapEnd
	}

	if segment.start.y <= minY || segment.start.y >= maxY {
		return false
	}
	overlapStart := max(minX, segMinX)
	overlapEnd := min(maxX, segMaxX)
	return overlapStart < overlapEnd
}

type Point struct {
	x, y int
}

type Pair struct {
	a, b Point
}

func (p Pair) Area() int {
	distX := abs(p.a.x-p.b.x) + 1
	distY := abs(p.a.y-p.b.y) + 1
	return distX * distY
}

func (p Pair) minMaxX() (int, int) {
	if p.a.x < p.b.x {
		return p.a.x, p.b.x
	}
	return p.b.x, p.a.x
}

func (p Pair) minMaxY() (int, int) {
	if p.a.y < p.b.y {
		return p.a.y, p.b.y
	}
	return p.b.y, p.a.y
}

type Segment struct {
	start, end Point
}

func (s Segment) minMaxX() (int, int) {
	if s.start.x < s.end.x {
		return s.start.x, s.end.x
	}
	return s.end.x, s.start.x
}

func (s Segment) minMaxY() (int, int) {
	if s.start.y < s.end.y {
		return s.start.y, s.end.y
	}
	return s.end.y, s.start.y
}

func getPairsAndSegments(points []Point) ([]Pair, []Segment) {
	var pairs []Pair
	var segments []Segment
	for i := range points {
		cur := points[i]
		next := points[(i+1)%len(points)]
		if cur.x == next.x || cur.y == next.y {
			segments = append(segments, Segment{cur, next})
		}

		for j := i + 1; j < len(points); j++ {
			pairs = append(pairs, Pair{points[i], points[j]})
		}
	}
	return pairs, segments
}

func getPoints(lines []string) []Point {
	points := make([]Point, len(lines))
	for i, line := range lines {
		splitted := strings.Split(line, ",")
		x, err := strconv.Atoi(splitted[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(splitted[1])
		if err != nil {
			panic(err)
		}
		points[i] = Point{x, y}
	}
	return points
}
