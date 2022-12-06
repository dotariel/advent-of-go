package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMarker(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", expected: 7},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", expected: 5},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", expected: 6},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", expected: 10},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", expected: 11},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, FindMarker(tc.input))
	}
}

func TestFindMesasge(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", expected: 19},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", expected: 23},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", expected: 23},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", expected: 29},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", expected: 26},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, FindMessage(tc.input))
	}
}
