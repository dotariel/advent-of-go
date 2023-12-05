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

	input := inputreader.ReadAll("input.txt")
	schematic := NewSchematic(input)

	for _, part := range schematic.FindParts() {
		total += part.Value()
	}

	return total
}

func Part2() interface{} {
	total := 0

	input := inputreader.ReadAll("input.txt")
	schematic := NewSchematic(input)

	for _, gear := range schematic.FindGears() {
		total += gear.Ratio()
	}

	return total
}
