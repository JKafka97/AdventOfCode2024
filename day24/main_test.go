package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const dayNumber = 24

func TestSolutionForPart1(t *testing.T) {
	helper.TestAdventOfCodeSolution(t, 49574189473968, dayNumber, 1, SolutionForPart1)
}

func TestSolutionForPart2(t *testing.T) {
	helper.TestAdventOfCodeSolutionString(t, "ckb,kbs,ksv,nbd,tqq,z06,z20,z39", dayNumber, 2, SolutionForPart2)
}
