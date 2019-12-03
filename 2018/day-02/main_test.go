package main

import "testing"

func TestHasCount(t *testing.T) {
	testCases := []struct {
		id       string
		hasTwo   bool
		hasThree bool
	}{
		{"abcdef", false, false},
		{"bababc", true, true},
		{"abbcde", true, false},
		{"abcccd", false, true},
		{"aabcdd", true, false},
		{"abcdee", true, false},
		{"ababab", false, true},
	}
	for _, tc := range testCases {
		hasTwo := HasCount(tc.id, 2)
		hasThree := HasCount(tc.id, 3)

		if hasTwo != tc.hasTwo {
			t.Errorf("Id %s expected to have 2", tc.id)
		}
		if hasThree != tc.hasThree {
			t.Errorf("Id %s expected to have 3", tc.id)
		}
	}
}

func TestChecksum(t *testing.T) {
	input := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	expectedChecksum := 12

	if checksum := Checksum(input); expectedChecksum != checksum {
		t.Errorf("Checksum did not match; expected %v, but got %v", expectedChecksum, checksum)
	}
}

func TestGetUnique(t *testing.T) {
	input := []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	expected := "fgij"

	if unique := Unique(input); unique != expected {
		t.Errorf("Expected unique value to be %s, but got %s", expected, unique)
	}
}
