package main

import (
	"fmt"
	"io/ioutil"

	"github.com/dotariel/advent-of-go/library/stack"
)

func main() {
	bytes, _ := ioutil.ReadFile("input")
	input := string(bytes)

	fmt.Println("Part 1: ", stack.Jump(stack.NewStack(input)))
}
