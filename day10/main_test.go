package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 10

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 629, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 1242, dayNumber, 2, SolutionForPart2)
}
