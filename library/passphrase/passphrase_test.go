package passphrase

import "testing"

func TestUnique(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}

	for _, tt := range testCases {
		if actual := Unique(tt.input); actual != tt.expected {
			t.Errorf("unique test failed for '%v'; wanted %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func TestNonAnagram(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"aa bb cc dd aa", false},
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
		{"aa bb cc dd ee", true},
	}

	for _, tt := range testCases {
		if actual := NonAnagram(tt.input); actual != tt.expected {
			t.Errorf("anagram test failed for '%v'; wanted %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func TestCountValid(t *testing.T) {
	testCases := []struct {
		input    string
		fn       ValidFunc
		expected int
	}{
		{"aa bb cc dd ee\naa bb cc dd aa\naa bb cc dd aaa", Unique, 2},
		{"aa bb cc dd ee\naa bb cc dd aa\naa bb cc dd aaa", NonAnagram, 2},
		{"abcde fghij\nabcde xyz ecdab\na ab abc abd abf abj\niiii oiii ooii oooi oooo\noiii ioii iioi iiio", NonAnagram, 3},
	}

	for _, tt := range testCases {
		if actual := CountValid(tt.input, tt.fn); actual != tt.expected {
			t.Errorf("validation count test failed for '%v'; wanted %v but got %v", tt.input, tt.expected, actual)
		}
	}
}
