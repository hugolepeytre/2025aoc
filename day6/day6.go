// Package day6 contains solution for day 6
package day6

import (
	"fmt"
	"strings"

	"aoc2025/parsing"
)

// Part1 solves first part of the puzzle.
func Part1(input string) {
	splits := strings.Split(input, "\n")
	ops := strings.Fields(splits[len(splits)-2])
	lists := parsing.ParseNumbers(input)

	var count int64
	for i := range ops {
		switch ops[i] {
		case "*":
			count += lists[0][i] * lists[1][i] * lists[2][i] * lists[3][i]
		case "+":
			count += lists[0][i] + lists[1][i] + lists[2][i] + lists[3][i]
		}
	}
	fmt.Println(count)
}

// Part2 solves second part of the puzzle.
func Part2(input string) {
	splits := strings.Split(input, "\n")
	lines := make([][]rune, 0)
	for _, split := range splits {
		lines = append(lines, []rune(split))
	}
	var totalSum int64
	var subSum int64
	op := '+'
	for i := range lines[0] {
		var val int64
		if lines[4][i] != ' ' {
			op = lines[4][i]
			switch op {
			case '*':
				subSum = 1
			case '+':
				subSum = 0
			default:
				panic("Unknown op")
			}
		}
		for j := range 4 {
			r := lines[j][i]
			if r != ' ' {
				val = 10*val + int64(r-'0')
			}
		}
		if val == 0 {
			totalSum += subSum
			subSum = 0
			op = 'x'
		} else {
			switch op {
			case '*':
				subSum *= val
			case '+':
				subSum += val
			default:
				panic("Unknown op")
			}
		}
	}
	count := totalSum + subSum
	fmt.Println(count)
}
