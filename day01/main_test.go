package main

import (
	"os"
	"testing"
)

const filePath = "input.txt"

func readFile(t *testing.T, filePath string) string {
	t.Helper()
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Error reading file %s: %v", filePath, err)
	}
	return string(content)
}

func TestDay1SolvePart1(t *testing.T) {
	// Arrange
	input := readFile(t, filePath)

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
	input := readFile(t, filePath)

	// Act
	result, err := SolveDay1Part2(input)

	// Assert
	if err != nil {
		t.Errorf("SolveDay1Part2() failed with error")
	}
	expected := 21070419
	if result != expected {
		t.Errorf("SolveDay1Part2() failed: expected %d, got %d", expected, result)
	}
}
