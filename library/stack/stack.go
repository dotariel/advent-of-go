package stack

import (
	"strconv"
	"strings"
)

// Stack represents a list of instructions
type Stack []int

// Trace processes the instructions in the Stack and returns the
// number of instructions processed before leaving the Stack.
func (stack Stack) Trace() int {
	pos := 0
	jumps := 0

	for {
		if pos < 0 || pos > len(stack)-1 {
			break
		}
		jump := stack[pos]
		stack[pos]++
		pos += jump
		jumps++
	}

	return jumps
}

// NewStack creates a new Stack from an input string
func NewStack(input string) Stack {
	instructions := make([]int, 0)

	for _, line := range strings.Split(input, "\n") {
		instruction, _ := strconv.ParseInt(line, 10, 64)
		instructions = append(instructions, int(instruction))
	}

	return Stack(instructions)
}
