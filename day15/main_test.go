package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 15

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 1577255, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 1597035, dayNumber, 2, SolutionForPart2)
}
