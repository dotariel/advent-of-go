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
	margin := 1

	for _, race := range Parse(inputreader.ReadAll("input.txt")) {
		margin *= len(race.GetChargeTimes())
	}

	return margin
}

func Part2() interface{} {
	race := ParseSingle(inputreader.ReadAll("input.txt"))

	return len(race.GetChargeTimes())
}
