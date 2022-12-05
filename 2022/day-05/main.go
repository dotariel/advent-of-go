package main

import (
	"dotariel/inputreader"
	"fmt"
)

func main() {
	fmt.Printf("Part 1: %v\n", Part1())
	fmt.Printf("Part 2: %v\n", Part2())
}

func Part1() string {
	result := make([]rune, 0)

	stacks, instructions := Setup()
	stacks.Execute(instructions)

	for _, stack := range stacks {
		result = append(result, stack[len(stack)-1])
	}

	return string(result)
}

func Part2() string {
	result := make([]rune, 0)

	stacks, instructions := Setup()
	stacks.Forklift(instructions)

	for _, stack := range stacks {
		result = append(result, stack[len(stack)-1])
	}

	return string(result)
}

func Setup() (Stacks, []Instruction) {
	stacks := Stacks([]Stack{
		Stack([]rune("TPZCSLQN")),
		Stack([]rune("LPTVHCG")),
		Stack([]rune("DCZF")),
		Stack([]rune("GWTDLMVC")),
		Stack([]rune("PWC")),
		Stack([]rune("PFJDCTSZ")),
		Stack([]rune("VWGBD")),
		Stack([]rune("NJSQHW")),
		Stack([]rune("RCQFSLV")),
	})

	instructions := make([]Instruction, 0)

	for _, i := range inputreader.ReadStrings("input.txt", "\n") {
		instructions = append(instructions, NewInstruction(i))
	}

	return stacks, instructions
}
