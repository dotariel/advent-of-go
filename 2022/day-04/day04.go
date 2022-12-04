package main

import (
	"strconv"
	"strings"
)

type RangeDef struct {
	Lower int
	Upper int
}

type Range []int

func Parse(s string) []RangeDef {
	rds := make([]RangeDef, 0)

	for _, r := range strings.Split(s, ",") {
		rds = append(rds, NewRangeDef(r))
	}

	return rds
}

func NewRangeDef(s string) RangeDef {
	bounds := strings.Split(s, "-")

	lower, _ := strconv.Atoi(bounds[0])
	upper, _ := strconv.Atoi(bounds[1])

	return RangeDef{Lower: lower, Upper: upper}
}

func (rd RangeDef) Contains(other RangeDef) bool {
	return other.Lower >= rd.Lower && other.Upper <= rd.Upper
}

func (rd RangeDef) ToRange() Range {
	r := Range(make([]int, 0))

	for i := rd.Lower; i <= rd.Upper; i++ {
		r = append(r, i)
	}

	return r
}

func (r Range) Overlaps(other Range) bool {
	for _, i := range other {
		for _, j := range r {
			if i == j {
				return true
			}
		}
	}

	return false
}
