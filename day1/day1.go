// Package day1 contains solution for day 1
package day1

import (
	"fmt"
	"strings"

	"aoc2025/parsing"
)

const dialSize = 100

// Part1 solves first part of the puzzle.
func Part1(input string) {
	var count, state int64 = 0, 50
	for _, instr := range parsing.ParseNumbersNoSplit(strings.ReplaceAll(input, "L", "-")) {
		if state == 0 {
			count++
		}

		state = (((state + instr) % dialSize) + dialSize) % dialSize // ensuring positive remainder
	}

	if state == 0 {
		count++
	}

	fmt.Println(count)
}

// Part2 solves second part of the puzzle.
func Part2(input string) {
	var count, prevClickVal int64 = 0, 50
	for _, instr := range parsing.ParseNumbersNoSplit(strings.ReplaceAll(input, "L", "-")) {
		newVal := (((prevClickVal + instr) % dialSize) + dialSize) % dialSize // ensuring positive remainder
		fullRots := abs(instr) / dialSize

		zeroClicked := ((instr < 0 && newVal > prevClickVal) || (instr > 0 && newVal < prevClickVal) || newVal == 0)
		zeroClicked = zeroClicked && prevClickVal != 0

		if zeroClicked {
			count++
		}

		count += fullRots
		prevClickVal = newVal
	}

	fmt.Println(count)
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}

	return n
}
