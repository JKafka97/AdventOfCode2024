package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 3

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 161289189, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 83595109, dayNumber, 2, SolutionForPart2)
}
