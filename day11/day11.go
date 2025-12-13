// Package day11 contains solution for day 11
package day11

import (
	"fmt"
	"strings"

	"aoc2025/graph"
)

// Part1 solves first part of the puzzle.
func Part1(input string) {
	g, nameTable := inputToGraph(input)
	cache := make(map[int64]int)
	count := countPaths(&g, nameTable["you"], nameTable["out"], &cache)
	fmt.Printf("Part 1 : %v\n", count)
}

// Part2 solves second part of the puzzle.
func Part2(input string) {
	g, nameTable := inputToGraph(input)
	paths := [][]string{{"svr", "dac", "fft", "out"}, {"svr", "fft", "dac", "out"}}
	count := 0
	for _, p := range paths {
		prod := 1
		for i := range len(p) - 1 {
			source := p[i]
			sink := p[i+1]
			cache := make(map[int64]int)
			prod *= countPaths(&g, nameTable[source], nameTable[sink], &cache)
		}
		count += prod
	}
	fmt.Printf("Part 2 : %v\n", count)
}

func countPaths(g *graph.Graph[int64], start, end int64, cache *map[int64]int) int {
	if v, ok := (*cache)[start]; ok {
		return v
	} else if start == end {
		return 1
	} else if l, ok := g.EdgeList[start]; ok {
		sol := 0
		for _, e := range l {
			sol += countPaths(g, e.Next, end, cache)
		}
		(*cache)[start] = sol

		return sol
	}

	return 0
}

func inputToGraph(input string) (graph.Graph[int64], map[string]int64) {
	noColonInput := strings.TrimSpace(strings.ReplaceAll(input, ":", ""))
	nameTable := make(map[string]int64)
	var idCount int64
	edgeList := make(map[int64][]graph.Edge[int64])

	for l := range strings.SplitSeq(noColonInput, "\n") {
		var k int64
		edges := make([]graph.Edge[int64], 0)
		for i, n := range strings.Fields(l) {
			name := nameToInt64(n, &nameTable, &idCount)
			if i == 0 {
				k = name
			} else {
				edges = append(edges, graph.Edge[int64]{Next: name, Cost: 1})
			}
		}
		edgeList[k] = edges
	}

	g := graph.Graph[int64]{EdgeList: edgeList}

	return g, nameTable
}

func nameToInt64(n string, table *map[string]int64, idCount *int64) int64 {
	if v, ok := (*table)[n]; ok {
		return v
	}
	(*table)[n] = *idCount
	(*idCount)++

	return *idCount - 1
}
