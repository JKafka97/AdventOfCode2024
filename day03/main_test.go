package main

import (
	helper "advent-of-code/shared"
	"testing"
)

const filePath = "../inputs/inputDay3.txt"

func TestDay3SolvePart1(t *testing.T) {
	// Arrange
	input := helper.ReadFile(t, filePath)

	// Act
	result, err := SolveDay3Part1(input)

	// Assert
	if err != nil {
		t.Errorf("TestDay3SolvePart1() failed with error")
	}
	expected := 161289189
	if result != expected {
		t.Errorf("TestDay3SolvePart1() failed: expected %d, got %d", expected, result)
	}
}

func TestDay3SolvePart2(t *testing.T) {
	// Arrange
	input := helper.ReadFile(t, filePath)

	// Act
	result, err := SolveDay3Part2(input)
	t.Log(result)

	// Assert
	if err != nil {
		t.Errorf("TestDay3SolvePart2() failed with error")
	}
	expected := 83595109
	if result != expected {
		t.Errorf("TestDay3SolvePart2() failed: expected %d, got %d", expected, result)
	}
}
