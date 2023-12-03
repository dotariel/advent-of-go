package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var digits = map[string]int{
	"1":     1,
	"one":   1,
	"2":     2,
	"two":   2,
	"3":     3,
	"three": 3,
	"4":     4,
	"four":  4,
	"5":     5,
	"five":  5,
	"6":     6,
	"six":   6,
	"7":     7,
	"seven": 7,
	"8":     8,
	"eight": 8,
	"9":     9,
	"nine":  9,
}

type Calibration struct {
	value string
}

func NewCalibration(s string) Calibration {
	return Calibration{value: s}
}

func (c Calibration) GetValue() int {

	values := make([]rune, 0)

	for _, rune := range c.value {
		if unicode.IsDigit(rune) {
			values = append(values, rune)
		}
	}

	val, err := strconv.Atoi(string(values[0]) + string(values[len(values)-1]))
	if err != nil {
		return 0
	}

	return val
}

func (c Calibration) GetExtendedValue() int {
	s := strings.Clone(c.value)
	values := make([]rune, 0)

	for {
		if len(s) == 0 {
			break
		}

		for key, value := range digits {
			if strings.HasPrefix(s, key) {
				values = append(values, rune(value))
			}
		}

		s = s[1:]
	}

	digits := fmt.Sprintf("%d%d", values[0], values[len(values)-1])

	val, err := strconv.Atoi(digits)
	if err != nil {
		return 0
	}

	return val
}
