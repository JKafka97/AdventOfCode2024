package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 17

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolutionString(t, "1,7,2,1,4,1,5,4,0", dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 37221261688304, dayNumber, 2, SolutionForPart2)
}
