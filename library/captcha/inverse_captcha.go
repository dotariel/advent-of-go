package captcha

import "strconv"

// AddFunc defines a type of function that adds the values of digits that
// meet a specific criteria.
type AddFunc func(string) int

// InverseCaptcha returns the sum of the list numbers represented by the
// digits based on the add strategy provided.
func InverseCaptcha(digits string, f AddFunc) int {
	return f(digits)
}

// WrapAroundAdd returns the sum of all the digits that match the next digit.
// The list is circular, so the digit after the last digit is the first
// digit in the list.
func WrapAroundAdd(digits string) (sum int) {
	return add(digits, 1)
}

// HalfwayAroundAdd returns the sum of all the digits that match the digit
// halfway around the circular list.
func HalfwayAroundAdd(digits string) (sum int) {
	return add(digits, len(digits)/2)
}

func add(digits string, lookahead int) (sum int) {
	for i := 0; i < len(digits); i++ {
		curr := digits[i]
		next := digits[(i+lookahead)%len(digits)]
		if curr == next {
			n, _ := strconv.ParseInt(string(curr), 10, 64)
			sum += int(n)
		}
	}

	return
}
