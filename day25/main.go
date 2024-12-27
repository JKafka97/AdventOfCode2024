package main

import (
	"bufio"
	"slices"
	"strings"
)

type adjacencyList map[string][]string

func SolutionForPart1(input string) (int, error) {
	locks, keys := parseInput(input)
	return calculatePartOne(locks, keys), nil
}

func parseInput(input string) ([][]int, [][]int) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var locks, keys [][]int

	for scanner.Scan() {
		curr := make([]int, 5)
		isLock := false

		for i := 0; i < 7; i++ {
			line := scanner.Text()
			for c, char := range line {
				if char == '#' {
					curr[c]++
				}
			}
			if i == 0 && slices.Equal(curr, []int{1, 1, 1, 1, 1}) {
				isLock = true
			}
			if !scanner.Scan() {
				break
			}
		}

		for i := range curr {
			curr[i]--
		}

		if isLock {
			locks = append(locks, curr)
		} else {
			keys = append(keys, curr)
		}
	}

	return locks, keys
}

func calculatePartOne(locks, keys [][]int) (count int) {
	for _, lock := range locks {
		for _, key := range keys {
			isValid := true
			for i := range lock {
				if lock[i]+key[i] > 5 {
					isValid = false
					break
				}
			}
			if isValid {
				count++
			}
		}
	}
	return count
}
