package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	Ax, Ay, Bx, By, Px, Py, CostA, CostB int
}

func parseInput(input string) ([]Machine, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var machines []Machine

	for i := 0; i < len(lines); i += 4 {
		var ax, ay, bx, by, px, py int
		_, err := fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &ax, &ay)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Button A: %w", err)
		}
		_, err = fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &bx, &by)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Button B: %w", err)
		}
		_, err = fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &px, &py)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Prize: %w", err)
		}

		machines = append(machines, Machine{Ax: ax, Ay: ay, Bx: bx, By: by, Px: px, Py: py, CostA: 3, CostB: 1})
	}

	return machines, nil
}

func minTokens(m Machine) (int, bool) {
	minTokens := -1
	found := false

	// Try to find a solution by iterating over a reasonable range for tokens
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			x := a*m.Ax + b*m.Bx
			y := a*m.Ay + b*m.By
			if x == m.Px && y == m.Py {
				tokens := a*m.CostA + b*m.CostB
				if !found || tokens < minTokens {
					minTokens = tokens
					found = true
				}
			}
		}
	}

	return minTokens, found
}

func SolutionForPart1(input string) (int, error) {
	machines, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	var totalTokens, prizesWon int
	for _, machine := range machines {
		tokens, found := minTokens(machine)
		if found {
			totalTokens += tokens
			prizesWon++
		}
	}

	if prizesWon == 0 {
		return 0, errors.New("no prizes could be won")
	}

	return totalTokens, nil
}

func parse(x string) ([2]int, [2]int, [2]int) {
	lines := strings.Split(x, "\n")
	// Regular expressions for parsing the values
	reA := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	reB := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	reP := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	aMatches := reA.FindStringSubmatch(lines[0])
	bMatches := reB.FindStringSubmatch(lines[1])
	pMatches := reP.FindStringSubmatch(lines[2])

	// Parse the matched groups into integers
	a := [2]int{parseInt(aMatches[1]), parseInt(aMatches[2])}
	b := [2]int{parseInt(bMatches[1]), parseInt(bMatches[2])}
	p := [2]int{parseInt(pMatches[1]), parseInt(pMatches[2])}

	return a, b, p
}

// Helper function to parse integers from string
func parseInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func SolutionForPart2(input string) (int, error) {
	lines := strings.Split(input, "\n\n")

	var totalTokens int
	for _, line := range lines {
		a, b, p := parse(line)
		// Shift large values to avoid negative numbers
		p[0] += 10000000000000
		p[1] += 10000000000000

		// Solve the system of linear equations using the determinant method
		i := (p[0]*b[1] - b[0]*p[1]) / (b[1]*a[0] - b[0]*a[1])
		j := (p[1] - a[1]*i) / b[1]

		// Ensure i and j are valid, then calculate the price
		if i >= 0 && j >= 0 && a[0]*i+b[0]*j == p[0] && a[1]*i+b[1]*j == p[1] {
			totalTokens += 3*i + j
		}
	}

	return totalTokens, nil
}
