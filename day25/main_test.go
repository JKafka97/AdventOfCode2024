package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 25

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 3255, dayNumber, 1, SolutionForPart1)
}
