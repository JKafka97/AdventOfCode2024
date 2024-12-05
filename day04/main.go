package main

import (
	"regexp"
	"strings"
)

func SolutionForPart1(input string) (int, error) {
	data := strings.Split(input, "\n")
	pattern := "XMAS"
	count := 0
	size := len(data)
	re := regexp.MustCompile(pattern)
	countPatternInString := func(s string) int {
		return len(re.FindAllString(s, -1)) + len(re.FindAllString(reverseString(s), -1))
	}

	// Rows and Columns
	for i := 0; i < size; i++ {
		// Row check
		count += countPatternInString(data[i])

		// Column check
		var colBuilder strings.Builder
		for _, row := range data {
			colBuilder.WriteByte(row[i])
		}
		count += countPatternInString(colBuilder.String())
	}

	// Diagonal Up + Right
	for i := 3; i < size; i++ {
		for j := 0; j < size-3; j++ {
			diag := string(data[i][j]) + string(data[i-1][j+1]) + string(data[i-2][j+2]) + string(data[i-3][j+3])
			count += countPatternInString(diag)
		}
	}

	// Diagonal Down + Right
	for i := 0; i < size-3; i++ {
		for j := 0; j < size-3; j++ {
			diag := string(data[i][j]) + string(data[i+1][j+1]) + string(data[i+2][j+2]) + string(data[i+3][j+3])
			count += countPatternInString(diag)
		}
	}

	return count, nil
}

func SolutionForPart2(input string) (int, error) {
	data := strings.Split(input, "\n")
	count := 0
	size := len(data)

	// Check for 'A' and valid diagonals
	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			if data[i][j] == 'A' {
				diag1 := string(data[i-1][j-1]) + string(data[i+1][j+1])
				diag2 := string(data[i+1][j-1]) + string(data[i-1][j+1])
				if (diag1 == "MS" || diag1 == "SM") && (diag2 == "MS" || diag2 == "SM") {
					count++
				}
			}
		}
	}

	return count, nil
}

// Helper function to reverse a string.
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
