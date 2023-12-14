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
	field := NewField(inputreader.ReadAll("input.txt"))
	path := field.Walk()

	return path.Max()
}

func Part2() interface{} {
	field := NewField(inputreader.ReadAll("input.txt"))
	path := field.Walk()

	return path.NumInteriorPoints()
}
