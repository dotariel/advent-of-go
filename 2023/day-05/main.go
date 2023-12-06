package main

import (
	"dotariel/inputreader"
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Part 1: %v\n", Part1())
	fmt.Printf("Part 2: %v\n", Part2())
}

func Part1() interface{} {
	almanac := NewAlmanac(inputreader.ReadAll("input.txt"))
	low := math.MaxInt32

	for _, seed := range almanac.Seeds {
		low = min(low, almanac.FindLocationBySeed(seed))
	}

	return low
}

func Part2() interface{} {
	low := math.MaxInt32

	almanac := NewAlmanac(inputreader.ReadAll("input.txt"))

	for _, rng := range almanac.GetSeedRanges() {
		for seed := rng[0]; seed <= rng[1]; seed++ {
			low = min(low, almanac.FindLocationBySeed(seed))
		}
	}

	return low
}
