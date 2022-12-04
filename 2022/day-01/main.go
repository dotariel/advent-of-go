package main

import (
	"dotariel/exercise"
	"dotariel/inputreader"
)

func main() {
	exercise.New(Part1, Part2).Run()
}

func Part1() int {
	input := inputreader.ReadAll("input.txt")
	elves := Parse(input)

	return TotalCalories(TopN(elves, 1))
}

func Part2() int {
	input := inputreader.ReadAll("input.txt")
	elves := Parse(input)

	return TotalCalories(TopN(elves, 3))
}
