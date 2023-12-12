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
	report := NewReport(inputreader.ReadAll("input.txt"))

	total := 0

	for _, history := range report {
		total += history.PredictNext()
	}

	return total
}

func Part2() interface{} {
	report := NewReport(inputreader.ReadAll("input.txt"))

	total := 0

	for _, history := range report {
		total += history.PredictPrevious()
	}

	return total
}
