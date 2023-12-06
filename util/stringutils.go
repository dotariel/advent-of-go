package util

import "strconv"

func ToInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}

	return 0
}

func ToInts(ss []string) []int {
	ints := make([]int, 0)

	for _, s := range ss {
		ints = append(ints, ToInt(s))
	}

	return ints
}
