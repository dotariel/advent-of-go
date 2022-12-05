package exercise

import "fmt"

type ExerciseFn func() interface{}

type Exercise struct {
	Part1 ExerciseFn
	Part2 ExerciseFn
}

func New(part1 ExerciseFn, part2 ExerciseFn) Exercise {
	return Exercise{
		Part1: part1,
		Part2: part2,
	}
}

func (e Exercise) Run() {
	fmt.Printf("Part 1: %v\n", e.Part1())
	fmt.Printf("Part 2: %v\n", e.Part2())
}
