package main

import (
	"dotariel/inputreader"
	"fmt"
	"slices"
)

func main() {
	fmt.Printf("Part 1: %v\n", Part1())
	fmt.Printf("Part 2: %v\n", Part2())
}

func Part1() interface{} {
	almanac := NewAlmanac(inputreader.ReadAll("input.txt"))

	locations := make([]int, 0)

	for _, seed := range almanac.Seeds {
		locations = append(locations, almanac.Traverse(seed))
	}

	return slices.Min(locations)
}

func Part2() interface{} {
	return 0
}
