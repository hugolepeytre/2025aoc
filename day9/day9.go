// Package day9 contains solution for day 9
package day9

import (
	"cmp"
	"fmt"
	"slices"

	"aoc2025/grid"
	"aoc2025/parsing"
)

// Part1 solves first part of the puzzle.
func Part1(input string) {
	coords := parsing.ParseNumbers(input)
	areas := make([]int64, 0)
	for i, c1 := range coords {
		for j, c2 := range coords {
			if i < j {
				diff := (abs(c1[0]-c2[0]) + 1) * (abs(c1[1]-c2[1]) + 1)
				areas = append(areas, diff)
			}
		}
	}

	slices.Sort(areas)
	count := areas[len(areas)-1]
	fmt.Printf("Part 1 : %v\n", count)
}

// Part2 solves second part of the puzzle.
func Part2(input string) {
	sampleInput := false
	lastStep := grid.Left
	if sampleInput {
		lastStep = grid.Up
	}
	lastOutsideDir := lastStep.CounterClockwise()

	coords := parsing.ParseNumbers(input)
	border := make(coordSet)
	perimeter := make(coordSet)

	c := grid.Coord{X: coords[0][0], Y: coords[0][1]}
	prevStep := lastStep.Value()
	outsideDir := lastOutsideDir

	for i := range coords {
		nextC := grid.CoordFrom(coords[(i+1)%len(coords)])
		step := grid.Coord{X: sign(nextC.X - c.X), Y: sign(nextC.Y - c.Y)}
		outsideDir = adjustOutsideDir(outsideDir, step, prevStep)
		n := c
		border[n.Add(outsideDir.Value())] = true
		for n != nextC {
			n = n.Add(step)
			border[n.Add(outsideDir.Value())] = true
			perimeter[n] = true
		}
		c = nextC
		prevStep = step
	}

	for c := range perimeter {
		delete(border, c)
	}

	rectangles := make([]rect, 0)
	for i, c1 := range coords {
		for j, c2 := range coords {
			if i < j {
				diff := (abs(c1[0]-c2[0]) + 1) * (abs(c1[1]-c2[1]) + 1)
				r := rect{c1: grid.CoordFrom(c1), c2: grid.CoordFrom(c2), area: diff}
				rectangles = append(rectangles, r)
			}
		}
	}

	slices.SortFunc(rectangles, func(a, b rect) int {
		return cmp.Compare(b.area, a.area)
	})

	var count int64
	i := 0
	for count == 0 {
		r := rectangles[i]
		if !crossesBorder(r.c1, r.c2, &border) {
			count = r.area
		}
		i++
	}
	fmt.Printf("Part 2 : %v\n", count)
}

func crossesBorder(c1 grid.Coord, c2 grid.Coord, border *coordSet) bool {
	maxX := max(c1.X, c2.X)
	maxY := max(c1.Y, c2.Y)
	minX := min(c1.X, c2.X)
	minY := min(c1.Y, c2.Y)
	for c := range *border {
		if minX <= c.X && c.X <= maxX && minY <= c.Y && c.Y <= maxY {
			return true
		}
	}

	return false
}

func adjustOutsideDir(d grid.Direction, step grid.Coord, prevStep grid.Coord) grid.Direction {
	// Optimized to the point of unreadability, but I wanted to make code shorter
	sameSign := step.Y*prevStep.X+step.X*prevStep.Y > 0
	if (step.Y != 0) != sameSign {
		return d.CounterClockwise()
	}

	return d.Clockwise()
}

func sign(a int64) int64 {
	if a > 0 {
		return 1
	}

	if a < 0 {
		return -1
	}

	return 0
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}

	return a
}

type rect struct {
	c1   grid.Coord
	c2   grid.Coord
	area int64
}

type coordSet = map[grid.Coord]bool
