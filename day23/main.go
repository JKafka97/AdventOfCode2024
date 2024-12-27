package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
)

type adjacencyList map[string][]string

func SolutionForPart1(input string) (int, error) {
	return findCycles(parseInput(input)), nil
}

func SolutionForPart2(input string) (string, error) {
	return findMaxClique(parseInput(input)), nil
}

func parseInput(input string) adjacencyList {

	scanner := bufio.NewScanner(strings.NewReader(input))

	adjList := make(adjacencyList)

	for scanner.Scan() {
		line := scanner.Text()
		sepIndex := strings.Index(line, "-")
		nodeA, nodeB := line[:sepIndex], line[sepIndex+1:]
		adjList[nodeA] = append(adjList[nodeA], nodeB)
		adjList[nodeB] = append(adjList[nodeB], nodeA)
	}

	return adjList
}

func findCycles(adjList adjacencyList) int {
	cycles := make(map[string][]string)
	for node1 := range adjList {
		if node1[0] != 't' {
			continue
		}
		for _, node2 := range adjList[node1] {
			for _, node3 := range adjList[node2] {
				if slices.Contains(adjList[node3], node1) {
					cycle := []string{node1, node2, node3}
					slices.Sort(cycle)
					key := fmt.Sprintf("%s%s%s", cycle[0], cycle[1], cycle[2])
					cycles[key] = cycle
				}
			}
		}
	}
	return len(cycles)
}

func findMaxClique(adjList adjacencyList) (password string) {
	nodes := make([]string, 0, len(adjList))

	for node := range adjList {
		nodes = append(nodes, node)
	}

	var largestClique []string
	var maxCliqueSize int

	BronKerbosch([]string{}, nodes, []string{}, &largestClique, &maxCliqueSize, adjList)

	slices.Sort(largestClique)
	for _, node := range largestClique {
		password += node + ","
	}

	return password[:len(password)-1]
}

func BronKerbosch(R, P, X []string, largestClique *[]string, maxCliqueSize *int, adjList adjacencyList) {
	if len(P) == 0 && len(X) == 0 && *maxCliqueSize < len(R) {
		rCopy := slices.Clone(R)
		*largestClique = rCopy
		*maxCliqueSize = len(rCopy)
		return
	}

	pCopy := slices.Clone(P)
	for _, v := range pCopy {
		newR := append(R, v)
		neighbours := adjList[v]

		newP := intersect(P, neighbours)
		newX := intersect(X, neighbours)

		BronKerbosch(newR, newP, newX, largestClique, maxCliqueSize, adjList)

		vIdx := slices.Index(P, v)
		P = slices.Delete(P, vIdx, vIdx+1)

		X = append(X, v)
	}
}

func intersect(a, b []string) (res []string) {
	intersection := make(map[string]bool)
	for _, val := range a {
		intersection[val] = true
	}
	for _, val := range b {
		if intersection[val] {
			res = append(res, val)
		}
	}
	return
}
