package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	input := []string{
		".#.",
		".#.",
		".#.",
		".#.",
	}

	grid := newGrid()
	grid.load(input)

	var slope Slope = Slope{1, 1}

	char, done := grid.move(slope)
	assert.Equal(t, char, "#")
	assert.False(t, done)

	char, done = grid.move(slope)
	assert.Equal(t, char, ".")
	assert.False(t, done)

	char, done = grid.move(slope)
	assert.Equal(t, char, ".")
	assert.True(t, done)
}
