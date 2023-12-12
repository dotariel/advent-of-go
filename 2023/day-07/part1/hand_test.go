package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHand_TieBreaker(t *testing.T) {
	h1 := NewHand("33332")
	h2 := NewHand("2AAAA")
	h3 := NewHand("77888")
	h4 := NewHand("77788")
	h5 := NewHand("KK677")
	h6 := NewHand("KTJJT")

	assert.Equal(t, RANK_FOUR_OF_A_KIND, h1.Rank())
	assert.Equal(t, RANK_FOUR_OF_A_KIND, h2.Rank())
	assert.Equal(t, RANK_FULL_HOUSE, h3.Rank())
	assert.Equal(t, RANK_FULL_HOUSE, h4.Rank())
	assert.Equal(t, RANK_TWO_PAIR, h5.Rank())
	assert.Equal(t, RANK_TWO_PAIR, h6.Rank())

	assert.Equal(t, 1, h1.Compare(h2))
	assert.Equal(t, 1, h3.Compare(h4))
	assert.Equal(t, 1, h5.Compare(h6))
}
