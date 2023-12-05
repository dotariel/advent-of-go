package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchematic_FindParts(t *testing.T) {
	testInput := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	schematic := NewSchematic(testInput)
	expected := []Part{
		{"467", Range{0, 2}},
		{"35", Range{2, 3}},
		{"633", Range{6, 8}},
		{"617", Range{0, 2}},
		{"592", Range{2, 4}},
		{"755", Range{6, 8}},
		{"664", Range{1, 3}},
		{"598", Range{5, 7}},
	}

	assert.Equal(t, expected, schematic.FindParts())
}

func TestSchematic_FindGears(t *testing.T) {
	testInput := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
....%.....
.875*362..`

	schematic := NewSchematic(testInput)
	expected := []Gear{
		{467, 35},
		{755, 598},
		{875, 362},
	}

	assert.Equal(t, expected, schematic.FindGears())
}
