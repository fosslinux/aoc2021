package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	s "strings"
)

type point struct {
	x int
	y int
}

func str2point(str string) point {
	sxy := s.Split(str, ",")
	x, _ := strconv.Atoi(sxy[0])
	y, _ := strconv.Atoi(sxy[1])
	return point{x: x, y: y}
}

type line struct {
	a point
	b point
}

func (l *line) points() []point {
	if l.a.x == l.b.x {
		pointCount := int(math.Abs(float64(l.a.y-l.b.y))) + 1
		lpoints := make([]point, pointCount)
		for i := 0; i < pointCount; i++ {
			smally := 0
			if l.a.y > l.b.y {
				smally = l.b.y
			} else {
				smally = l.a.y
			}
			lpoints[i] = point{x: l.a.x, y: smally + i}
		}
		return lpoints
	} else if l.a.y == l.b.y {
		pointCount := int(math.Abs(float64(l.a.x-l.b.x))) + 1
		lpoints := make([]point, pointCount)
		for i := 0; i < pointCount; i++ {
			smallx := 0
			if l.a.x > l.b.x {
				smallx = l.b.x
			} else {
				smallx = l.a.x
			}
			lpoints[i] = point{x: smallx + i, y: l.b.y}
		}
		return lpoints
	}
	return make([]point, 0)
}

func parta(in []string) int {
	// Create grid
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
		for j := range grid[i] {
			grid[i][j] = 0
		}
	}

	// Populate grid
	for _, sLine := range in {
		rawPoints := s.Split(sLine, " ")
		pointa := str2point(rawPoints[0])
		pointb := str2point(rawPoints[2])
		line := line{a: pointa, b: pointb}
		for _, point := range line.points() {
			grid[point.x][point.y]++
		}
	}

	// Calculate
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell >= 2 {
				count++
			}
		}
	}
	return count
}

func main() {
	bin, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Unable to read file: %s", err))
	}
	in := s.Split(string(bin), "\n")
	in = in[:len(in)-1]

	if os.Args[2] == "a" {
		fmt.Println(parta(in))
	}
}
