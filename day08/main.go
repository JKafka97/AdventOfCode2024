package main

import (
	"strings"
)

// isInBounds checks if the coordinates (x, y) are within the grid of size n x n.
func isInBounds(x, y, n int) bool {
	return 0 <= x && x < n && 0 <= y && y < n
}

// getAntinodePositions calculates potential antinodes given two points a and b in an n x n grid.
func getAntinodePositions(a, b [2]int, n int) [][2]int {
	ax, ay := a[0], a[1]
	bx, by := b[0], b[1]

	// Calculate the positions of potential antinodes
	cx, cy := ax-(bx-ax), ay-(by-ay) // Direction: back from a
	dx, dy := bx+(bx-ax), by+(by-ay) // Direction: forward from b

	var antinodes [][2]int
	// Add valid antinodes within bounds
	if isInBounds(cx, cy, n) {
		antinodes = append(antinodes, [2]int{cx, cy})
	}
	if isInBounds(dx, dy, n) {
		antinodes = append(antinodes, [2]int{dx, dy})
	}

	return antinodes
}

// getExtendedAntinodePositions calculates extended antinodes by moving backward and forward from points a and b.
func getExtendedAntinodePositions(a, b [2]int, n int) [][2]int {
	ax, ay := a[0], a[1]
	bx, by := b[0], b[1]
	dx, dy := bx-ax, by-ay // Direction vector from a to b

	var antinodes [][2]int

	// Extend antinodes backward from a
	i := 0
	for {
		x, y := ax-dx*i, ay-dy*i
		if isInBounds(x, y, n) {
			antinodes = append(antinodes, [2]int{x, y})
		} else {
			break
		}
		i++
	}

	// Extend antinodes forward from b
	i = 0
	for {
		x, y := bx+dx*i, by+dy*i
		if isInBounds(x, y, n) {
			antinodes = append(antinodes, [2]int{x, y})
		} else {
			break
		}
		i++
	}

	return antinodes
}

// Processes the input string and returns the number of unique antinodes for part 1 of the problem.
func SolutionForPart1(input string) (int, error) {
	// Read the grid from input string
	grid := strings.Split(strings.TrimSpace(input), "\n")
	n := len(grid)

	// Collect all locations of interest (not empty cells)
	locsByChar := make(map[byte][][2]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '.' {
				locsByChar[grid[i][j]] = append(locsByChar[grid[i][j]], [2]int{i, j})
			}
		}
	}

	// Collect all unique antinodes
	antinodes := make(map[[2]int]struct{})
	for _, locs := range locsByChar {
		// Generate pairs of points from the same character
		for i := 0; i < len(locs); i++ {
			for j := i + 1; j < len(locs); j++ {
				a, b := locs[i], locs[j]
				// Add antinodes from the pair of points
				for _, antinode := range getAntinodePositions(a, b, n) {
					antinodes[antinode] = struct{}{}
				}
			}
		}
	}

	return len(antinodes), nil
}

// Processes the input string and returns the number of unique extended antinodes for part 2 of the problem.
func SolutionForPart2(input string) (int, error) {
	// Read the grid from input string
	grid := strings.Split(strings.TrimSpace(input), "\n")
	n := len(grid)

	// Collect all locations of interest (not empty cells)
	locsByChar := make(map[byte][][2]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '.' {
				locsByChar[grid[i][j]] = append(locsByChar[grid[i][j]], [2]int{i, j})
			}
		}
	}

	// Collect all unique extended antinodes
	antinodes := make(map[[2]int]struct{})
	for _, locs := range locsByChar {
		// Generate pairs of points from the same character
		for i := 0; i < len(locs); i++ {
			for j := i + 1; j < len(locs); j++ {
				a, b := locs[i], locs[j]
				// Add extended antinodes from the pair of points
				for _, antinode := range getExtendedAntinodePositions(a, b, n) {
					antinodes[antinode] = struct{}{}
				}
			}
		}
	}

	return len(antinodes), nil
}
