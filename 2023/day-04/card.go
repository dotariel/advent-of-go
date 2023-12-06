package main

import (
	"dotariel/util"
	"regexp"
	"strings"
)

type Card struct {
	id      int
	numbers []int
	winners []int
}

func NewCard(s string) Card {
	r := regexp.MustCompile(`(?:Card)(?:[\s]+)([\d]+): ([\d\s]+)+ \| (.*)`)
	matches := r.FindStringSubmatch(s)

	return Card{
		id:      util.ToInt(matches[1]),
		numbers: util.ToInts(strings.Fields(matches[2])),
		winners: util.ToInts(strings.Fields(matches[3])),
	}
}

func (c Card) GetMatches() []int {
	matches := []int{}

	for _, number := range c.numbers {
		for _, winner := range c.winners {
			if number == winner {
				matches = append(matches, number)
			}
		}
	}

	return matches
}

func (c Card) Points() int {
	points := 0
	matches := c.GetMatches()

	if len(matches) == 0 {
		return 0
	}

	for i := range matches {
		if i == 0 {
			points += 1
			continue
		}

		points *= 2
	}

	return points
}

func (c Card) IsWinner() bool {
	return len(c.GetMatches()) > 0
}
