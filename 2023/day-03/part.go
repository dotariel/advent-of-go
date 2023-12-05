package main

import "strconv"

type Part struct {
	string
	Range
}

func (p Part) Value() int {
	if value, err := strconv.Atoi(p.string); err == nil {
		return value
	}

	return 0
}
