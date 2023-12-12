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
	assert.Equal(t, -1, Card("T").Compare(Card("J")))
	assert.Equal(t, -1, Card("10").Compare(Card("J")))
}
