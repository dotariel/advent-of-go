package main

import (
	"dotariel/exercise"
	"dotariel/inputreader"
)

func main() {
	exercise.New(Part1, Part2).Run()
}

func Part1() int {
	entries := inputreader.ReadStrings("input.txt", "\n")
	total := 0

	for _, entry := range entries {
		total += Score(entry)
	}

	return total
}

func Part2() int {
	entries := inputreader.ReadStrings("input.txt", "\n")
	total := 0

	for _, entry := range entries {
		total += AlternateScore(entry)
	}

	return total
}
