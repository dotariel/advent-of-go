package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	game := NewGame(input)
	expected := Game{
		Id: 1,
		CubeSets: []CubeSet{
			{"blue": 3, "red": 4},
			{"blue": 6, "green": 2, "red": 1},
			{"green": 2},
		},
	}

	assert.Equal(t, expected, game)
}

func TestGame_GetMinimumCubeSet(t *testing.T) {
	testCases := []struct {
		input    string
		expected CubeSet
	}{
		{input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", expected: CubeSet{"red": 4, "green": 2, "blue": 6}},
		{input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", expected: CubeSet{"red": 1, "green": 3, "blue": 4}},
		{input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", expected: CubeSet{"red": 20, "green": 13, "blue": 6}},
		{input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", expected: CubeSet{"red": 14, "green": 3, "blue": 15}},
		{input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", expected: CubeSet{"red": 6, "green": 3, "blue": 2}},
	}

	for _, tc := range testCases {
		game := NewGame(tc.input)

		assert.Equal(t, tc.expected, game.GetMinimumCubeSet())
	}
}

func Test_extractId(t *testing.T) {
	assert.Equal(t, 0, extractId("Game"))
	assert.Equal(t, 1, extractId("Game 1"))
	assert.Equal(t, 22, extractId("Game 22"))
}
