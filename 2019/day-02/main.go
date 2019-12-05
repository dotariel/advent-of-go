package main

import (
	"dotariel/inputreader"
	"fmt"
)

type Program struct {
	InitialState []int
}

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
	fmt.Printf("Part 1: %v\n", NewProgram(inputs).Run(12, 2))
}

func Part2(inputs []int) {
	p := NewProgram(inputs)
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

func NewProgram(state []int) Program {
	return Program{InitialState: state}
}

func (p Program) Run(noun int, verb int) int {
	memory := append([]int(nil), p.InitialState...)
	memory[1] = noun
	memory[2] = verb

	return Intcode(memory)[0]
}

func Intcode(input []int) []int {
	output := append([]int(nil), input...)

	for i := 0; i < len(output); i++ {
		code := output[i]

		if code == 99 {
			break
		}

		if code == 1 || code == 2 {
			op1 := output[output[i+1]]
			op2 := output[output[i+2]]
			pos := output[i+3]

			if code == 1 {
				output[pos] = op1 + op2
			} else {
				output[pos] = op1 * op2
			}

			i += 3
		}
	}

	return output
}
