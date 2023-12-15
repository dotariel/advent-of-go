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
	mp := NewMap(inputreader.ReadAll("input.txt"))

	return mp.GetTotalDistance(1)
}

func Part2() interface{} {
	mp := NewMap(inputreader.ReadAll("input.txt"))

	return mp.GetTotalDistance(1000000 - 1)
}
