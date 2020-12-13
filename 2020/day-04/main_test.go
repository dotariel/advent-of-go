package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassport(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", true},
		{"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929", false},
		{"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm", true},
		{"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in", false},
	}

	for _, tc := range testCases {
		assert.Equal(t, NewPassport(tc.input).IsValid(), tc.expected)
	}
}

func Test_byr(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"x", false},
		{"1", false},
		{"1919", false},
		{"2003", false},
		{"1920", true},
		{"2002", true},
	}

	for _, tc := range testCases {
		assert.Equal(t, byr(tc.input), tc.expected)
	}
}

func Test_iyr(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"x", false},
		{"1", false},
		{"2009", false},
		{"2021", false},
		{"2010", true},
		{"2020", true},
	}

	for _, tc := range testCases {
		assert.Equal(t, iyr(tc.input), tc.expected)
	}
}

func Test_eyr(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"x", false},
		{"1", false},
		{"2019", false},
		{"2031", false},
		{"2020", true},
		{"2030", true},
	}

	for _, tc := range testCases {
		assert.Equal(t, eyr(tc.input), tc.expected)
	}
}

func Test_hgt(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"x", false},
		{"1", false},
		{"149cm", false},
		{"194cm", false},
		{"150cm", true},
		{"193cm", true},
		{"58in", false},
		{"77in", false},
		{"59in", true},
		{"76in", true},
	}

	for _, tc := range testCases {
		assert.Equal(t, hgt(tc.input), tc.expected)
	}
}

func Test_hcl(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"x", false},
		{"#", false},
		{"#fffff", false},
		{"#fffffff", false},
		{"#abcxyz", false},
		{"#abcdef", true},
	}

	for _, tc := range testCases {
		assert.Equal(t, hcl(tc.input), tc.expected)
	}
}

func Test_ecl(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"x", false},
		{"amb", true},
		{"blu", true},
		{"brn", true},
		{"gry", true},
		{"grn", true},
		{"hzl", true},
		{"oth", true},
	}

	for _, tc := range testCases {
		assert.Equal(t, ecl(tc.input), tc.expected)
	}
}

func Test_pid(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"x", false},
		{"1111111", false},
		{"0000000001", false},
		{"aaaaaaaaa", false},
		{"000000001", true},
	}

	for _, tc := range testCases {
		assert.Equal(t, pid(tc.input), tc.expected)
	}
}

func Test_cid(t *testing.T) {
	assert.True(t, cid("s"))
}

func Test_getField(t *testing.T) {
	m := map[string]string{
		"foo": "bar",
	}

	assert.Equal(t, getField(m, "foo"), "bar")
	assert.Equal(t, getField(m, "baz"), "")
}
