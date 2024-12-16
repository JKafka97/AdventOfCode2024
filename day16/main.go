package main

import (
	"bufio"
	"strings"
)

type Point struct {
	x, y int
}

type State struct {
	loc Point
	dir Point
}

const (
	inf = 1 << 30 // A large number to represent infinity
)

var directions = []Point{
	{-1, 0}, {0, 1}, {1, 0}, {0, -1},
}

func add(p1, p2 Point) Point {
	return Point{p1.x + p2.x, p1.y + p2.y}
}

func nextStates(state State) map[State]int {
	loc := state.loc
	d := state.dir
	// Normal move (cost 1), turn left or right (cost 1000)
	return map[State]int{
		{add(loc, d), d}:        1,    // Move forward (cost 1)
		{loc, Point{-d.y, d.x}}: 1000, // Turn left (cost 1000)
		{loc, Point{d.y, -d.x}}: 1000, // Turn right (cost 1000)
	}
}

func prevStates(state State) map[State]int {
	loc := state.loc
	d := state.dir
	// Reversing the direction (cost 1), turn left or right (cost 1000)
	return map[State]int{
		{add(loc, Point{-d.x, -d.y}), d}: 1,    // Move backward (cost 1)
		{loc, Point{-d.y, d.x}}:          1000, // Turn left (cost 1000)
		{loc, Point{d.y, -d.x}}:          1000, // Turn right (cost 1000)
	}
}

func solve(fileName string) (int, int) {
	scanner := bufio.NewScanner(strings.NewReader(fileName))
	grid := make(map[Point]rune)
	var start Point
	var end Point
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, c := range line {
			grid[Point{x, y}] = c
			if c == 'S' {
				start = Point{x, y}
			} else if c == 'E' {
				end = Point{x, y}
			}
		}
		y++
	}

	initState := State{start, directions[1]} // Start facing East (index 1)
	states := make(map[State]int)
	states[initState] = 0

	toUpdate := map[State]struct{}{initState: {}}

	// Part 1: Dijkstra-like search to find the minimal cost
	for len(toUpdate) > 0 {
		var current State
		for s := range toUpdate {
			current = s
			delete(toUpdate, s)
			break
		}

		cost := states[current]
		newStates := nextStates(current)

		for newState, costIncrease := range newStates {
			newLoc := newState.loc
			if grid[newLoc] == '#' {
				continue
			}
			newCost := cost + costIncrease
			if existingCost, exists := states[newState]; !exists || newCost < existingCost {
				states[newState] = newCost
				toUpdate[newState] = struct{}{}
			}
		}
	}

	// Find minimum cost for reaching the end state
	endState := State{}
	minCost := inf
	for _, d := range directions {
		stateCandidate := State{end, d}
		if cost, exists := states[stateCandidate]; exists && cost < minCost {
			minCost = cost
			endState = stateCandidate
		}
	}
	p1 := minCost

	// Part 2: Trace back the path to count how many locations were on the shortest path
	locsOnPath := map[Point]struct{}{end: {}}
	toCheck := map[State]struct{}{endState: {}}

	for len(toCheck) > 0 {
		var current State
		for s := range toCheck {
			current = s
			delete(toCheck, s)
			break
		}

		cost := states[current]

		for prevState, costIncrease := range prevStates(current) {
			prevLoc := prevState.loc

			if grid[prevLoc] == '#' {
				continue
			}

			if prevCost, exists := states[prevState]; exists && prevCost+costIncrease == cost {
				toCheck[prevState] = struct{}{}
				locsOnPath[prevLoc] = struct{}{}
			}
		}
	}

	p2 := len(locsOnPath)
	return p1, p2
}

func SolutionForPart1(input string) (int, error) {
	solution, _ := solve(input)
	return solution, nil
}

func SolutionForPart2(input string) (int, error) {
	_, solution := solve(input)
	return solution, nil
}
