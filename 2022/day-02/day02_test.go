package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert.Equal(t, Map("A"), ROCK)
	assert.Equal(t, Map("X"), ROCK)
	assert.Equal(t, Map("B"), PAPER)
	assert.Equal(t, Map("Y"), PAPER)
	assert.Equal(t, Map("C"), SCISSORS)
	assert.Equal(t, Map("Z"), SCISSORS)
}

func TestScore(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		score int
	}{
		{name: "rock (A) loses to paper (Y)", input: "A Y", score: 8},
		{name: "paper (B) beats rock (X)", input: "B X", score: 1},
		{name: "scissors (C) ties scissors (Z)", input: "C Z", score: 6},
		{name: "rock (A) ties rock (X)", input: "A X", score: 4},
		{name: "scissors (C) beats paper (Y)", input: "C Y", score: 2},
	}

	for _, tc := range testCases {
		actual := Score(tc.input)
		assert.Equal(t, tc.score, actual, `(%v) wanted: %v got: %v`, tc.name, tc.score, actual)
	}
}

func TestCheat(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		score int
	}{
		{name: "opponent chooses rock, so player chooses rock to tie", input: "A Y", score: 4},
		{name: "opponent chooses paper, so player chooses rock to lose", input: "B X", score: 1},
		{name: "opponent chooses scissors, so player chooses rock to win", input: "C Z", score: 7},
	}

	for _, tc := range testCases {
		actual := AlternateScore(tc.input)
		assert.Equal(t, tc.score, actual, `(%v) wanted: %v got: %v`, tc.name, tc.score, actual)
	}
}
