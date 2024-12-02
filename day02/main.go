package main

import (
	"math"
	"strconv"
	"strings"
)

func SolveDay2Part1(input string) (int, error) {
	pairs := strings.Split(input, "\n")
	result := 0
	for _, pair := range pairs {
		numberRow, err := parseNumbers(pair)
		if err != nil {
			return 0, err
		}
		if len(numberRow) < 2 {
			continue
		}
		isIncreasing := numberRow[1] > numberRow[0]
		isValid := true
		for i := 1; i < len(numberRow); i++ {
			diff := numberRow[i] - numberRow[i-1]
			if absDiff := int(math.Abs(float64(diff))); absDiff < 1 || absDiff > 3 || (diff > 0) != isIncreasing {
				isValid = false
				break
			}
		}
		if isValid {
			result++
		}
	}
	return result, nil
}

func SolveDay2Part2(input string) (int, error) {
	rows := strings.Split(input, "\n")
	result := 0

	for _, row := range rows {
		numbers, err := parseNumbers(row)
		if err != nil {
			return 0, err
		}
		if len(numbers) < 2 {
			continue
		}

		if isValid(numbers) || canBeValidWithDampener(numbers) {
			result++
		}
	}

	return result, nil
}

func parseNumbers(row string) ([]int, error) {
	fields := strings.Fields(row)
	numbers := make([]int, len(fields))
	for i, f := range fields {
		num, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
		numbers[i] = num
	}
	return numbers, nil
}

func isValid(seq []int) bool {
	isIncreasing := seq[1] > seq[0]
	for i := 1; i < len(seq); i++ {
		diff := seq[i] - seq[i-1]
		absDiff := int(math.Abs(float64(diff)))
		if absDiff < 1 || absDiff > 3 || (diff > 0) != isIncreasing {
			return false
		}
	}
	return true
}

func canBeValidWithDampener(seq []int) bool {
	for i := range seq {
		newSeq := append(append([]int{}, seq[:i]...), seq[i+1:]...)
		if len(newSeq) > 1 && isValid(newSeq) {
			return true
		}
	}
	return false
}
