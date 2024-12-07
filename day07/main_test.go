package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 7

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 5030892084481, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 91377448644679, dayNumber, 2, SolutionForPart2)
}
