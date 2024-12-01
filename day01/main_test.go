package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const filePath = "../inputs/inputDay1.txt"

func TestDay1SolvePart1(t *testing.T) {
	// Arrange
	input := helper.ReadFile(t, filePath)

	// Act
	result, err := SolveDay1Part1(input)

	// Assert
	if err != nil {
		t.Errorf("SolveDay1Part1() failed with error")
	}
	expected := 1223326
	if result != expected {
		t.Errorf("SolveDay1Part1() failed: expected %d, got %d", expected, result)
	}
}

func TestDay1SolvePart2(t *testing.T) {
	// Arrange
	input := helper.ReadFile(t, filePath)

	// Act
	result, err := SolveDay1Part2(input)
	t.Log(result)

	// Assert
	if err != nil {
		t.Errorf("SolveDay1Part2() failed with error")
	}
	expected := 21070419
	if result != expected {
		t.Errorf("SolveDay1Part2() failed: expected %d, got %d", expected, result)
	}
}
