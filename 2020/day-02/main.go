package main

import (
	"dotariel/inputreader"
	"strconv"
	"strings"
)

type policy struct {
	char rune
	min  int
	max  int
}

var entries = make([]policy, 0)

func main() {
	valid := 0

	for _, val := range inputreader.ReadStrings("input.txt", "\n") {
		policy := parsePolicy(val)
		password := parsePassword(val)

		if policy.IsPositionalValid((password)) {
			valid++
		}
	}

	println("Part 1: %v", valid)
}

func parsePolicy(s string) policy {
	parts := strings.Split(s, " ")
	allowed := strings.Split(parts[0], "-")
	char := parts[1][0]

	min, _ := strconv.Atoi(allowed[0])
	max, _ := strconv.Atoi(allowed[1])

	return policy{
		char: rune(char),
		min:  min,
		max:  max,
	}
}

func parsePassword(s string) string {
	return strings.Split(s, ":")[1][1:]
}

func (p policy) IsValid(password string) bool {
	charCount := 0

	for _, c := range password {
		if c == p.char {
			charCount++
		}
	}

	return charCount >= p.min && charCount <= p.max
}

func (p policy) IsPositionalValid(password string) bool {
	pos1 := p.min - 1
	pos2 := p.max - 1

	return (rune(password[pos1]) == p.char || (rune(password[pos2]) == p.char)) &&
		rune(password[pos1]) != (rune(password[pos2]))
}
