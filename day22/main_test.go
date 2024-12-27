package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 22

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 15303617151, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 1727, dayNumber, 2, SolutionForPart2)
}
