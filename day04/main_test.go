package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 4

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 2454, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 1858, dayNumber, 2, SolutionForPart2)
}
