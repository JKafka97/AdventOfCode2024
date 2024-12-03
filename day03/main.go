package main

import (
	"regexp"
	"strconv"
	"strings"
)

// SolveDay3Part1 computes the sum of products from "mul(x, y)" in the input.
func SolveDay3Part1(input string) (int, error) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	result := 0
	for _, match := range re.FindAllStringSubmatch(input, -1) {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		result += x * y
	}
	return result, nil
}

func SolveDay3Part2(input string) (int, error) {
	rows := strings.Split(input, "\n")
	result := 0

	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	mulEnabled := true

	for _, line := range rows {
		parts := mulRegex.FindAllStringSubmatchIndex(line, -1) // Positions of mul matches
		doMatches := doRegex.FindAllStringIndex(line, -1)
		dontMatches := dontRegex.FindAllStringIndex(line, -1)

		i := 0
		for i < len(line) {
			if len(dontMatches) > 0 && i == dontMatches[0][0] {
				mulEnabled = false
				dontMatches = dontMatches[1:] // Remove processed don't()
				i += len("don't()")
				continue
			}
			if len(doMatches) > 0 && i == doMatches[0][0] {
				mulEnabled = true
				doMatches = doMatches[1:] // Remove processed do()
				i += len("do()")
				continue
			}
			if len(parts) > 0 && i == parts[0][0] {
				if mulEnabled {
					match := line[parts[0][0]:parts[0][1]]
					mulParts := mulRegex.FindStringSubmatch(match)
					if len(mulParts) >= 3 { // Ensure the match contains both numbers
						x, _ := strconv.Atoi(mulParts[1])
						y, _ := strconv.Atoi(mulParts[2])
						result += x * y
					}
				}
				i += parts[0][1] - parts[0][0] // Skip past the mul()
				parts = parts[1:]
				continue
			}

			i++
		}

	}
	return result, nil
}
