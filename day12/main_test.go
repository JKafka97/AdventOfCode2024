package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 12

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 1396298, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 853588, dayNumber, 2, SolutionForPart2)
}
