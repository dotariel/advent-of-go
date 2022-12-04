package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRucksack(t *testing.T) {
	input := "vJrwpWtwJgWrhcsFMMfFFhFp"

	r := NewRucksack(input)

	assert.Equal(t, 2, len(r.Compartments))
	assert.Equal(t, 12, len(r.Compartments[0].Items))
	assert.Equal(t, 12, len(r.Compartments[1].Items))
}

func TestRucksack_FindDuplicates(t *testing.T) {
	testCases := []struct {
		input string
		dupes []rune
	}{
		{input: "vJrwpWtwJgWrhcsFMMfFFhFp", dupes: []rune{'p'}},
		{input: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", dupes: []rune{'L'}},
		{input: "PmmdzqPrVvPwwTWBwg", dupes: []rune{'P'}},
		{input: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", dupes: []rune{'v'}},
		{input: "ttgJtRGJQctTZtZT", dupes: []rune{'t'}},
		{input: "CrZsJsPPZsGzwwsLwLmpwMDw", dupes: []rune{'s'}},
	}

	for _, tc := range testCases {
		dupes := NewRucksack(tc.input).FindDuplicates()

		assert.Equal(t, tc.dupes, dupes)
	}
}

func TestGetPriority(t *testing.T) {
	testCases := []struct {
		char     rune
		priority int
	}{
		{char: 'p', priority: 16},
		{char: 'L', priority: 38},
		{char: 'P', priority: 42},
		{char: 'v', priority: 22},
		{char: 't', priority: 20},
		{char: 's', priority: 19},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.priority, GetPriority(tc.char))
	}
}

func TestElfGroup_FindBadge(t *testing.T) {
	g1 := ElfGroup{
		NewRucksack("vJrwpWtwJgWrhcsFMMfFFhFp"),
		NewRucksack("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"),
		NewRucksack("PmmdzqPrVvPwwTWBwg"),
	}

	g2 := ElfGroup{
		NewRucksack("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"),
		NewRucksack("ttgJtRGJQctTZtZT"),
		NewRucksack("CrZsJsPPZsGzwwsLwLmpwMDw"),
	}

	assert.Equal(t, 'r', g1.FindBadge())
	assert.Equal(t, 'Z', g2.FindBadge())
}
