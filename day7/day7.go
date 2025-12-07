// Package day7 contains solution for day 7
package day7

import (
	"fmt"
	"strings"
)

// Part1 solves first part of the puzzle.
func Part1(input string) {
	arr := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, 0)
	var count int64
	for i, l := range arr {
		grid = append(grid, []rune(l))
		for j, r := range grid[i] {
			if r == '^' && grid[i-2][j] == 'S' {
				if 0 <= j-1 {
					grid[i][j-1] = 'S'
				}
				if j+1 < len(arr[0]) {
					grid[i][j+1] = 'S'
				}
				count++
			}
			if r == '.' && i > 1 && grid[i-2][j] == 'S' {
				grid[i][j] = 'S'
			}
		}
	}

	fmt.Printf("Part 1 : %v\n", count)
}

// Part2 solves second part of the puzzle.
func Part2(input string) {
	arr := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]int64, 0)
	for i, l := range arr {
		grid = append(grid, make([]int64, len(l)))
		for j, r := range []rune(l) {
			if r == 'S' {
				grid[i][j] = 1
			}
			if r == '^' {
				if 0 <= j-1 {
					grid[i][j-1] += grid[i-2][j]
				}
				if j+1 < len(arr[0]) {
					grid[i][j+1] += grid[i-2][j]
				}
			}
			if r == '.' && i > 1 {
				grid[i][j] += grid[i-2][j]
			}
		}
	}

	var count int64
	for _, n := range grid[len(grid)-2] {
		count += n
	}
	fmt.Printf("Part 2 : %v\n", count)
}
