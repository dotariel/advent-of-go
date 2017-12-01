package captcha

import (
	"strconv"
)

// InverseCaptcha returns the sum of the list numbers represented by the
// digits based on the add strategy provided.
func InverseCaptcha(digits string, addfunc func([]int) int) int {
	return addfunc(convert(digits))
}

// WrapAroundAdd returns the sum of all the digits that match the next digit.
// The list is circular, so the digit after the last digit is the first
// digit in the list.
func WrapAroundAdd(numbers []int) (sum int) {
	return add(numbers, 1)
}

// HalfwayAroundAdd returns the sum of all the digits that match the digit
// halfway around the circular list.
func HalfwayAroundAdd(numbers []int) (sum int) {
	return add(numbers, len(numbers)/2)
}

func add(numbers []int, lookahead int) (sum int) {
	for i := 0; i < len(numbers); i++ {
		curr := numbers[i]
		next := numbers[(i+lookahead)%len(numbers)]
		if curr == next {
			sum += curr
		}
	}

	return
}

func convert(digits string) []int {
	numbers := make([]int, 0)

	for _, r := range digits {
		n, _ := strconv.ParseInt(string(r), 10, 64)
		numbers = append(numbers, int(n))
	}

	return numbers
}
