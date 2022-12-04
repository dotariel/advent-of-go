package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	elves := Parse(input)

	expected := []Elf{
		{
			Food: []Food{
				{Calories: 1000},
				{Calories: 2000},
				{Calories: 3000},
			},
		},
		{
			Food: []Food{
				{Calories: 4000},
			},
		},
		{
			Food: []Food{
				{Calories: 5000},
				{Calories: 6000},
			},
		},
		{
			Food: []Food{
				{Calories: 7000},
				{Calories: 8000},
				{Calories: 9000},
			},
		},
		{
			Food: []Food{
				{Calories: 10000},
			},
		},
	}

	assert.Equal(t, expected, elves)
}

func Test_TotalCalories(t *testing.T) {
	elf := Elf{
		Food: []Food{
			{Calories: 1000},
			{Calories: 2000},
			{Calories: 3000},
		},
	}

	assert.Equal(t, elf.TotalCalories(), 6000)
}

func TestTopNCalories(t *testing.T) {
	e1 := Elf{Food: []Food{{Calories: 5}}}
	e2 := Elf{Food: []Food{{Calories: 3}}}
	e3 := Elf{Food: []Food{{Calories: 9}}}
	e4 := Elf{Food: []Food{{Calories: 4}}}
	e5 := Elf{Food: []Food{{Calories: 8}}}

	elves := []Elf{e1, e2, e3, e4, e5}

	assert.Equal(t, TopN(elves, 1), []Elf{e3})
	assert.Equal(t, TopN(elves, 2), []Elf{e3, e5})
	assert.Equal(t, TopN(elves, 3), []Elf{e3, e5, e1})
}
