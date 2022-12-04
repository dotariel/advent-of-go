package main

import (
	"unicode"

	"golang.org/x/exp/slices"
)

type ElfGroup []Rucksack

type Compartment struct {
	Items []rune
}

type Rucksack struct {
	Compartments []Compartment
}

func NewRucksack(s string) Rucksack {
	r := Rucksack{}

	c1 := s[0 : len(s)/2]
	c2 := s[len(s)/2:]

	r.Compartments = []Compartment{
		{Items: []rune(c1)},
		{Items: []rune(c2)},
	}

	return r
}

func (r Rucksack) AllItems() []rune {
	items := make([]rune, 0)

	for _, c := range r.Compartments {
		for _, i := range c.Items {
			items = append(items, i)
		}
	}

	return items
}

func (r Rucksack) FindDuplicates() []rune {
	dupes := make([]rune, 0)

	for _, i := range r.Compartments[0].Items {
		for _, j := range r.Compartments[1].Items {
			if i == j {
				if !slices.Contains(dupes, i) {
					dupes = append(dupes, i)
				}
			}
		}
	}

	return dupes
}

func (r Rucksack) ToString() string {
	return string(r.AllItems())
}

func (g ElfGroup) FindBadge() rune {
	for _, item := range g[0].AllItems() {
		matches := true
		for _, r := range g[1:] {
			matches = matches && slices.Contains(r.AllItems(), item)
		}

		if matches {
			return item
		}
	}

	return '0'
}

func GetPriority(r rune) int {
	mod := 38
	if unicode.IsLower(r) {
		mod = 96
	}

	return int(r) - mod
}
