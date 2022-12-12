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
	return GetIncreases(inputreader.ReadInts("input.txt", "\n"), 1)
}

func Part2() interface{} {
	return GetIncreases(inputreader.ReadInts("input.txt", "\n"), 3)
}
