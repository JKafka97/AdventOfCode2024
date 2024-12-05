package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 1

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 1223326, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 21070419, dayNumber, 2, SolutionForPart2)
}
