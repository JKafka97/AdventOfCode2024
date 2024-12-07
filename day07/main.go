package main

import (
	"fmt"
	"strconv"
	"strings"
)

func evaluate(nums []int, combo []rune) (int, error) {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		switch combo[i-1] {
		case '+':
			result += nums[i]
		case '*':
			result *= nums[i]
		case '|':
			concatenated, err := strconv.Atoi(fmt.Sprintf("%d%d", result, nums[i]))
			if err != nil {
				return 0, err
			}
			result = concatenated
		default:
			return 0, fmt.Errorf("unknown operator: %c", combo[i-1])
		}
	}
	return result, nil
}

func generateCombinations(chars string, n int) [][]rune {
	if n == 0 {
		return [][]rune{{}}
	}
	subCombos := generateCombinations(chars, n-1)
	var combos [][]rune
	for _, sub := range subCombos {
		for _, ch := range chars {
			combos = append(combos, append([]rune{ch}, sub...))
		}
	}
	return combos
}

func solve(input, operators string) (int, error) {
	rows := strings.Split(strings.TrimSpace(input), "\n")
	total := 0

	for _, line := range rows {
		parts := strings.Fields(line)
		expected, err := strconv.Atoi(strings.TrimSuffix(parts[0], ":"))
		if err != nil {
			return 0, fmt.Errorf("invalid target number: %v", err)
		}

		nums, err := parseNumbers(parts[1:])
		if err != nil {
			return 0, err
		}

		combinations := generateCombinations(operators, len(nums)-1)
		for _, combo := range combinations {
			if result, err := evaluate(nums, combo); err == nil && result == expected {
				total += expected
				break
			}
		}
	}
	return total, nil
}

func parseNumbers(parts []string) ([]int, error) {
	var nums []int
	for _, p := range parts {
		num, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %v", err)
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func SolutionForPart1(input string) (int, error) {
	return solve(input, "*+")
}

func SolutionForPart2(input string) (int, error) {
	return solve(input, "*+|")
}
