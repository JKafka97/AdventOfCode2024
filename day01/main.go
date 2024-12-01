package main

import (
	"sort"
	"strconv"
	"strings"
)

func SolveDay1Part1(input string) (int, error) {
	pairs := strings.Split(input, "\r\n")
	firstRow := make([]int, 0)
	secondRow := make([]int, 0)
	result := 0
	for _, pair := range pairs {
		numbers := strings.Fields(pair)
		firstNumber, err := strconv.Atoi(numbers[0])
		if err != nil {
			return 0, err
		}
		firstRow = append(firstRow, firstNumber)
		secondNumber, err := strconv.Atoi(numbers[1])
		if err != nil {
			return 0, err
		}
		secondRow = append(secondRow, secondNumber)
	}
	sort.Ints(firstRow)
	sort.Ints(secondRow)
	for i := range firstRow {
		result += absInt(firstRow[i] - secondRow[i])
	}
	return result, nil
}

func SolveDay1Part2(input string) (int, error) {
	pairs := strings.Split(input, "\r\n")
	firstRow := make([]int, 0)
	secondRow := make([]int, 0)
	result := 0
	for _, pair := range pairs {
		numbers := strings.Fields(pair)
		firstNumber, err := strconv.Atoi(numbers[0])
		if err != nil {
			return 0, err
		}
		firstRow = append(firstRow, firstNumber)
		secondNumber, err := strconv.Atoi(numbers[1])
		if err != nil {
			return 0, err
		}
		secondRow = append(secondRow, secondNumber)
	}

	secondRowNumbersCountMap := countElements(secondRow)
	for _, numberInFirstRow := range firstRow {
		result += numberInFirstRow * secondRowNumbersCountMap[numberInFirstRow]
	}
	return result, nil
}

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func countElements(slice []int) map[int]int {
	counts := make(map[int]int)
	for _, item := range slice {
		counts[item]++
	}
	return counts
}
