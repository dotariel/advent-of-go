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
	network := NewNetwork(inputreader.ReadAll("input.txt"))

	return network.Traverse()
}

func Part2() interface{} {
	network := NewNetwork(inputreader.ReadAll("input.txt"))

	return network.TraverseParallel()
}
