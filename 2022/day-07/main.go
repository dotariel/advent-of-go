package main

import (
	"dotariel/inputreader"
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %v\n", Part1())
	fmt.Printf("Part 2: %v\n", Part2())
}

func Part1() interface{} {
	n := NewNavigator()
	inputs := inputreader.ReadStrings("input.txt", "\n")

	n.ParseHistory(inputs)
	n.Go("/")

	total := 0
	for _, d := range n.CurrentDir.Collect(100000) {
		total += d.Size()
	}

	return total
}

func Part2() interface{} {
	n := NewNavigator()
	inputs := inputreader.ReadStrings("input.txt", "\n")

	n.ParseHistory(inputs)
	n.Go("/")

	totalSpace := 70000000
	requiredSpace := 30000000
	usedSpace := n.CurrentDir.Size()
	unusedSpace := totalSpace - usedSpace
	amountToDelete := requiredSpace - unusedSpace

	dirs := n.CurrentDir.Collect(9999999999)
	sort.Sort(BySize(dirs))

	for _, dir := range dirs {
		if dir.Size() >= amountToDelete {
			return dir.Size()
		}
	}

	return 0
}
