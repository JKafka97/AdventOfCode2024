package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 19

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 324, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 575227823167869, dayNumber, 2, SolutionForPart2)
}
