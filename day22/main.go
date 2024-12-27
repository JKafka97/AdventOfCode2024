package main

import (
	"bufio"
	"strconv"
	"strings"
)

func SolutionForPart1(input string) (int, error) {
	result, _ := computeSolutions(input)
	return result, nil
}

func SolutionForPart2(input string) (int, error) {
	_, result := computeSolutions(input)
	return result, nil
}

const MODULO = (1 << 24) - 1

func computeSolutions(input string) (int, int) {
	reader := bufio.NewReader(strings.NewReader(input))
	total := 0
	sequenceCounts := make(map[[4]int]int)

	for {
		line, _, err := reader.ReadLine()
		if line == nil || err != nil {
			break
		}
		value, _ := strconv.Atoi(string(line))
		sequence := make(map[[4]int]int)
		total += processValue(value, 2000, &sequence)

		for key, count := range sequence {
			sequenceCounts[key] += count
		}
	}

	best := 0
	for _, count := range sequenceCounts {
		if count > best {
			best = count
		}
	}
	return total, best
}

func processValue(value, iterations int, sequence *map[[4]int]int) int {
	window := [4]int{0, 0, 0, 0}
	for i := 0; i < iterations; i++ {
		previous := value
		value = transformValue(value)
		difference := value%10 - previous%10
		digit := value % 10

		if i < 4 {
			window[i] = difference
			if i < 3 {
				continue
			}
		} else {
			copy(window[:3], window[1:])
			window[3] = difference
		}

		if _, exists := (*sequence)[window]; !exists {
			(*sequence)[window] = digit
		}
	}
	return value
}

func transformValue(value int) int {
	value = (value ^ (value << 6)) & MODULO
	value = (value ^ (value >> 5)) & MODULO
	return (value ^ (value << 11)) & MODULO
}
