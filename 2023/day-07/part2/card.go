package main

import (
	"slices"
	"strconv"
)

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

func (c Card) IsJoker() bool {
	return c == "J"
}

func (c Card) Value() int {
	if val, err := strconv.Atoi(string(c)); err == nil {
		return val
	}

	if c == "T" {
		return 10
	}

	if c.IsJoker() {
		return 1
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

func (cs Cards) Exists(c Card) bool {
	for _, card := range cs {
		if card == c {
			return true
		}
	}

	return false
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

func (cm CardMap) HasSetAndPair() bool {
	return cm.hasN(3, 2)
}

func (cm CardMap) HasQuads() bool {
	return cm.hasN(4)
}

func (cm CardMap) HasFive() bool {
	return cm.hasN(5) || len(cm) == 1
}

func (cm CardMap) hasN(ns ...int) bool {
	cardmap := cm.clone()

	jokers := 0
	for k, v := range cardmap {
		if Card(k).IsJoker() {
			jokers = len(v)
			delete(cardmap, k)
		}
	}

	haslist := make([]int, 0)

	for _, targetCount := range ns {
		for c, cards := range cardmap {

			if len(cards)+jokers >= targetCount {
				haslist = append(haslist, targetCount)

				jokers -= targetCount - len(cards)

				delete(cardmap, c)
			}
		}
	}

	has := true

	for _, target := range ns {
		has = has && slices.Contains(haslist, target)
	}

	return has

}

func (cm CardMap) numJokers() int {
	for k, v := range cm {
		if Card(k).IsJoker() {
			return len(v)
		}
	}

	return 0
}

func (cm CardMap) clone() CardMap {
	copy := CardMap{}

	for k, v := range cm {
		if _, exists := copy[k]; !exists {
			copy[k] = []Card{}
		}

		copy[k] = append(copy[k], v...)
	}

	return copy
}
