// Package parsing for various parsing functions mostly
package parsing

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ReadInput reads to input for the required day to a string.
func ReadInput(dayPick int) string {
	b, err := os.ReadFile(fmt.Sprintf("inputs/%v", dayPick))
	if err != nil {
		panic(fmt.Sprintf("Could not read input day %v\nError : %v", dayPick, err))
	}

	return string(b)
}

// ParseDigits returns all digits of each line of a string.
func ParseDigits(s string) [][]int64 {
	re := regexp.MustCompile(`\d`)

	res := make([][]int64, 0)
	for l := range strings.Lines(s) {
		matches := re.FindAllString(l, -1)
		subRes := make([]int64, len(matches))
		for i, m := range matches {
			n, err := strconv.ParseInt(m, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("Could not convert %v to number", m))
			}

			subRes[i] = n
		}
		res = append(res, subRes)
	}

	return res
}

// ParseNumbers returns all numbers of each line of a string.
func ParseNumbers(s string) [][]int64 {
	re := regexp.MustCompile(`-?\d+`)

	res := make([][]int64, 0)
	for l := range strings.Lines(s) {
		matches := re.FindAllString(l, -1)
		subRes := make([]int64, len(matches))
		for i, m := range matches {
			n, err := strconv.ParseInt(m, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("Could not convert %v to number", m))
			}

			subRes[i] = n
		}
		res = append(res, subRes)
	}

	return res
}

// ParseNumbersNoSplit returns all numbers of a string.
func ParseNumbersNoSplit(s string) []int64 {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(s, -1)

	res := make([]int64, len(matches))

	for i, m := range matches {
		n, err := strconv.ParseInt(m, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("Could not convert %v to number", m))
		}

		res[i] = n
	}

	return res
}
