package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHand_Joker(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "all jokers", input: "JJJJJ", expected: RANK_FIVE_OF_KIND},
		{name: "four of a kind with 1 joker", input: "TTTTJ", expected: RANK_FIVE_OF_KIND},
		{name: "three of a kind with 2 jokers", input: "TTTJJ", expected: RANK_FIVE_OF_KIND},
		{name: "three of a kind with 1 joker", input: "TTTJ3", expected: RANK_FOUR_OF_A_KIND},
		{name: "two pair with 1 joker", input: "TT33J", expected: RANK_FULL_HOUSE},
		{name: "pair with 3 jokers", input: "TTJJJ", expected: RANK_FIVE_OF_KIND},
		{name: "pair with 2 jokers", input: "TTJJ3", expected: RANK_FOUR_OF_A_KIND},
		{name: "pair with 1 joker", input: "TTJ93", expected: RANK_THREE_OF_A_KIND},
		{name: "pair of jokers", input: "JJ293", expected: RANK_THREE_OF_A_KIND},
		{name: "high card with 1 joker", input: "2345J", expected: RANK_PAIR},
		{name: "high card with 0 jokers", input: "32T9K", expected: RANK_HIGH_CARD},
	}

	for _, tc := range testCases {
		actual := NewHand(tc.input).Rank()
		assert.Equal(t, tc.expected, actual, "scenario '%v' failed; wanted:%v got %v\n", tc.name, tc.expected, actual)
	}
}

func TestHand_Compare(t *testing.T) {
	h1 := NewHand("JKKK2")
	h2 := NewHand("QQQQ2")
	assert.Equal(t, h1.Compare(h2), -1)
	assert.Equal(t, h1.Rank(), RANK_FOUR_OF_A_KIND)
	assert.Equal(t, h2.Rank(), RANK_FOUR_OF_A_KIND)

	h3 := NewHand("22T39")
	h4 := NewHand("J2T39")
	assert.Equal(t, h3.Compare(h4), 1)

	h5 := NewHand("JJ293")
	assert.Equal(t, RANK_THREE_OF_A_KIND, h5.Rank())
}
