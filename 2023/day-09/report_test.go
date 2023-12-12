package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewReport(t *testing.T) {
	input := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

	expected := Report{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}

	assert.Equal(t, expected, NewReport(input))
}

func TestHistory_PredictNext(t *testing.T) {
	testCases := []struct {
		input    History
		expected int
	}{
		{input: History{0, 3, 6, 9, 12, 15}, expected: 18},
		{input: History{1, 3, 6, 10, 15, 21}, expected: 28},
		{input: History{10, 13, 16, 21, 30, 45}, expected: 68},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, tc.input.PredictNext())
	}
}

func TestHistory_PredictPrevious(t *testing.T) {
	testCases := []struct {
		input    History
		expected int
	}{
		{input: History{10, 13, 16, 21, 30, 45}, expected: 5},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, tc.input.PredictPrevious())
	}
}
