package main

import (
	"bufio"
	"container/heap"
	"strings"
)

type Point struct{ x, y int }
type State struct{ loc, dir Point }
type Point2 struct {
	row, col int
	dir      rune
}
type Item struct {
	distance int
	point    Point2
}
type PriorityQueue []*Item

var directions = []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // N, E, S, W

func add(p1, p2 Point) Point { return Point{p1.x + p2.x, p1.y + p2.y} }

func nextStates(state State) map[State]int {
	loc, d := state.loc, state.dir
	return map[State]int{
		{add(loc, d), d}:        1,
		{loc, Point{-d.y, d.x}}: 1000,
		{loc, Point{d.y, -d.x}}: 1000,
	}
}

func prevStates(state State) map[State]int {
	loc, d := state.loc, state.dir
	return map[State]int{
		{add(loc, Point{-d.x, -d.y}), d}: 1,
		{loc, Point{-d.y, d.x}}:          1000,
		{loc, Point{d.y, -d.x}}:          1000,
	}
}

func solve(fileName string) int {
	scanner := bufio.NewScanner(strings.NewReader(fileName))
	grid := make(map[Point]rune)
	var start, end Point
	y := 0

	for scanner.Scan() {
		for x, c := range scanner.Text() {
			grid[Point{x, y}] = c
			if c == 'S' {
				start = Point{x, y}
			}
			if c == 'E' {
				end = Point{x, y}
			}
		}
		y++
	}

	initState := State{start, directions[1]}
	states := map[State]int{initState: 0}
	toUpdate := map[State]struct{}{initState: {}}

	for len(toUpdate) > 0 {
		var current State
		for s := range toUpdate {
			current = s
			delete(toUpdate, s)
			break
		}

		cost := states[current]
		for newState, costIncrease := range nextStates(current) {
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

	locsOnPath := map[Point]struct{}{end: {}}
	toCheck := map[State]struct{}{{end, directions[1]}: {}}

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
	return len(locsOnPath)
}

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].distance < pq[j].distance }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func parse(lines []string) ([][]rune, Point2, Point2) {
	grid := make([][]rune, len(lines))
	var start, end Point2

	for r, line := range lines {
		grid[r] = []rune(strings.TrimSpace(line))
		for c, ch := range grid[r] {
			if ch == 'S' {
				start = Point2{r, c, 'E'}
			}
			if ch == 'E' {
				end = Point2{r, c, 'E'}
			}
		}
	}
	return grid, start, end
}

func dijkstra(grid [][]rune, starts []Point2) map[Point2]int {
	delta := map[rune][2]int{'E': {0, 1}, 'W': {0, -1}, 'N': {-1, 0}, 'S': {1, 0}}
	dist := make(map[Point2]int)
	pq := &PriorityQueue{}

	for _, start := range starts {
		dist[start] = 0
		heap.Push(pq, &Item{distance: 0, point: start})
	}

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		current := item.point

		if dist[current] < item.distance {
			continue
		}

		for _, nextDir := range "EWNS" {
			if nextDir != current.dir {
				nextPoint := Point2{current.row, current.col, nextDir}
				if _, exists := dist[nextPoint]; !exists || dist[nextPoint] > item.distance+1000 {
					dist[nextPoint] = item.distance + 1000
					heap.Push(pq, &Item{distance: dist[nextPoint], point: nextPoint})
				}
			}
		}

		dr, dc := delta[current.dir][0], delta[current.dir][1]
		nextRow, nextCol := current.row+dr, current.col+dc

		if nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0]) && grid[nextRow][nextCol] != '#' {
			nextPoint := Point2{nextRow, nextCol, current.dir}
			if _, exists := dist[nextPoint]; !exists || dist[nextPoint] > item.distance+1 {
				dist[nextPoint] = item.distance + 1
				heap.Push(pq, &Item{distance: dist[nextPoint], point: nextPoint})
			}
		}
	}

	return dist
}

func part1(input [][]rune, start Point2, end Point2) int {
	distances := dijkstra(input, []Point2{start})
	bestDistance := int(1e9)

	for _, dir := range "EWNS" {
		if distance, exists := distances[Point2{end.row, end.col, rune(dir)}]; exists {
			bestDistance = min(bestDistance, distance)
		}
	}
	return bestDistance
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func SolutionForPart1(input string) (int, error) {
	lines := strings.Split(input, "\n")
	grid, start, end := parse(lines)
	return part1(grid, start, end), nil
}

func SolutionForPart2(input string) (int, error) {
	return solve(input), nil
}
