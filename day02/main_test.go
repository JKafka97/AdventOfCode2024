package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 2

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 564, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 604, dayNumber, 2, SolutionForPart2)
}
