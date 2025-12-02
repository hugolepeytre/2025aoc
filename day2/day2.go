// Package day2 contains solution for day 2
package day2

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"aoc2025/parsing"
)

// Part1 solves first part of the puzzle.
func Part1(input string) {
	const nConcats = 2
	ranges, maxVal := rangesFromInput(input)

	var idx, count int64 = 1, 0
	testedVal := selfConcatNumber(idx, nConcats)
	for testedVal < maxVal {
		testedVal = selfConcatNumber(idx, nConcats)
		for _, p := range ranges {
			low, high := p.one, p.two
			if low <= testedVal && testedVal <= high {
				count += testedVal
			}
		}

		idx++
	}

	fmt.Println(count)
}

// Part2 solves second part of the puzzle.
func Part2(input string) {
	ranges, maxVal := rangesFromInput(input)

	var idx int64 = 1
	invalidIDs := make(map[int64]bool)
	for selfConcatNumber(idx, 2) < maxVal {
		var nConcat uint32 = 2
		for selfConcatNumber(idx, nConcat) < maxVal {
			testedID := selfConcatNumber(idx, nConcat)
			for _, p := range ranges {
				low, high := p.one, p.two
				if low <= testedID && testedID <= high {
					invalidIDs[testedID] = true
				}
			}
			nConcat++
		}

		idx++
	}

	var count int64
	for id := range invalidIDs {
		count += id
	}

	fmt.Println(count)
}

func rangesFromInput(input string) ([]pair[int64, int64], int64) {
	parsedIntegers := parsing.ParseNumbersNoSplit(strings.ReplaceAll(input, "-", "x"))

	ranges := make([]pair[int64, int64], len(parsedIntegers)/2)
	for i := 0; i < len(parsedIntegers)-1; i += 2 {
		ranges[i/2] = pair[int64, int64]{one: parsedIntegers[i], two: parsedIntegers[i+1]}
	}

	maxVal := slices.Max(parsedIntegers)

	return ranges, maxVal
}

func selfConcatNumber(num int64, n uint32) int64 {
	nDigits := int(math.Log10(float64(num))) + 1

	var res int64
	for i := range n {
		res += num * int64(math.Pow10(int(i)*nDigits))
	}

	return res
}

type pair[T, U any] struct {
	one T
	two U
}
