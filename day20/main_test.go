package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 20

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 1363, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 1007186, dayNumber, 2, SolutionForPart2)
}
