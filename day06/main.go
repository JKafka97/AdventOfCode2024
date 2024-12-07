package main

import (
	"errors"
	"strings"
)

const (
	Up    = '^'
	Right = '>'
	Down  = 'v'
	Left  = '<'
)

var directions = [][2]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

func SolutionForPart1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	n, m := len(lines), len(lines[0])

	var startRow, startCol int
	found := false

	for i := 0; i < n && !found; i++ {
		for j := 0; j < m; j++ {
			if lines[i][j] == '^' {
				startRow, startCol = i, j
				found = true
				break
			}
		}
	}

	i, j := startRow, startCol
	dirIndex := 0
	visited := make(map[[2]int]bool)

	for {
		visited[[2]int{i, j}] = true

		nextI := i + directions[dirIndex][0]
		nextJ := j + directions[dirIndex][1]

		if nextI < 0 || nextI >= n || nextJ < 0 || nextJ >= m {
			break
		}

		if lines[nextI][nextJ] == '#' {
			dirIndex = (dirIndex + 1) % 4
		} else {
			i, j = nextI, nextJ
		}
	}

	return len(visited), nil
}

func SolutionForPart2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	n, m := len(lines), len(lines[0])

	grid := make([][]rune, n)
	for i := range lines {
		grid[i] = []rune(lines[i])
	}

	var startRow, startCol int
	found := false

	for i := 0; i < n && !found; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == Up {
				startRow, startCol = i, j
				found = true
				break
			}
		}
	}

	if !found {
		return 0, errors.New("no guard starting position found")
	}

	dirIndex := 0
	i, j := startRow, startCol
	originalVisited := make(map[[2]int]bool)

	for {
		originalVisited[[2]int{i, j}] = true

		nextI, nextJ := i+directions[dirIndex][0], j+directions[dirIndex][1]

		if nextI < 0 || nextI >= n || nextJ < 0 || nextJ >= m {
			break
		}

		if grid[nextI][nextJ] == '#' {
			dirIndex = (dirIndex + 1) % 4
		} else {
			i, j = nextI, nextJ
		}
	}

	isLoop := func(oi, oj int) bool {
		if grid[oi][oj] == '#' {
			return false
		}

		original := grid[oi][oj]
		grid[oi][oj] = '#' // Temporarily place an obstacle

		i, j, dirIndex := startRow, startCol, 0
		seenStates := make(map[[3]int]bool)

		for {
			state := [3]int{i, j, dirIndex}
			if seenStates[state] {
				grid[oi][oj] = original // Restore the grid
				return true
			}
			seenStates[state] = true

			nextI, nextJ := i+directions[dirIndex][0], j+directions[dirIndex][1]

			if nextI < 0 || nextI >= n || nextJ < 0 || nextJ >= m {
				break
			}

			if grid[nextI][nextJ] == '#' {
				dirIndex = (dirIndex + 1) % 4
			} else {
				i, j = nextI, nextJ
			}
		}

		grid[oi][oj] = original // Restore the grid
		return false
	}

	loopCount := 0
	for pos := range originalVisited {
		oi, oj := pos[0], pos[1]
		if oi == startRow && oj == startCol {
			continue // Cannot place obstacle where the guard starts
		}
		if isLoop(oi, oj) {
			loopCount++
		}
	}

	return loopCount, nil
}
