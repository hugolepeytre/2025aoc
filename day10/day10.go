// Package day10 contains solution for day 10
package day10

import (
	"fmt"
	"strings"

	"aoc2025/parsing"
)

// Part1 solves first part of the puzzle.
func Part1(input string) {
	count := 0
	for l := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		count += solveLightProblem(lineToProblem(l))
	}
	fmt.Printf("Part 1 : %v\n", count)
}

// Part2 solves second part of the puzzle.
func Part2(input string) {
	var count int64
	for l := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		p := lineToProblem(l)
		m := makeParityMatrix(p.buttons)
		cache := make(joltageCache)

		var wrongSolVal int64
		for _, t := range p.joltageTarget {
			wrongSolVal += t + 1
		}

		count += solveJoltageProblem(p, &cache, &m, wrongSolVal)
	}
	fmt.Printf("Part 2 : %v\n", count)
}

func solveLightProblem(p problem) int {
	minPresses := len(p.buttons)
	for _, sol := range powerSet(p.buttons) {
		if (len(sol) < minPresses) && (buttonParity(&sol) == p.lightTarget) {
			minPresses = len(sol)
		}
	}

	return minPresses
}

// General solution idea inspired by
// https://www.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory
// Caching divides runtime by about 10.
func solveJoltageProblem(p problem, cache *joltageCache, m *parityMatrix, wrongSolVal int64) int64 {
	var sum int64
	for _, t := range p.joltageTarget {
		sum += t
	}
	if sum == 0 {
		return 0
	}
	if v, ok := (*cache)[sliceToString(&p.joltageTarget)]; ok {
		return v
	}

	minSol := wrongSolVal
	for _, sol := range (*m)[p.joltageTargetParity] {
		if newP, ok := p.pushAndHalve(sol); ok {
			v := 2*solveJoltageProblem(newP, cache, m, wrongSolVal) + int64(len(sol))
			minSol = min(minSol, v)
		}
	}

	(*cache)[sliceToString(&p.joltageTarget)] = minSol

	return minSol
}

func lineToProblem(l string) problem {
	elems := strings.Fields(l)
	diagram := elems[0]
	size := len(diagram) - 2

	var lightTarget int64
	for _, c := range diagram {
		if c == '.' {
			lightTarget <<= 1
		}
		if c == '#' {
			lightTarget = (lightTarget << 1) + 1
		}
	}

	buttons := make([]int64, len(elems)-2)
	for i, b := range elems[1 : len(elems)-1] {
		var button int64
		for _, n := range parsing.ParseNumbersNoSplit(b) {
			button |= 1 << (size - 1 - int(n))
		}
		buttons[i] = button
	}

	joltageTarget := make([]int64, 0)
	joltageTarget = append(joltageTarget, parsing.ParseNumbersNoSplit(elems[len(elems)-1])...)

	return problem{
		lightTarget:         lightTarget,
		joltageTargetParity: joltageParity(&joltageTarget),
		joltageTarget:       joltageTarget,
		buttons:             buttons,
		size:                size,
	}
}

func joltageParity(joltages *[]int64) int64 {
	var r int64
	for _, j := range *joltages {
		r <<= 1
		if j&1 == 1 {
			r++
		}
	}

	return r
}

func powerSet(buttons []int64) [][]int64 {
	powerset := [][]int64{make([]int64, 0)}

	for _, button := range buttons {
		for _, elem := range powerset {
			newElem := make([]int64, len(elem)+1)
			copy(newElem, elem)
			newElem[len(elem)] = button
			powerset = append(powerset, newElem)
		}
	}

	return powerset
}

func sliceToString(t *[]int64) string {
	return fmt.Sprint(t)
}

func buttonParity(buttons *[]int64) int64 {
	var parity int64
	for _, b := range *buttons {
		parity ^= b
	}

	return parity
}

type problem struct {
	lightTarget         int64
	joltageTarget       []int64
	joltageTargetParity int64
	buttons             []int64
	size                int
}

func makeParityMatrix(buttons []int64) parityMatrix {
	m := make(parityMatrix)
	for _, sol := range powerSet(buttons) {
		if v, ok := m[buttonParity(&sol)]; ok {
			v[sliceToString(&sol)] = sol
		} else {
			newV := make(map[string][]int64)
			newV[sliceToString(&sol)] = sol
			m[buttonParity(&sol)] = newV
		}
	}

	return m
}

func (p *problem) pushAndHalve(pushes []int64) (problem, bool) {
	newTargets := make([]int64, len(p.joltageTarget))
	copy(newTargets, p.joltageTarget)
	for _, button := range pushes {
		i := p.size
		for button > 0 {
			i--
			if button&1 == 1 {
				if newTargets[i] == 0 {
					return *p, false
				}
				newTargets[i]--
			}
			button >>= 1
		}
	}

	for i := range newTargets {
		newTargets[i] /= 2
	}

	return problem{
		joltageTarget:       newTargets,
		joltageTargetParity: joltageParity(&newTargets),
		lightTarget:         p.lightTarget,
		size:                p.size,
		buttons:             p.buttons,
	}, true
}

type (
	joltageCache = map[string]int64
	parityMatrix = map[int64]map[string][]int64
)
