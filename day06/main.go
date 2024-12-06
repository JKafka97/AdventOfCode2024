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

var directionVectors = map[rune][2]int{
	Up:    {0, -1},
	Right: {1, 0},
	Down:  {0, 1},
	Left:  {-1, 0},
}

var turnRight = map[rune]rune{
	Up:    Right,
	Right: Down,
	Down:  Left,
	Left:  Up,
}

func SolutionForPart1(input string) (int, error) {
	if input == "" {
		return 0, errors.New("input cannot be empty")
	}

	rows := strings.Split(input, "\n")
	height := len(rows)
	if height == 0 || len(rows[0]) == 0 {
		return 0, errors.New("invalid grid dimensions")
	}

	width := len(rows[0])
	var guardX, guardY int
	var direction rune
	obstacles := make(map[[2]int]bool)

	for y, row := range rows {
		for x, char := range row {
			switch char {
			case Up, Right, Down, Left:
				guardX, guardY = x, y
				direction = char
			case '#':
				obstacles[[2]int{x, y}] = true
			}
		}
	}

	visited := make(map[[2]int]bool)
	visited[[2]int{guardX, guardY}] = true

	for {
		dx, dy := directionVectors[direction][0], directionVectors[direction][1]
		nextX, nextY := guardX+dx, guardY+dy

		if nextX < 0 || nextX >= width || nextY < 0 || nextY >= height {
			break
		}

		if obstacles[[2]int{nextX, nextY}] {
			direction = turnRight[direction] // Turn right on obstacle
		} else {
			guardX, guardY = nextX, nextY // Move forward
			visited[[2]int{guardX, guardY}] = true
		}
	}

	return len(visited), nil
}

func SolutionForPart2(input string) (int, error) {
	if input == "" {
		return 0, errors.New("input cannot be empty")
	}

	lines := strings.Split(input, "\n")
	n := len(lines)
	if n == 0 || len(lines[0]) == 0 {
		return 0, errors.New("invalid grid dimensions")
	}

	m := len(lines[0])
	grid := make([][]rune, n)
	for i := range lines {
		grid[i] = []rune(lines[i])
	}

	startRow, startCol, found := -1, -1, false
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == Up {
				startRow, startCol = i, j
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		return 0, errors.New("no guard starting position found")
	}

	directions := [][2]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}

	dirIndex := 0
	observedPositions := make(map[[2]int]bool)
	i, j := startRow, startCol

	for {
		observedPositions[[2]int{i, j}] = true
		nextI := i + directions[dirIndex][0]
		nextJ := j + directions[dirIndex][1]

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

		grid[oi][oj] = '#'
		i, j := startRow, startCol
		dirIndex := 0
		seenStates := make(map[[3]int]bool)

		for {
			if seenStates[[3]int{i, j, dirIndex}] {
				grid[oi][oj] = '.' // Restore original state
				return true
			}
			seenStates[[3]int{i, j, dirIndex}] = true

			nextI := i + directions[dirIndex][0]
			nextJ := j + directions[dirIndex][1]

			if nextI < 0 || nextI >= n || nextJ < 0 || nextJ >= m {
				grid[oi][oj] = '.' // Restore original state
				return false
			}

			if grid[nextI][nextJ] == '#' {
				dirIndex = (dirIndex + 1) % 4
			} else {
				i, j = nextI, nextJ
			}
		}
	}

	loopCount := 0
	for pos := range observedPositions {
		if isLoop(pos[0], pos[1]) {
			loopCount++
		}
	}

	return loopCount - 1, nil
}
