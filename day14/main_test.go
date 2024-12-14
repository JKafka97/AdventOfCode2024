package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 14

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 230436441, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 8270, dayNumber, 2, SolutionForPart2)
}
