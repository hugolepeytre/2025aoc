// Package day3 contains solution for day 3
package day3

import (
	"fmt"
	"math"
	"slices"

	"aoc2025/parsing"
)

// Part1 solves first part of the puzzle.
func Part1(input string) {
	var count int64
	for _, n := range parsing.ParseDigits(input) {
		count += findMaxBatteryVal(n, 2)
	}

	fmt.Println(count)
}

// Part2 solves second part of the puzzle.
func Part2(input string) {
	var count int64
	for _, n := range parsing.ParseDigits(input) {
		count += findMaxBatteryVal(n, 12)
	}

	fmt.Println(count)
}

func findMaxBatteryVal(v []int64, n int) int64 {
	if n == 1 {
		return slices.Max(v)
	}

	var maxFirstDigit int64
	maxFirstDigitIdx := 0
	for i, num := range v[:len(v)+1-n] {
		if num > maxFirstDigit {
			maxFirstDigit = num
			maxFirstDigitIdx = i
		}
	}
	maxRemainingVal := findMaxBatteryVal(v[(maxFirstDigitIdx+1):], n-1)

	return maxFirstDigit*int64(math.Pow10(n-1)) + maxRemainingVal
}
