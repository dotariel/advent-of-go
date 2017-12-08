package register

import (
	"regexp"
	"strconv"
	"strings"
)

type operation func(int, int) int
type condition func(int, int) bool

type Registers map[string]int

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

func inc(n, x int) int {
	return n + x
}

func dec(n, x int) int {
	return n - x
}

func noop(n, x int) int {
	return n
}

func eq(a, b int) bool {
	return a == b
}

func neq(a, b int) bool {
	return !eq(a, b)
}

func gt(a, b int) bool {
	return a > b
}

func gte(a, b int) bool {
	return a >= b
}

func lt(a, b int) bool {
	return a < b
}

func lte(a, b int) bool {
	return a <= b
}

func New() Registers {
	return Registers(make(map[string]int))
}

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

	cFn := conditionFromString(result["cond"])
	cX, _ := strconv.ParseInt(result["condx"], 10, 64)

	if condition := cFn(r[cr], int(cX)); condition {
		oFn := operationFromString(result["op"])
		oX, _ := strconv.ParseInt(result["opx"], 10, 64)

		r[or] = oFn(r[or], int(oX))
	}
}

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

func operationFromString(op string) operation {
	switch op {
	case "inc":
		return inc
	case "dec":
		return dec
	}

	return noop
}

func conditionFromString(c string) condition {
	if c == "<" {
		return lt
	}

	if c == "<=" {
		return lte
	}

	if c == ">" {
		return gt
	}

	if c == ">=" {
		return gte
	}

	if c == "!=" {
		return neq
	}

	return eq
}
