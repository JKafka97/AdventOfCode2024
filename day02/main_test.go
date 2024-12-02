package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const filePath = "../inputs/inputDay2.txt"

func TestDay2SolvePart1(t *testing.T) {
	// Arrange
	input := helper.ReadFile(t, filePath)

	// Act
	result, err := SolveDay2Part1(input)

	// Assert
	if err != nil {
		t.Errorf("TestDay2SolvePart1() failed with error")
	}
	expected := 564
	if result != expected {
		t.Errorf("TestDay2SolvePart1() failed: expected %d, got %d", expected, result)
	}
}

func TestDay2SolvePart2(t *testing.T) {
	// Arrange
	input := helper.ReadFile(t, filePath)

	// Act
	result, err := SolveDay2Part2(input)
	t.Log(result)

	// Assert
	if err != nil {
		t.Errorf("TestDay2SolvePart2() failed with error")
	}
	expected := 604
	if result != expected {
		t.Errorf("TestDay2SolvePart2() failed: expected %d, got %d", expected, result)
	}
}
