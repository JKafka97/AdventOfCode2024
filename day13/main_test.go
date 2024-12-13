package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 13

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 29877, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 99423413811305, dayNumber, 2, SolutionForPart2)
}
