package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGames_Sort(t *testing.T) {
	sorted := Games{
		NewGame("32T3K 765"),
		NewGame("T55J5 684"),
		NewGame("23456 12"),
		NewGame("KK677 28"),
		NewGame("JJJJJ 28"),
		NewGame("KTJJT 220"),
		NewGame("QQQJA 483"),
	}

	expected := Games{
		NewGame("23456 12"),
		NewGame("32T3K 765"),
		NewGame("KTJJT 220"),
		NewGame("KK677 28"),
		NewGame("T55J5 684"),
		NewGame("QQQJA 483"),
		NewGame("JJJJJ 28"),
	}

	sort.Sort(sorted)

	assert.Equal(t, expected, sorted)
}

func TestGames_Winnings(t *testing.T) {
	sorted := Games{
		NewGame("32T3K 765"),
		NewGame("T55J5 684"),
		NewGame("KK677 28"),
		NewGame("KTJJT 220"),
		NewGame("QQQJA 483"),
	}

	assert.Equal(t, 6440, sorted.Winnings())
}
