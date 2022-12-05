package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Pop(t *testing.T) {
	stack := Stack([]rune("abc"))

	assert.Equal(t, 'c', stack.Pop())
	assert.Equal(t, 'b', stack.Pop())
	assert.Equal(t, 'a', stack.Pop())
	assert.Empty(t, stack)
}

func TestStack_Push(t *testing.T) {
	stack := Stack([]rune("abc"))

	stack.Push('d')
	stack.Push('e')
	stack.Push('f')

	assert.Equal(t, "abcdef", string(stack))
}

func TestNewInstruction(t *testing.T) {
	i := NewInstruction("move 1 from 2 to 3")

	assert.Equal(t, 1, i.Count)
	assert.Equal(t, 2, i.From)
	assert.Equal(t, 3, i.To)
}

func TestStack_Move(t *testing.T) {
	stacks := Stacks([]Stack{
		Stack([]rune("NZ")),
		Stack([]rune("DCM")),
	})

	instructions := []Instruction{
		NewInstruction("move 1 from 1 to 2"),
	}

	stacks.Execute(instructions)

	assert.Equal(t, Stack([]rune("N")), stacks[0])
	assert.Equal(t, Stack([]rune("DCMZ")), stacks[1])
}

func TestStack_Graft(t *testing.T) {
	s1 := Stack([]rune("ABCDEF"))
	s2 := Stack([]rune("GHIJKL"))

	s1.Graft(&s2, 3)

	assert.Equal(t, Stack([]rune("ABC")), s1)
	assert.Equal(t, Stack([]rune("GHIJKLDEF")), s2)
}
