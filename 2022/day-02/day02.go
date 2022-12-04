package main

import (
	"strings"
)

type Rock struct{}
type Paper struct{}
type Scissors struct{}

type Option interface {
	Defeats() Option
	DefeatedBy() Option
	Value() int
}

var (
	ROCK     = Rock{}
	PAPER    = Paper{}
	SCISSORS = Scissors{}
)

var (
	C_WIN  = 6
	C_LOSS = 0
	C_TIE  = 3
)

func (r Rock) Defeats() Option {
	return SCISSORS
}

func (r Rock) DefeatedBy() Option {
	return PAPER
}

func (r Rock) Value() int {
	return 1
}

func (p Paper) Defeats() Option {
	return ROCK
}

func (p Paper) DefeatedBy() Option {
	return SCISSORS
}

func (p Paper) Value() int {
	return 2
}

func (s Scissors) Defeats() Option {
	return PAPER
}

func (s Scissors) DefeatedBy() Option {
	return ROCK
}

func (s Scissors) Value() int {
	return 3
}

func AlternateScore(s string) int {
	options := strings.Fields(s)

	opponent := Map(options[0])
	player := GetResponse(opponent, options[1])

	return Compare(player, opponent)
}

func Score(s string) int {
	options := strings.Fields(s)

	return Compare(Map(options[1]), Map(options[0]))
}

func Compare(player Option, opponent Option) int {
	if player.Defeats() == opponent {
		return player.Value() + C_WIN
	}

	if opponent.Defeats() == player {
		return player.Value() + C_LOSS
	}

	return player.Value() + C_TIE
}

func GetResponse(o Option, s string) Option {
	if s == "X" {
		return o.Defeats()
	}

	if s == "Y" {
		return o
	}

	if s == "Z" {
		return o.DefeatedBy()
	}

	return nil
}

func Map(s string) Option {
	if s == "A" || s == "X" {
		return ROCK
	}

	if s == "B" || s == "Y" {
		return PAPER
	}

	if s == "C" || s == "Z" {
		return SCISSORS
	}

	return nil
}
