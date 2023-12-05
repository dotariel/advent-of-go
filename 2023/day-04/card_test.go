package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCard(t *testing.T) {
	input := "Card     1: 41 48 83 86 17 | 83 86 6 31 17  9 48 53"

	card := NewCard(input)

	assert.IsType(t, card, Card{})
	assert.Equal(t, 1, card.id)
	assert.Equal(t, []int{41, 48, 83, 86, 17}, card.numbers)
	assert.Equal(t, []int{83, 86, 6, 31, 17, 9, 48, 53}, card.winners)
}

func TestCard_GetMatches(t *testing.T) {
	for _, tc := range []struct {
		card    string
		id      int
		matches []int
		points  int
	}{
		{card: "Card   1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", id: 1, matches: []int{48, 83, 17, 86}, points: 8},
		{card: "Card   2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", id: 2, matches: []int{32, 61}, points: 2},
		{card: "Card   3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", id: 3, matches: []int{1, 21}, points: 2},
		{card: "Card   4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", id: 4, matches: []int{84}, points: 1},
		{card: "Card   5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", id: 5, matches: []int{}, points: 0},
		{card: "Card   6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", id: 6, matches: []int{}, points: 0},
		{card: "Card  99:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", id: 99, matches: []int{1, 21}, points: 2},
		{card: "Card 100:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", id: 100, matches: []int{1, 21}, points: 2},
	} {
		c := NewCard(tc.card)
		assert.Equal(t, tc.id, c.id)
		assert.ElementsMatch(t, tc.matches, c.GetMatches())
		assert.Equal(t, tc.points, c.Points())
	}
}
