// Package day4 contains solution for day 4
package day4

import (
	"fmt"

	"aoc2025/grid"
	"aoc2025/parsing"
)

// Part1 solves first part of the puzzle.
func Part1(input string) {
	g := parsing.StringToGrid(input)

	initialCount := g.Count('@')
	removeRound(&g)
	finalCount := g.Count('@')

	count := initialCount - finalCount
	fmt.Println(count)
}

// Part2 solves second part of the puzzle.
func Part2(input string) {
	g := parsing.StringToGrid(input)

	initialCount := g.Count('@')
	for removeRound(&g) {
	}
	finalCount := g.Count('@')

	count := initialCount - finalCount
	fmt.Println(count)
}

func removeRound(g *grid.Grid[rune]) bool {
	adjCounts := make([]int, g.Height*g.Width)
	initialCount := g.Count('@')

	for i := range adjCounts {
		currPos := g.IdxToPos(int64(i))
		for _, d := range grid.IterDirs() {
			newPos, ok := g.Mvt(currPos, d.Value())
			if ok && g.Get(newPos) == '@' {
				adjCounts[g.PosToIdx(currPos)]++
			}
		}
	}

	for i, c := range adjCounts {
		if c < 4 && g.Get(g.IdxToPos(int64(i))) == '@' {
			g.Set(g.IdxToPos(int64(i)), '.')
		}
	}

	finalCount := g.Count('@')

	return initialCount != finalCount
}
