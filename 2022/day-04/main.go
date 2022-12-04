package main

import (
	"dotariel/inputreader"
	"fmt"
)

func main() {
	fmt.Printf("Part 1: %v\n", Part1())
	fmt.Printf("Part 2: %v\n", Part2())
}

func Part1() int {
	overlap := 0
	for _, line := range inputreader.ReadStrings("input.txt", "\n") {
		rds := Parse(line)

		rd1 := rds[0]
		rd2 := rds[1]

		if rd1.Contains(rd2) || rd2.Contains(rd1) {
			overlap++
		}
	}

	return overlap
}

func Part2() int {
	overlap := 0
	for _, line := range inputreader.ReadStrings("input.txt", "\n") {
		rds := Parse(line)

		r1 := rds[0].ToRange()
		r2 := rds[1].ToRange()

		if r1.Overlaps(r2) {
			overlap++
		}
	}

	return overlap
}
