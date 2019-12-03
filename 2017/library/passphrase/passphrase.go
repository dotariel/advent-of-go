package passphrase

import (
	"sort"
	"strings"
)

type ValidFunc func(string) bool

func CountValid(input string, valid ValidFunc) (count int) {
	for _, phrase := range strings.Split(input, "\n") {
		if valid(phrase) {
			count++
		}
	}

	return
}

func Unique(phrase string) bool {
	seen := make(map[string]int)
	for _, word := range strings.Fields(phrase) {
		if _, exists := seen[word]; exists {
			return false
		}
		seen[word]++
	}
	return true
}

func NonAnagram(phrase string) bool {
	words := strings.Fields(phrase)

	for i, word := range words {
		for j, other := range words {
			if isAnagram(word, other) && i != j {
				return false
			}
		}
	}

	return true
}

type byRunes []rune

func (s byRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s byRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byRunes) Len() int {
	return len(s)
}

func sorted(s string) string {
	runes := []rune(s)
	sort.Sort(byRunes(runes))
	return string(runes)
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	this := strings.ToLower(a)
	that := strings.ToLower(b)

	return sorted(this) == sorted(that)
}
