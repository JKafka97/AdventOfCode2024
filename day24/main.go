package main

import (
	"bufio"
	"maps"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type adjacencyList map[string][]string

func SolutionForPart1(input string) (int, error) {
	value, dependencies := parseInput(input)

	return partOne(value, dependencies), nil
}

func SolutionForPart2(input string) (string, error) {
	_, dependencies := parseInput(input)

	return partTwo(dependencies), nil
}

type dependency struct {
	w1, w2 string
	op     string
}

func parseInput(input string) (map[string]int8, map[string]dependency) {
	instrRegex := regexp.MustCompile(`([a-z0-9]*) ([A-Z]*) ([a-z0-9]*) -> ([a-z0-9]*)`)
	wireValueRegex := regexp.MustCompile(`([a-zA-Z0-9]*): ([0-9])`)

	sc := bufio.NewScanner(strings.NewReader(input))

	value := make(map[string]int8)
	dependencies := make(map[string]dependency)

	for sc.Scan() && sc.Text() != "" {
		matches := wireValueRegex.FindStringSubmatch(sc.Text())
		w := matches[1]
		v := int8(matches[2][0] - '0')
		value[w] = v
	}

	for sc.Scan() {
		matches := instrRegex.FindStringSubmatch(sc.Text())
		w := matches[4]
		op := matches[2]
		w1, w2 := matches[1], matches[3]

		dependencies[w] = dependency{
			w1: w1,
			w2: w2,
			op: op,
		}
	}

	return value, dependencies
}

func partOne(value map[string]int8, dependencies map[string]dependency) (res int) {
	var resolve func(string) int8

	resolve = func(curr string) int8 {
		if v, ok := value[curr]; ok {
			return v
		}

		d := dependencies[curr]
		v1 := resolve(d.w1)
		v2 := resolve(d.w2)

		switch d.op {
		case "XOR":
			value[curr] = v1 ^ v2
		case "AND":
			value[curr] = v1 & v2
		case "OR":
			value[curr] = v1 | v2
		}

		return value[curr]
	}

	for n := range dependencies {
		resolve(n)
	}

	for n, v := range value {
		if n[0] == 'z' {
			temp, _ := strconv.Atoi(n[1:])
			res |= int(v) << temp
		}
	}

	return
}

func partTwo(dependencies map[string]dependency) (res string) {
	temp := make(map[string]bool)

	for w, d := range dependencies {
		if w[0] == 'z' {
			val, _ := strconv.Atoi(w[1:])
			if d.op != "XOR" && val != 45 {
				temp[w] = true
			}
		} else if !isXOrY(d.w1) && !isXOrY(d.w2) && d.w1[0] != d.w2[0] && d.op == "XOR" {
			temp[w] = true
		}

		if d.op == "XOR" && isXOrY(d.w1) && isXOrY(d.w2) && d.w1[0] != d.w2[0] {
			isValid := false
			for _, dp := range dependencies {
				if dp.op == "XOR" && (dp.w1 == w || dp.w2 == w) {
					isValid = true
				}
			}
			if !isValid {
				temp[w] = true
			}
		}

		if d.op == "AND" && isXOrY(d.w1) && isXOrY(d.w2) && d.w1[0] != d.w2[0] {
			isValid := false
			for _, dp := range dependencies {
				if dp.op == "OR" && (dp.w1 == w || dp.w2 == w) {
					isValid = true
				}
			}
			if !isValid {
				temp[w] = true
			}
		}
	}
	ans := slices.Collect(maps.Keys(temp))
	slices.Sort(ans)

	for _, w := range ans {
		res += w + ","
	}

	return res[:len(res)-1]
}

func isXOrY(wire string) bool {
	temp, _ := strconv.Atoi(wire[1:])
	return (wire[0] == 'x' || wire[0] == 'y') && temp != 0
}
