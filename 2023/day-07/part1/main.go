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
	games := Games{}

	for _, line := range inputreader.ReadStrings("input.txt", "\n") {
		games = append(games, NewGame(line))
	}

	return games.Winnings()
}

func Part2() interface{} {
	return 0
}
