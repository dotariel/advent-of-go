package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath_Area(t *testing.T) {
	field := NewField(file("square.txt"))
	path := field.Walk()

	assert.Equal(t, 9, path.Area())
}

func TestPath_NumInteriorPoints(t *testing.T) {
	testCases := []struct {
		inputFile string
		expected  int
	}{
		{inputFile: "big.txt", expected: 4},
		{inputFile: "larger.txt", expected: 8},
		{inputFile: "largest.txt", expected: 10},
	}

	for _, tc := range testCases {
		field := NewField(file(tc.inputFile))
		path := field.Walk()
		interior := path.NumInteriorPoints()

		assert.Equal(t, tc.expected, interior)
	}
}
