package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 18

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 226, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolutionString(t, "60,46", dayNumber, 2, SolutionForPart2)
}
