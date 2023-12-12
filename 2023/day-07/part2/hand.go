package main

import (
	"strings"
)

const (
	RANK_HIGH_CARD = iota
	RANK_PAIR
	RANK_TWO_PAIR
	RANK_THREE_OF_A_KIND
	RANK_FULL_HOUSE
	RANK_FOUR_OF_A_KIND
	RANK_FIVE_OF_KIND
)

type Hand struct {
	Cards   Cards
	cardMap CardMap
}

type Hands []Hand

func NewHand(input string) Hand {
	hand := Hand{Cards: Cards{}, cardMap: CardMap{}}

	for _, card := range strings.Split(input, "") {
		hand.Cards = append(hand.Cards, Card(card))

		if _, ok := hand.cardMap[card]; !ok {
			hand.cardMap[card] = make([]Card, 0)
		}

		hand.cardMap[card] = append(hand.cardMap[card], Card(card))
	}

	return hand
}

func (h Hand) Rank() int {
	if h.cardMap.HasFive() {
		return RANK_FIVE_OF_KIND
	}

	if h.cardMap.HasQuads() {
		return RANK_FOUR_OF_A_KIND
	}

	if h.cardMap.HasSetAndPair() {
		return RANK_FULL_HOUSE
	}

	if h.cardMap.HasSet() {
		return RANK_THREE_OF_A_KIND
	}

	if h.cardMap.HasTwoPair() {
		return RANK_TWO_PAIR
	}

	if h.cardMap.HasPair() {
		return RANK_PAIR
	}

	return RANK_HIGH_CARD
}

func (h Hand) Compare(other Hand) int {
	if h.Rank() > other.Rank() {
		return 1
	}

	if h.Rank() < other.Rank() {
		return -1
	}

	for i := range h.Cards {
		if h.Cards[i].Compare(other.Cards[i]) == 0 {
			continue
		}

		return h.Cards[i].Compare(other.Cards[i])
	}

	return 0
}

func (hs Hands) Less(i, j int) bool {
	return hs[i].Compare(hs[j]) == -1
}

func (hs Hands) Len() int {
	return len(hs)
}

func (hs Hands) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}
