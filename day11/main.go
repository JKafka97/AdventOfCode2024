package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var memory = make(map[string]int)

func solve(stone int, blinks int) int {
	if blinks == 0 {
		return 1
	}

	key := fmt.Sprintf("%d,%d", stone, blinks)
	if val, found := memory[key]; found {
		return val
	}

	var val int
	if stone == 0 {
		val = solve(1, blinks-1)
	} else if strStone := strconv.Itoa(stone); len(strStone)%2 == 0 {
		mid := len(strStone) / 2
		left, _ := strconv.Atoi(strStone[:mid])
		right, _ := strconv.Atoi(strStone[mid:])
		val = solve(left, blinks-1) + solve(right, blinks-1)
	} else {
		val = solve(stone*2024, blinks-1)
	}

	memory[key] = val
	return val
}

func SolutionForPart1(input string) (int, error) {
	if input == "" {
		return 0, errors.New("input cannot be empty")
	}
	stones := strings.Fields(input)
	sum := 0
	for _, s := range stones {
		stone, _ := strconv.Atoi(s)
		sum += solve(stone, 25)
	}
	return sum, nil
}

func SolutionForPart2(input string) (int, error) {
	if input == "" {
		return 0, errors.New("input cannot be empty")
	}
	stones := strings.Fields(input)
	sum := 0
	for _, s := range stones {
		stone, _ := strconv.Atoi(s)
		sum += solve(stone, 75)
	}
	return sum, nil
}

