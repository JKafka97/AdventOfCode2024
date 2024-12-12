package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 11

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 220722, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 261952051690787, dayNumber, 2, SolutionForPart2)
}
