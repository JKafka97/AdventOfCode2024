package main

import (
	"strings"
)

// SolutionForPart1 calculates the sum of scores of all trailheads.
func SolutionForPart1(input string) (int, error) {
	topographicMap := parseInput(input)
	rows := len(topographicMap)
	cols := len(topographicMap[0])
	scoreSum := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if topographicMap[r][c] == 0 {
				scoreSum += bfsScore(topographicMap, r, c, rows, cols)
			}
		}
	}

	return scoreSum, nil
}

// SolutionForPart2 calculates the sum of ratings of all trailheads.
func SolutionForPart2(input string) (int, error) {
	topographicMap := parseInput(input)
	rows := len(topographicMap)
	cols := len(topographicMap[0])
	ratingSum := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if topographicMap[r][c] == 0 {
				ratingSum += countDistinctTrails(topographicMap, r, c, rows, cols)
			}
		}
	}

	return ratingSum, nil
}

func parseInput(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	topoMap := make([][]int, len(lines))
	for i, line := range lines {
		topoMap[i] = make([]int, len(line))
		for j, ch := range line {
			topoMap[i][j] = int(ch - '0')
		}
	}
	return topoMap
}

func bfsScore(grid [][]int, startRow, startCol, rows, cols int) int {
	reachableNines := map[[2]int]bool{}
	visited := map[[2]int]bool{}
	queue := [][2]int{{startRow, startCol}}
	visited[[2]int{startRow, startCol}] = true

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if grid[current[0]][current[1]] == 9 {
			reachableNines[current] = true
			continue
		}

		for _, d := range directions {
			nr, nc := current[0]+d[0], current[1]+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				if !visited[[2]int{nr, nc}] && grid[nr][nc] == grid[current[0]][current[1]]+1 {
					visited[[2]int{nr, nc}] = true
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
	}

	return len(reachableNines)
}

func countDistinctTrails(grid [][]int, startRow, startCol, rows, cols int) int {
	dp := map[[3]int]int{}
	return dfs(grid, startRow, startCol, rows, cols, dp)
}

func dfs(grid [][]int, r, c, rows, cols int, dp map[[3]int]int) int {
	if grid[r][c] == 9 {
		return 1
	}

	if val, exists := dp[[3]int{r, c, grid[r][c]}]; exists {
		return val
	}

	count := 0
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, d := range directions {
		nr, nc := r+d[0], c+d[1]
		if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == grid[r][c]+1 {
			count += dfs(grid, nr, nc, rows, cols, dp)
		}
	}
	dp[[3]int{r, c, grid[r][c]}] = count
	return count
}
