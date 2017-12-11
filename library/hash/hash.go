package hash

import (
	"fmt"
	"math"
)

func sublist(n []uint8, p int, l int) ([]uint8, []int, error) {
	if l > len(n) {
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

	if l > have {
		for i, it := range n[0:(l - have)] {
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

func do(n []uint8, p int, l int) ([]uint8, error) {
	if l > len(n) {
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
