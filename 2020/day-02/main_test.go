package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolicy_IsValid(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{input: "1-3 a: abcde", expected: true},
		{input: "1-3 b: cdefg", expected: false},
		{input: "2-9 c: ccccccccc", expected: true},
		{input: "11-13 d: dddddddddblddmddk", expected: true},
	}

	for _, tc := range testCases {
		policy := parsePolicy((tc.input))
		password := parsePassword((tc.input))

		assert.Equal(t, tc.expected, policy.IsValid(password))
	}
}

func TestPolicy_IsPositionalValid(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{input: "1-3 a: abcde", expected: true},
		{input: "1-3 b: cdefg", expected: false},
		{input: "2-9 c: ccccccccc", expected: false},
		{input: "11-13 d: dddddddddblddmddk", expected: true},
		{input: "6-12 r: crxrrzrnprrr", expected: true},
	}

	for _, tc := range testCases {
		policy := parsePolicy((tc.input))
		password := parsePassword((tc.input))

		assert.Equal(t, tc.expected, policy.IsPositionalValid(password))
	}
}
