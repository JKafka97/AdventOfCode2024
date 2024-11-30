package main

import "testing"

func TestSolvePart1(t *testing.T) {
	input := "nějaký vstup"
	expected := 14
	result := SolvePart1(input)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := "nějaký vstup"
	expected := 14
	result := SolvePart2(input)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
