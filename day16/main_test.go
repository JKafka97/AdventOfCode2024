package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 16

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 109496, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 551, dayNumber, 2, SolutionForPart2)
}
