package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 8

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 359, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 1293, dayNumber, 2, SolutionForPart2)
}
