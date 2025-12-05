// Package day5 contains solution for day 5
package day5

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"aoc2025/parsing"
)

// Part1 solves first part of the puzzle.
func Part1(input string) {
	rangesStr, idsStr, _ := strings.Cut(strings.ReplaceAll(input, "-", ";"), "\n\n")
	ranges := make([]pair[int64, int64], 0)

	for _, r := range parsing.ParseNumbers(rangesStr) {
		ranges = append(ranges, pair[int64, int64]{one: r[0], two: r[1]})
	}

	ids := parsing.ParseNumbersNoSplit(idsStr)
	count := 0
	for _, id := range ids {
		for _, p := range ranges {
			low, high := p.one, p.two
			if low <= id && id <= high {
				count++

				break
			}
		}
	}

	fmt.Println(count)
}

// Part2 solves second part of the puzzle.
func Part2(input string) {
	rangesStr, _, _ := strings.Cut(strings.ReplaceAll(input, "-", ";"), "\n\n")
	ranges := make([]pair[int64, int64], 0)

	for _, r := range parsing.ParseNumbers(rangesStr) {
		ranges = append(ranges, pair[int64, int64]{one: r[0], two: r[1]})
	}

	rangeCmp := func(a, b pair[int64, int64]) int {
		if cmp.Compare(a.one, b.one) != 0 {
			return cmp.Compare(a.one, b.one)
		}

		return cmp.Compare(a.two, b.two)
	}
	slices.SortFunc(ranges, rangeCmp)

	var count int64
	currLow := ranges[0].one
	currHigh := ranges[0].two
	for _, r := range ranges {
		low, high := r.one, r.two
		if low <= currHigh {
			currHigh = max(currHigh, high)
		} else {
			count += currHigh - currLow + 1
			currHigh = high
			currLow = low
		}
	}
	count += currHigh - currLow + 1

	fmt.Println(count)
}

type pair[T, U any] struct {
	one T
	two U
}
