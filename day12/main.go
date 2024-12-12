package main

import (
	"bufio"
	"errors"
	"strings"
)

type GardenMap struct {
	grid                map[complex128]string
	connectedComponents []map[complex128]struct{}
}

func NewGardenMapFromInput(input string) (*GardenMap, error) {
	grid := make(map[complex128]string)
	scanner := bufio.NewScanner(strings.NewReader(input))
	r := 0
	for scanner.Scan() {
		line := scanner.Text()
		for c, chr := range line {
			grid[complex(float64(c), float64(r))] = string(chr)
		}
		r++
	}

	return &GardenMap{
		grid: grid,
	}, nil
}

func (g *GardenMap) getNeighbors(z complex128) []complex128 {
	directions := []complex128{1, -1, 1i, -1i}
	var nbrs []complex128
	for _, d := range directions {
		if g.grid[z+d] == g.grid[z] {
			nbrs = append(nbrs, z+d)
		}
	}
	return nbrs
}

func (g *GardenMap) getConnectedComponents() {
	if len(g.connectedComponents) > 0 {
		return
	}

	visited := make(map[complex128]struct{})
	for z := range g.grid {
		if _, ok := visited[z]; ok {
			continue
		}
		component := make(map[complex128]struct{})
		stack := []complex128{z}
		for len(stack) > 0 {
			s := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if _, ok := component[s]; !ok {
				component[s] = struct{}{}
				stack = append(stack, g.getNeighbors(s)...)
			}
		}
		if len(component) > 0 {
			g.connectedComponents = append(g.connectedComponents, component)
			for p := range component {
				visited[p] = struct{}{}
			}
		}
	}
}

func (g *GardenMap) calculateFencingCost(bulkDiscount bool) int {
	g.getConnectedComponents()
	totalCost := 0

	for _, comp := range g.connectedComponents {
		area := len(comp)
		perimeter := make(map[[2]complex128]bool)

		for z := range comp {
			for _, d := range []complex128{1, -1, 1i, -1i} {
				adj := z + d
				if _, ok := comp[adj]; !ok {
					perimeter[[2]complex128{z, adj}] = true
				}
			}
		}

		if bulkDiscount {
			filtered := make(map[[2]complex128]bool)
			for edge := range perimeter {
				keep := true
				for _, d := range []complex128{1, 1i} {
					if perimeter[[2]complex128{edge[0] + d, edge[1] + d}] {
						keep = false
					}
				}
				if keep {
					filtered[edge] = true
				}
			}
			totalCost += area * len(filtered)
		} else {
			totalCost += area * len(perimeter)
		}
	}
	return totalCost
}

func SolutionForPart1(input string) (int, error) {
	garden, err := NewGardenMapFromInput(input)
	if err != nil {
		return 0, errors.New("failed to initialize garden map")
	}
	return garden.calculateFencingCost(false), nil
}

func SolutionForPart2(input string) (int, error) {
	garden, err := NewGardenMapFromInput(input)
	if err != nil {
		return 0, errors.New("failed to initialize garden map")
	}
	return garden.calculateFencingCost(true), nil
}
