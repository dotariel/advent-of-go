package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	input   int
	matches bool
}{
	{111111, false},
	{223450, false},
	{123789, false},
	{123345, true},
	{523345, false},
	{123444, false},
	{112233, true},
	{111122, true},
}

func TestMatches(t *testing.T) {
	for _, tc := range testCases {
		assert.Equal(t, tc.matches, Matches(tc.input))
	}
}

func TestCount(t *testing.T) {
	t.Logf("Count: %v", Count(264793, 803935))
}
