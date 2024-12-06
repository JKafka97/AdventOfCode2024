package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 6

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 4776, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 1586, dayNumber, 2, SolutionForPart2)
}
