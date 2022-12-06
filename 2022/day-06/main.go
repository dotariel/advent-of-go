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
	return FindMarker(inputreader.ReadAll("input.txt"))
}

func Part2() interface{} {
	return FindMessage(inputreader.ReadAll("input.txt"))
}
