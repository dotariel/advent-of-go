package main

import (
	"dotariel/inputreader"
)

func main() {
	inputs := inputreader.ReadStrings("input.txt", "\n")
	checksum := Checksum(inputs)
	unique := Unique(inputs)

	println(checksum)
	println(unique)
}

func Checksum(ids []string) int {
	hasTwo := 0
	hasThree := 0

	for _, id := range ids {
		if HasCount(id, 2) {
			hasTwo++
		}
		if HasCount(id, 3) {
			hasThree++
		}
	}

	return hasTwo * hasThree
}

func HasCount(id string, count int) bool {
	lexicon := parse(id)

	for _, v := range lexicon {
		if v == count {
			return true
		}
	}

	return false
}

func Unique(ids []string) string {
	matches := make([]rune, 0)

	for _, id := range ids {
		for _, other := range ids {
			matches = make([]rune, 0)

			if id == other {
				break
			}

			a := []rune(id)
			b := []rune(other)

			for pos := range a {
				if a[pos] == b[pos] {
					matches = append(matches, a[pos])
				}
			}

			if len(matches) > len(id)-2 {
				return string(matches)
			}
		}
	}

	return ""
}

func parse(s string) map[rune]int {
	lexicon := make(map[rune]int)

	for _, char := range s {
		if _, ok := lexicon[char]; !ok {
			lexicon[char] = 0
		}

		lexicon[char]++
	}

	return lexicon
}
