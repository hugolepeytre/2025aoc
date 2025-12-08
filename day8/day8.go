// Package day8 contains solution for day 8
package day8

import (
	"cmp"
	"fmt"
	"math"
	"slices"

	"aoc2025/parsing"
)

// Part1 solves first part of the puzzle.
func Part1(input string) {
	const nFirstEdges = 1000
	edges, _ := getEdges(input)

	nextSetID := 0
	boxToSet := make(map[[3]int64]int)
	boxSets := make(mapOfSets)
	for i := range nFirstEdges {
		addEdge(&boxToSet, &boxSets, &nextSetID, edges[i])
	}

	lengths := make([]int, nextSetID)
	for id, s := range boxSets {
		lengths[id] = len(s)
	}
	slices.Sort(lengths)

	count := lengths[nextSetID-1] * lengths[nextSetID-2] * lengths[nextSetID-3]
	fmt.Printf("Part 1 : %v\n", count)
}

// Part2 solves second part of the puzzle.
func Part2(input string) {
	edges, nBoxes := getEdges(input)

	nextSetID := 0
	i := 0
	boxToSet := make(map[[3]int64]int)
	boxSets := make(mapOfSets)
	for !isOver(&boxSets, nBoxes) {
		addEdge(&boxToSet, &boxSets, &nextSetID, edges[i])
		i++
	}

	lastEdge := edges[i-1]
	count := lastEdge.b1[0] * lastEdge.b2[0]
	// fmt.Printf("%+v\n", edges)
	// fmt.Printf("%+v\n", lastEdge)
	// fmt.Printf("%+v\n", i)
	fmt.Printf("Part 2 : %v\n", count)
}

func getEdges(input string) ([]edge, int) {
	boxes := parsing.ParseNumbers(input)

	edges := make([]edge, 0)
	for i, b1 := range boxes {
		for _, b2 := range boxes[i+1:] {
			dx := math.Pow(float64(b1[0]-b2[0]), 2)
			dy := math.Pow(float64(b1[1]-b2[1]), 2)
			dz := math.Pow(float64(b1[2]-b2[2]), 2)
			dist := math.Sqrt(dx + dy + dz)
			box1 := [3]int64{b1[0], b1[1], b1[2]}
			box2 := [3]int64{b2[0], b2[1], b2[2]}
			edges = append(edges, edge{b1: box1, b2: box2, dist: dist})
		}
	}

	slices.SortFunc(edges, func(a, b edge) int {
		return cmp.Compare(a.dist, b.dist)
	})

	return edges, len(boxes)
}

func addEdge(boxToSet *map[[3]int64]int, boxSets *mapOfSets, nextSetID *int, e edge) {
	sid1, contains1 := (*boxToSet)[e.b1]
	sid2, contains2 := (*boxToSet)[e.b2]

	if contains1 && contains2 && sid1 != sid2 {
		for box := range (*boxSets)[sid2] {
			(*boxToSet)[box] = sid1
			(*boxSets)[sid1][box] = true
		}
		delete(*boxSets, sid2)
	}
	if contains1 && !contains2 {
		(*boxToSet)[e.b2] = sid1
		(*boxSets)[sid1][e.b2] = true
	}
	if !contains1 && contains2 {
		(*boxToSet)[e.b1] = sid2
		(*boxSets)[sid2][e.b1] = true
	}
	if !contains1 && !contains2 {
		(*boxToSet)[e.b1] = *nextSetID
		(*boxToSet)[e.b2] = *nextSetID
		(*boxSets)[*nextSetID] = make(setOfCoords)
		(*boxSets)[*nextSetID][e.b1] = true
		(*boxSets)[*nextSetID][e.b2] = true
		(*nextSetID)++
	}
}

// Check if the first set of the map contains all boxes.
func isOver(boxSets *mapOfSets, nBoxes int) bool {
	for _, v := range *boxSets {
		return len(v) == nBoxes
	}

	return false
}

type edge struct {
	b1   [3]int64
	b2   [3]int64
	dist float64
}

type (
	setOfCoords = map[[3]int64]bool
	mapOfSets   = map[int]setOfCoords
)
