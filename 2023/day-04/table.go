package main

import (
	"errors"
)

type Cards []Card

func Accumulate(cards Cards) Cards {
	all := Cards{}

	for _, c := range cards {
		all = append(all, c.acc(&Cards{}, cards)...)
	}

	return all
}

func (c Card) acc(acc *Cards, base Cards) Cards {
	*acc = append(*acc, c)

	for _, copy := range base.GetSubWinners(c) {
		copy.acc(acc, base)
	}

	return *acc
}

func (t Cards) GetSubWinners(card Card) Cards {
	subs := make(Cards, 0)

	for i := range card.GetMatches() {
		if copy, err := t.FindById(card.id + i + 1); err == nil {
			subs = append(subs, copy)
		}
	}

	return subs
}

func (t Cards) FindById(id int) (Card, error) {
	for _, card := range t {
		if card.id == id {
			return card, nil
		}
	}

	return Card{}, errors.New("card not found")
}
