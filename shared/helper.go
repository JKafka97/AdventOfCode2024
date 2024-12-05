package helper

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func ReadFile(t *testing.T, filePath string) string {
	t.Helper()
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Error reading file %s: %v", filePath, err)
	}
	return strings.ReplaceAll(string(content), "\r\n", "\n")
}

func TestAdventOfCodeSolution(t *testing.T, expectedResult int, dayNumber int, partNumber int, solver func(string) (int, error)) {
	// Arrange
	input := ReadFile(t, fmt.Sprintf("../inputs/inputDay%d.txt", dayNumber))

	// Act
	result, err := solver(input)

	// Assert
	if err != nil {
		t.Errorf("TestSolvePart1() failed with error")
	}
	if result != expectedResult {
		t.Errorf("TestDay %d SolvePart %d failed: expected %d, got %d", dayNumber, partNumber, expectedResult, result)
	}
}
