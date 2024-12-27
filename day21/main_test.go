package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 21

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 155252, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 195664513288128, dayNumber, 2, SolutionForPart2)
}
