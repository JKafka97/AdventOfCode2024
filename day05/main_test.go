package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 5

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 5108, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 7380, dayNumber, 2, SolutionForPart2)
}
