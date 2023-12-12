package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCard_Compare(t *testing.T) {
	assert.Equal(t, 1, Card("3").Compare(Card("2")))
	assert.Equal(t, 0, Card("2").Compare(Card("2")))
	assert.Equal(t, -1, Card("2").Compare(Card("3")))
	assert.Equal(t, 1, Card("A").Compare(Card("K")))
	assert.Equal(t, -1, Card("J").Compare(Card("3")))
	assert.Equal(t, 1, Card("3").Compare(Card("J")))
}

func TestCardMap_HasN(t *testing.T) {
	assert.True(t, NewHand("JJJJJ").cardMap.HasFive())
	assert.True(t, NewHand("TTJJJ").cardMap.HasFive())
	assert.True(t, NewHand("TTTJJ").cardMap.HasFive())
}
