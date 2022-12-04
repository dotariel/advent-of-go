package main

import (
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	Food []Food
}

type Food struct {
	Calories int
}

type ByCalories []Elf

func Parse(s string) []Elf {
	elves := make([]Elf, 0)

	elfGroups := strings.Split(s, "\n\n")

	for _, group := range elfGroups {
		elf := Elf{Food: []Food{}}

		foods := strings.Split(group, "\n")
		for _, f := range foods {
			if calories, err := strconv.Atoi(f); err == nil {
				food := Food{Calories: calories}
				elf.Food = append(elf.Food, food)
			}
		}

		elves = append(elves, elf)
	}

	return elves
}

func (e Elf) TotalCalories() int {
	calories := 0

	for _, food := range e.Food {
		calories += food.Calories
	}

	return calories
}

func TopN(es []Elf, top int) []Elf {
	sort.Sort(sort.Reverse(ByCalories(es)))

	return es[0:top]
}

func (c ByCalories) Len() int           { return len(c) }
func (c ByCalories) Less(i, j int) bool { return c[i].TotalCalories() < c[j].TotalCalories() }
func (c ByCalories) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func TotalCalories(es []Elf) int {
	total := 0

	for _, elf := range es {
		total += elf.TotalCalories()
	}

	return total
}
