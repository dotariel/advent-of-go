package stack

import (
	"strconv"
	"strings"
)

// Stack represents a list of instructions
type Stack []int

// OffsetResolver represents a function that can resolve the value
// of a given offset
type OffsetResolver func(int) int

// SimpleIncrementer is an OffsetResolver function that increase
// the offset by 1
func SimpleIncrementer(offset int) int {
	return offset + 1
}

// BiasedDecrementer is an OffsetResolver function that decrements
// the offset by 1 if the offset is 3 or more, and increments by 1
// otherwise
func BiasedDecrementer(offset int) int {
	if offset >= 3 {
		return offset - 1
	}

	return offset + 1
}

// Trace processes the instructions in the Stack and returns the
// number of instructions processed before leaving the Stack.
func (stack Stack) Trace(resolver OffsetResolver) int {
	pos := 0
	jumps := 0

	for {
		if pos < 0 || pos > len(stack)-1 {
			break
		}

		offset := stack[pos]

		// Resolve the offset based on the provided resolver
		stack[pos] = resolver(stack[pos])

		// Jump to the next instruction
		pos += offset

		// Increment the number of jumps
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
