package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBag_Validate(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", expected: true},
		{input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", expected: true},
		{input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", expected: false},
		{input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", expected: false},
		{input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", expected: true},
	}

	bag := NewBag(CubeSet{"red": 12, "green": 13, "blue": 14})

	for _, tc := range testCases {
		game := NewGame(tc.input)

		assert.Equal(t, tc.expected, bag.Validate(game))
	}
}
