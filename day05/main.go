package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// ProcessRulesAndUpdates processes the input string and returns rules and updates.
func ProcessRulesAndUpdates(input string) (rules [][2]string, updates [][]string, err error) {
	parts := strings.Split(input, "\n\n")
	if len(parts) < 2 {
		return nil, nil, fmt.Errorf("invalid input format")
	}

	// Parse rules and updates simultaneously
	for _, line := range strings.Split(parts[0], "\n") {
		parts := strings.Split(line, "|")
		if len(parts) == 2 {
			rules = append(rules, [2]string{parts[0], parts[1]})
		}
	}
	for _, line := range strings.Split(parts[1], "\n") {
		updates = append(updates, strings.Split(line, ","))
	}

	return rules, updates, nil
}

func SolutionForPart1(input string) (int, error) {
	rules, updates, err := ProcessRulesAndUpdates(input)
	if err != nil {
		return 0, err
	}

	result := 0
	for _, update := range updates {
		if !isUpdateIncorrect(update, rules) {
			midIndex, _ := strconv.Atoi(update[len(update)/2])
			result += midIndex
		}
	}
	return result, nil
}

func SolutionForPart2(input string) (int, error) {
	rules, updates, err := ProcessRulesAndUpdates(input)
	if err != nil {
		return 0, err
	}

	result := 0
	for _, update := range updates {
		if isUpdateIncorrect(update, rules) {
			result += fixAndCalculate(update, rules)
		}
	}
	return result, nil
}

// isUpdateIncorrect checks if any rule is violated in the update.
func isUpdateIncorrect(update []string, rules [][2]string) bool {
	for _, rule := range rules {
		if slices.Contains(update, rule[0]) && slices.Contains(update, rule[1]) &&
			slices.Index(update, rule[0]) > slices.Index(update, rule[1]) {
			return true
		}
	}
	return false
}

// fixAndCalculate fixes the update and calculates the result.
func fixAndCalculate(update []string, rules [][2]string) int {
	relevantRules := filterRules(update, rules)
	for i := 0; i < len(update); i++ {
		for _, value := range update[i:] {
			if !isValueMapped(value, relevantRules) {
				update = moveToPosition(update, value, i)
				break
			}
		}
		relevantRules = removeAppliedRules(relevantRules, update[i])
	}
	val, _ := strconv.Atoi(update[len(update)/2])
	return val
}

// isValueMapped checks if a value is mapped by any rule.
func isValueMapped(value string, rules [][2]string) bool {
	for _, rule := range rules {
		if rule[1] == value {
			return true
		}
	}
	return false
}

// filterRules filters the rules that apply to the given update.
func filterRules(update []string, rules [][2]string) [][2]string {
	var relevantRules [][2]string
	for _, rule := range rules {
		if slices.Contains(update, rule[0]) && slices.Contains(update, rule[1]) {
			relevantRules = append(relevantRules, rule)
		}
	}
	return relevantRules
}

// removeAppliedRules removes rules where the specified value has been applied.
func removeAppliedRules(rules [][2]string, value string) [][2]string {
	var remainingRules [][2]string
	for _, rule := range rules {
		if rule[0] != value {
			remainingRules = append(remainingRules, rule)
		}
	}
	return remainingRules
}

// moveToPosition moves a value in the slice to the specified position.
func moveToPosition(slice []string, value string, position int) []string {
	index := slices.Index(slice, value)
	if index != -1 {
		slice = append(slice[:index], slice[index+1:]...)
		slice = append(slice[:position], append([]string{value}, slice[position:]...)...)
	}
	return slice
}
