package main

import (
	"bufio"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func SolutionForPart1(input string) (int, error) {
	return calculateSolution(parseInput(input), 2), nil
}

func SolutionForPart2(input string) (int, error) {
	return calculateSolution(parseInput(input), 25), nil
}

type index struct {
	row, col int
}

type direction struct {
	dRow, dCol int
}

var (
	directionMap = map[rune]direction{
		'^': {-1, 0},
		'v': {1, 0},
		'>': {0, 1},
		'<': {0, -1},
	}

	numericKeypad = map[rune]index{
		'7': {0, 0}, '8': {0, 1}, '9': {0, 2},
		'4': {1, 0}, '5': {1, 1}, '6': {1, 2},
		'1': {2, 0}, '2': {2, 1}, '3': {2, 2},
		'0': {3, 1}, 'A': {3, 2},
	}

	directionKeypad = map[rune]index{
		'^': {0, 1}, 'A': {0, 2},
		'<': {1, 0}, 'v': {1, 1}, '>': {1, 2},
	}

	reverseDirectionKeypad = reverseMap(directionKeypad)
	reverseNumericKeypad   = reverseMap(numericKeypad)
	minDistanceCache       = make(map[string]int)
	pathsCache             = make(map[string][]string)
)

func parseInput(input string) []string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func calculateSolution(input []string, depth int) int {
	var total int
	for _, line := range input {
		cost := computeCost("A"+line, depth)
		coeff, _ := strconv.Atoi(line[:len(line)-1])
		total += cost * coeff
	}
	return total
}

func computeCost(str string, depth int) int {
	var totalCost int
	for i := 0; i < len(str)-1; i++ {
		totalCost += calculatePairCost(rune(str[i]), rune(str[i+1]), numericKeypad, reverseNumericKeypad, depth)
	}
	return totalCost
}

func calculatePairCost(a, b rune, charToIndex map[rune]index, indexToChar map[index]rune, depth int) int {
	keypadType := 'd'
	if _, exists := charToIndex['0']; exists {
		keypadType = 'n'
	}
	cacheKey := fmt.Sprintf("%c%c%c%d", a, b, keypadType, depth)

	if dist, exists := minDistanceCache[cacheKey]; exists {
		return dist
	}

	if depth == 0 {
		minLength := math.MaxInt
		for _, path := range findAllPaths(a, b, directionKeypad, reverseDirectionKeypad) {
			minLength = min(minLength, len(path))
		}
		return minLength
	}

	allPaths := findAllPaths(a, b, charToIndex, indexToChar)
	minCost := math.MaxInt

	for _, path := range allPaths {
		path = "A" + path
		var currentCost int

		for i := 0; i < len(path)-1; i++ {
			currentCost += calculatePairCost(rune(path[i]), rune(path[i+1]), directionKeypad, reverseDirectionKeypad, depth-1)
		}
		minCost = min(minCost, currentCost)
	}

	minDistanceCache[cacheKey] = minCost
	return minCost
}

func findAllPaths(a, b rune, charToIndex map[rune]index, indexToChar map[index]rune) []string {
	cacheKey := fmt.Sprintf("%c %c", a, b)
	if paths, exists := pathsCache[cacheKey]; exists {
		return paths
	}
	var allPaths []string
	depthFirstSearch(charToIndex[a], charToIndex[b], []rune{}, charToIndex, indexToChar, make(map[index]bool), &allPaths)
	pathsCache[cacheKey] = allPaths
	return allPaths
}

func depthFirstSearch(current, end index, path []rune, charToIndex map[rune]index, indexToChar map[index]rune, visited map[index]bool, allPaths *[]string) {
	if current == end {
		*allPaths = append(*allPaths, string(path)+"A")
		return
	}
	visited[current] = true
	for char, dir := range directionMap {
		nextIndex := index{current.row + dir.dRow, current.col + dir.dCol}
		if _, exists := indexToChar[nextIndex]; exists && !visited[nextIndex] {
			newPath := slices.Clone(path)
			depthFirstSearch(nextIndex, end, append(newPath, char), charToIndex, indexToChar, visited, allPaths)
		}
	}
	visited[current] = false
}

func reverseMap(m map[rune]index) map[index]rune {
	reverseMap := make(map[index]rune)
	for r, i := range m {
		reverseMap[i] = r
	}
	return reverseMap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
