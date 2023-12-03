package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalibration_GetValue(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: "1abc2", expected: 12},
		{input: "pqr3stu8vwx", expected: 38},
		{input: "a1b2c3d4e5f", expected: 15},
		{input: "treb7uchet", expected: 77},
	}

	for _, tc := range testCases {
		c := NewCalibration(tc.input)

		assert.Equal(t, tc.expected, c.GetValue())
	}
}

func TestCalibration_GetExtendedValue(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{

		{input: "two1nine", expected: 29},
		{input: "eightwothree", expected: 83},
		{input: "abcone2threexyz", expected: 13},
		{input: "xtwone3four", expected: 24},
		{input: "4nineeightseven2", expected: 42},
		{input: "zoneight234", expected: 14},
		{input: "7pqrstsixteen", expected: 76},
		{input: "3bs", expected: 33},
		{input: "14nsnlqqlgfourxbzzxfztvbxsnxttjmktcxkkkzfphppsczqoneightgc", expected: 18},
	}

	for _, tc := range testCases {
		c := NewCalibration(tc.input)

		assert.Equal(t, tc.expected, c.GetExtendedValue())
	}
}

func Test(t *testing.T) {

	s := "12345"

	assert.Equal(t, "45", strings.TrimPrefix(s, "123"))
	assert.Equal(t, "2345", s[1:])
}
