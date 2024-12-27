package main

import (
	"strings"
)

func SolutionForPart1(input string) (int, error) {
	result := solve(input, false)
	return result, nil
}

func SolutionForPart2(input string) (int, error) {
	result := solve(input, true)
	return result, nil
}

func solve(input string, countWays bool) int {
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	patterns := strings.Split(split[0], ", ")

	var ways func(string) int
	cache := map[string]int{}

	ways = func(design string) (n int) {
		if n, ok := cache[design]; ok {
			return n
		}
		defer func() { cache[design] = n }()

		if design == "" {
			return 1
		}
		for _, s := range patterns {
			if strings.HasPrefix(design, s) {
				n += ways(design[len(s):])
			}
		}
		return n
	}

	result := 0
	for _, s := range strings.Fields(split[1]) {
		if w := ways(s); w > 0 {
			if countWays {
				result += w
			} else {
				result++
			}
		}
	}
	return result
}
