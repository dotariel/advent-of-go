package main

import (
	"dotariel/inputreader"
)

var entries = make([]int, 0)

func main() {
	load()

	println("Part 1: %v", findSumTwo((2020)))
	println("Part 2: %v", findSumThree((2020)))
}

func load() {
	for _, val := range inputreader.ReadInts("input.txt", "\n") {
		entries = append(entries, val)
	}
}

func findSumTwo(n int) int {
	for _, a := range entries {
		for _, b := range entries {
			if (a + b) == n {
				return a * b
			}
		}
	}

	return 0
}

func findSumThree(n int) int {
	for _, a := range entries {
		for _, b := range entries {
			for _, c := range entries {
				if (a + b + c) == n {
					return a * b * c
				}
			}
		}
	}

	return 0
}
