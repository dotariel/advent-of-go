package main

import (
	"dotariel/util"
	"sort"
	"strings"
)

type Game struct {
	Hand
	Bid int
}

type Games []Game

func NewGame(s string) Game {
	parts := strings.Fields(s)

	return Game{
		Hand: NewHand(parts[0]),
		Bid:  util.ToInt(parts[1]),
	}
}

func (gs Games) Winnings() int {
	winnings := 0

	sort.Sort(gs)

	for i, game := range gs {
		rank := i + 1
		win := rank * game.Bid
		winnings += win
	}

	return winnings
}

func (gs Games) Less(i, j int) bool {
	return gs[i].Hand.Compare(gs[j].Hand) == -1
}

func (gs Games) Len() int {
	return len(gs)
}

func (gs Games) Swap(i, j int) {
	gs[i], gs[j] = gs[j], gs[i]
}
