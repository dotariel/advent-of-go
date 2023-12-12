package main

import "strconv"

type Card string

type CardMap map[string][]Card

type Cards []Card

func (c Card) Compare(other Card) int {
	if c.Value() < other.Value() {
		return -1
	}

	if c.Value() > other.Value() {
		return 1
	}

	return 0
}

func (c Card) Value() int {
	if val, err := strconv.Atoi(string(c)); err == nil {
		return val
	}

	if c == "T" {
		return 10
	}

	if c == "J" {
		return 11
	}

	if c == "Q" {
		return 12
	}

	if c == "K" {
		return 13
	}

	if c == "A" {
		return 14
	}

	return 0
}

func (cs Cards) Less(i, j int) bool {
	return cs[i].Value() < cs[j].Value()
}

func (cs Cards) Len() int {
	return len(cs)
}

func (cs Cards) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func (cs Cards) Values() []int {
	values := make([]int, 0)

	for _, c := range cs {
		values = append(values, c.Value())
	}

	return values
}

func (cm CardMap) HasPair() bool {
	return cm.hasN(2)
}

func (cm CardMap) HasTwoPair() bool {
	return len(cm) == 3 && cm.HasPair()
}

func (cm CardMap) HasSet() bool {
	return cm.hasN(3)
}

func (cm CardMap) HasQuads() bool {
	return cm.hasN(4)
}

func (cm CardMap) HasFive() bool {
	return cm.hasN(5)
}

func (cm CardMap) hasN(n int) bool {
	for _, v := range cm {
		if len(v) == n {
			return true
		}
	}

	return false
}
