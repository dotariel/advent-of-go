package main

import (
	"fmt"
	"strconv"
)

var min = 264793
var max = 803935

func main() {
	fmt.Printf("Count: %v", Count(min, max))
}

func Count(min int, max int) int {
	matches := 0
	for i := min; i <= max; i++ {
		if Matches(i) {
			matches++
		}
	}
	return matches
}

func Matches(number int) bool {
	var lastValue int

	digits := make(map[rune]int)

	str := []rune(strconv.Itoa(number))
	for i := 0; i < len(str); i++ {
		digit := str[i]
		value, _ := strconv.Atoi(string(str[i]))

		if value < lastValue {
			return false
		}

		if _, exists := digits[digit]; !exists {
			digits[digit] = 0
		}

		digits[digit]++
		lastValue = value
	}

	pairs := 0

	for _, count := range digits {
		if count == 2 {
			pairs++
		}
	}

	return pairs > 0
}
