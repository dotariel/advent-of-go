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
	input := inputreader.ReadAll("input.txt")

	grid := Parse(input)

	return grid.CountVisible()
}

func Part2() interface{} {
	return 0
}
