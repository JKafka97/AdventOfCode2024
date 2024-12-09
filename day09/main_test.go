package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 9

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 6242766523059, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 6272188244509, dayNumber, 2, SolutionForPart2)
}
