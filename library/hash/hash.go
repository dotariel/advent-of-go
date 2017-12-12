package hash

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Hash simulates tying a knot in a circle of string with 256 marks on it.
// Based on the input to be hashed, the function repeatedly selects a span of string,
// brings the ends together, and gives the span a half-twist to reverse the order of
// the marks within it. After doing this many times, the order of the marks is used
// to build the resulting hash.
func Hash(list, lengths []uint8) []uint8 {
	ret := append([]uint8{}, list...)
	curPos := 0
	skip := 0

	for _, l := range lengths {
		t, _ := twist(ret, curPos, uint8(l))
		ret = t
		curPos = (curPos + int(l) + skip) % len(list)
		skip++
	}

	return ret
}

// Parse parses a comma-separated string of uint8 values into a valid input for Hash
func Parse(input string) (list []uint8) {
	if len(input) == 0 {
		return
	}

	for _, item := range strings.Split(input, ",") {
		n, _ := strconv.ParseInt(item, 10, 16)
		list = append(list, uint8(n))
	}
	return
}

func twist(n []uint8, p int, l uint8) ([]uint8, error) {
	if int(l) > len(n) {
		return nil, fmt.Errorf("invalid length")
	}

	if p >= len(n) || p < 0 {
		return nil, fmt.Errorf("invalid position")
	}

	ret := append([]uint8{}, n...)
	sublist, positions, _ := sublist(n, p, l)
	reverse := reverse(sublist)

	for i, v := range positions {
		ret[v] = reverse[i]
	}

	return ret, nil
}

func sublist(n []uint8, p int, l uint8) ([]uint8, []int, error) {
	if int(l) > len(n) {
		return nil, nil, fmt.Errorf("invalid length")
	}

	if p >= len(n) || p < 0 {
		return nil, nil, fmt.Errorf("invalid position")
	}

	s := make([]uint8, 0)
	pos := make([]int, 0)

	have := len(n) - p
	next := p + int(math.Min(float64(len(n)-p), float64(l)))

	for i, it := range n[p:next] {
		s = append(s, it)
		pos = append(pos, p+i)
	}

	if int(l) > have {
		for i, it := range n[0:(int(l) - have)] {
			s = append(s, it)
			pos = append(pos, i)
		}
	}

	return s, pos, nil
}

func reverse(n []uint8) []uint8 {
	var rev []uint8

	for _, v := range n {
		rev = append(rev, v)
	}

	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}

	return rev
}
