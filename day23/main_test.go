package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 23

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 1368, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolutionString(t, "dd,ig,il,im,kb,kr,pe,ti,tv,vr,we,xu,zi", dayNumber, 2, SolutionForPart2)
}
