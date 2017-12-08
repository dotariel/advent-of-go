package register

import (
	"regexp"
	"strconv"
	"strings"
)

// Registers is an alias for a map of string,int
type Registers map[string]int

// New constructs an empty Registers
func New() Registers {
	return Registers(make(map[string]int))
}

// Max returns the highest value in any of the registers
func (r Registers) Max() (max int) {
	values := make([]int, 0)

	for _, value := range r {
		values = append(values, value)
	}

	for i, value := range values {
		if i == 0 {
			max = value
		}

		if value > max {
			max = value
		}
	}

	return
}

// Process takes a single instruction and process it
func (r Registers) Process(instruction string) {
	exp := regexp.MustCompile("(?P<or>[a-z]+) (?P<op>[inc|dec]{3}) (?P<opx>-?[0-9]+) if (?P<cr>[a-z]+) (?P<cond>[!<>=]{1,2}) (?P<condx>-?[0-9]+)")
	match := exp.FindStringSubmatch(instruction)

	result := make(map[string]string)
	for i, name := range exp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	or := result["or"]
	cr := result["cr"]

	if _, ok := r[or]; !ok {
		r[or] = 0
	}

	if _, ok := r[cr]; !ok {
		r[cr] = 0
	}

	cFn := NewCondition(result["cond"])
	cX, _ := strconv.ParseInt(result["condx"], 10, 64)

	if condition := cFn(r[cr], int(cX)); condition {
		oFn := NewOperation(result["op"])
		oX, _ := strconv.ParseInt(result["opx"], 10, 64)

		r[or] = oFn(r[or], int(oX))
	}
}

// ProcessBatch processes several instructions from a
// line-separated list of strings
func (r Registers) ProcessBatch(input string) int {
	highest := 0

	for _, line := range strings.Split(input, "\n") {
		r.Process(line)
		if max := r.Max(); max > highest {
			highest = max
		}
	}

	return highest
}
