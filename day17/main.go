package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SolutionForPart1(input string) (string, error) {

	lines := strings.Split(strings.TrimSpace(input), "\n")
	A, _ := strconv.Atoi(strings.Split(lines[0], " ")[2])
	B, _ := strconv.Atoi(strings.Split(lines[1], " ")[2])
	C, _ := strconv.Atoi(strings.Split(lines[2], " ")[2])
	programStr := strings.Split(strings.Split(lines[4], " ")[1], ",")
	program := make([]int, len(programStr))
	for i, s := range programStr {
		program[i], _ = strconv.Atoi(s)
	}

	run := func(program []int, A, B, C int) []int {
		ptr := 0
		var result []int

		combo := func() int {
			x := program[ptr+1]
			if x <= 3 {
				return x
			}
			if x == 4 {
				return A
			}
			if x == 5 {
				return B
			}
			if x == 6 {
				return C
			}
			return 0
		}

		lit := func() int {
			return program[ptr+1]
		}

		for {
			if ptr >= len(program) {
				break
			}

			opcode := program[ptr]

			switch opcode {
			case 0:
				A = A / (1 << combo())
			case 1:
				B = B ^ lit()
			case 2:
				B = combo() % 8
			case 3:
				if A != 0 {
					ptr = lit()
					continue
				} else {
					ptr += 2
				}
			case 4:
				B = B ^ C
			case 5:
				result = append(result, combo()%8)
			case 6:
				B = A / (1 << combo())
			case 7:
				C = A / (1 << combo())
			}

			if opcode != 3 {
				ptr += 2
			}
		}

		return result
	}

	ans := run(program, A, B, C)
	result := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ans)), ","), "[]")

	return result, nil
}

func SolutionForPart2(input string) (int, error) {

	parts := strings.Split(input, "\n\n")
	reg := parts[0]
	progStr := strings.Split(parts[1], ": ")[1]
	progStrs := strings.Split(progStr, ",")
	prog := make([]int, len(progStrs))
	for i, s := range progStrs {
		prog[i], _ = strconv.Atoi(s)
	}

	regLines := strings.Split(reg, "\n")
	rega, _ := strconv.Atoi(strings.Split(regLines[0], ": ")[1])
	regb, _ := strconv.Atoi(strings.Split(regLines[1], ": ")[1])
	regc, _ := strconv.Atoi(strings.Split(regLines[2], ": ")[1])

	rega = 0
	j := 1
	istart := 0
	for j <= len(prog) && j >= 0 {
		rega <<= 3
		found := false
		for i := istart; i < 8; i++ {
			if equal(prog[len(prog)-j:], run(prog, rega+i, regb, regc)) {
				found = true
				break
			}
		}
		if !found {
			j--
			rega >>= 3
			istart = rega%8 + 1
			rega >>= 3
			continue
		}
		j++
		rega += istart
		istart = 0
	}
	return rega, nil
}

func getCombo(oper, rega, regb, regc int) int {
	switch oper {
	case 0, 1, 2, 3:
		return oper
	case 4:
		return rega
	case 5:
		return regb
	case 6:
		return regc
	default:
		return 0
	}
}

func run(prog []int, rega, regb, regc int) []int {
	ip := 0
	out := []int{}
	for ip < len(prog) {
		oper := prog[ip+1]
		combo := getCombo(oper, rega, regb, regc)
		switch prog[ip] {
		case 0:
			rega /= 1 << combo
		case 1:
			regb ^= oper
		case 2:
			regb = combo % 8
		case 3:
			if rega != 0 {
				ip = oper
				continue
			}
		case 4:
			regb ^= regc
		case 5:
			out = append(out, combo%8)
		case 6:
			regb = rega / (1 << combo)
		case 7:
			regc = rega / (1 << combo)
		}
		ip += 2
	}
	return out
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
