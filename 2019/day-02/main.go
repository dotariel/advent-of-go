package main

import (
	"dotariel/inputreader"
	"dotariel/intcode"
	"fmt"
)

const (
	add      = 1
	multiply = 2
	halt     = 99
)

func main() {
	inputs := inputreader.ReadInts("input.txt", ",")

	Part1(inputs)
	Part2(inputs)
}

func Part1(inputs []int) {
	fmt.Printf("Part 1: %v\n", intcode.NewProgram(inputs).Run(12, 2))
}

func Part2(inputs []int) {
	p := intcode.NewProgram(inputs)
	desired := 19690720
	noun := 0
	verb := 0

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if actual := p.Run(i, j); actual == desired {
				noun = i
				verb = j
				break
			}
		}
	}

	fmt.Printf("Part 2: noun:%v, verb:%v; total:%v\n", noun, verb, (100*noun + verb))
}
