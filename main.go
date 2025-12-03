// Package main runs solution for any day given as command line argument
package main

import (
	"fmt"
	"os"
	"strconv"

	"aoc2025/day1"
	"aoc2025/day2"
	"aoc2025/day3"
	"aoc2025/parsing"
)

var part1 = map[int]func(string){
	1: day1.Part1,
	2: day2.Part1,
	3: day3.Part1,
}

var part2 = map[int]func(string){
	1: day1.Part2,
	2: day2.Part2,
	3: day3.Part2,
}

const argsNeeded = 1

func main() {
	if len(os.Args) < argsNeeded+1 {
		fmt.Println("Specify the day number to run as the first argument to the program")
		os.Exit(0)
	}

	dayPick, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Could not read day. Specify the day number to run as the first argument to the program")
		os.Exit(0)
	}

	input := parsing.ReadInput(dayPick)

	p1, ok := part1[dayPick]
	if !ok {
		panic(fmt.Sprintf("No function found for day %d part 1", dayPick))
	}

	p1(input)

	p2, ok := part2[dayPick]
	if !ok {
		panic(fmt.Sprintf("No function found for day %d part 2", dayPick))
	}

	p2(input)
}
