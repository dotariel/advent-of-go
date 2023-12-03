package main

import (
	"dotariel/inputreader"
	"fmt"
)

func main() {
	fmt.Printf("Part 1: %v\n", Part1())
	fmt.Printf("Part 2: %v\n", Part2())
}

func Part1() interface{} {
	total := 0
	bag := NewBag(map[string]int{"red": 12, "green": 13, "blue": 14})

	for _, input := range inputreader.ReadStrings("input.txt", "\n") {
		game := NewGame(input)
		if bag.Validate(game) {
			total += game.Id
		}
	}

	return total
}

func Part2() interface{} {
	total := 0

	for _, input := range inputreader.ReadStrings("input.txt", "\n") {
		game := NewGame(input)
		total += game.GetMinimumCubeSet().Power()
	}

	return total
}
