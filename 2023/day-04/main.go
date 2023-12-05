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

	for _, input := range inputreader.ReadStrings("input.txt", "\n") {
		total += NewCard(input).Points()
	}

	return total
}

func Part2() interface{} {
	table := Cards{}

	for _, input := range inputreader.ReadStrings("input.txt", "\n") {
		table = append(table, NewCard(input))
	}

	return len(Accumulate(table))
}
