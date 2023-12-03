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
		total += NewCalibration(input).GetValue()
	}

	return total
}

func Part2() interface{} {
	total := 0
	for _, input := range inputreader.ReadStrings("input.txt", "\n") {
		total += NewCalibration(input).GetExtendedValue()
	}

	return total
}
