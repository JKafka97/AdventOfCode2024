package main

import (
	"bufio"
	"strings"
)

func SolutionForPart1(input string) (int, error) {
	grid, sIdx, eIdx := parseInput(input)

	result := partOne(grid, sIdx, eIdx)
	return result, nil
}

func SolutionForPart2(input string) (int, error) {
	grid, sIdx, eIdx := parseInput(input)

	result := partTwo(grid, sIdx, eIdx)
	return result, nil
}

const (
	start = 'S'
	end   = 'E'
	track = '.'
	wall  = '#'
)

type index struct {
	r, c int
}

type direction struct {
	dr, dc int
}

var directions = []direction{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func parseInput(input string) (grid [][]rune, sIdx, eIdx index) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		grid = append(grid, make([]rune, 0, len(scanner.Text())))
		r := len(grid) - 1
		for c, block := range scanner.Text() {
			if block == start {
				sIdx = index{r: r, c: c}
			}
			if block == end {
				eIdx = index{r: r, c: c}
			}
			grid[r] = append(grid[r], block)
		}
	}
	return
}

func partOne(grid [][]rune, sIdx, eIdx index) (count int) {
	distanceFromStart := BFS(grid, sIdx)
	distanceFromEnd := BFS(grid, eIdx)

	minDistanceWithoutCheating := distanceFromStart[eIdx]

	for idx, ds := range distanceFromStart {
		for _, cIdx := range getCheatIndices(idx, len(grid), len(grid[0])) {
			if de, ok := distanceFromEnd[cIdx]; ok {
				dist := ds + de + getManhattenDistance(idx, cIdx)
				if minDistanceWithoutCheating-dist >= 100 {
					count++
				}
			}
		}
	}
	return
}

func partTwo(grid [][]rune, sIdx, eIdx index) int {
	N := len(grid)
	path := []index{sIdx}
	for path[len(path)-1] != eIdx {
		i, j := path[len(path)-1].r, path[len(path)-1].c
		for _, d := range directions {
			ii, jj := i+d.dr, j+d.dc
			if !inGrid(ii, jj, N) || grid[ii][jj] == wall || (len(path) > 1 && ii == path[len(path)-2].r && jj == path[len(path)-2].c) {
				continue
			}
			path = append(path, index{ii, jj})
			break
		}
	}

	og := len(path) - 1
	times := make(map[index]int)
	for t, coord := range path {
		times[coord] = og - t
	}

	maxLen := 20
	counts := make(map[int]int)
	saved := make(map[[4]int]int)
	for t, coord := range path {
		i, j := coord.r, coord.c
		for ii := i - maxLen; ii <= i+maxLen; ii++ {
			for jj := j - maxLen; jj <= j+maxLen; jj++ {
				timeUsed := abs(ii-i) + abs(jj-j)
				if !inGrid(ii, jj, N) || timeUsed > maxLen || grid[ii][jj] == wall {
					continue
				}
				remT := times[index{ii, jj}]
				saved[[4]int{i, j, ii, jj}] = og - (t + remT + timeUsed)
			}
		}
	}

	ans := 0
	for _, v := range saved {
		if v >= 0 {
			counts[v]++
		}
		if v >= 100 {
			ans++
		}
	}

	return ans
}

func inGrid(i, j, N int) bool {
	return 0 <= i && i < N && 0 <= j && j < N
}

func getCheatIndices(idx index, maxRow, maxCol int) (cheatIndices []index) {
	for i := 1; i <= 2; i++ {
		for _, dir := range directions {
			nIdx := index{r: idx.r + i*dir.dr, c: idx.c + i*dir.dc}
			if isWithinBounds(nIdx, maxRow, maxCol) {
				cheatIndices = append(cheatIndices, nIdx)
			}
		}
	}
	return
}

func getManhattenDistance(i1 index, i2 index) int {
	return abs(i1.r-i2.r) + abs(i1.c-i2.c)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func BFS(grid [][]rune, sIdx index) map[index]int {
	type qNode struct {
		idx      index
		distance int
	}
	maxRow, maxCol := len(grid), len(grid[0])

	distMap := make(map[index]int)
	queue := []qNode{{idx: sIdx, distance: 0}}
	visited := make(map[index]bool)
	visited[sIdx] = true

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		distMap[front.idx] = front.distance

		for _, dir := range directions {
			nIdx := index{r: front.idx.r + dir.dr, c: front.idx.c + dir.dc}
			if !visited[nIdx] && isWithinBounds(nIdx, maxRow, maxCol) && grid[nIdx.r][nIdx.c] != wall {
				queue = append(queue, qNode{idx: nIdx, distance: front.distance + 1})
				visited[nIdx] = true
			}
		}
	}
	return distMap
}

func isWithinBounds(idx index, maxRow, maxCol int) bool {
	return idx.r >= 0 && idx.c >= 0 && idx.r < maxRow && idx.c < maxCol
}
