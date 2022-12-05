package main

import (
	"regexp"
	"strconv"
)

type Stack []rune

type Stacks []Stack

type Instruction struct {
	Count int
	From  int
	To    int
}

func (self *Stack) Pop() rune {
	var r rune

	slice := *self
	length := len(slice) - 1
	r, *self = slice[length], slice[:length]

	return r
}

func (self *Stack) Push(r rune) {
	*self = append(*self, r)
}

func (self *Stack) Move(target *Stack, n int) {
	for i := 0; i < n; i++ {
		target.Push(self.Pop())
	}
}

func (self *Stack) Graft(target *Stack, n int) {
	popped := make([]rune, 0)
	for i := 0; i < n; i++ {
		popped = append(popped, self.Pop())
	}

	for i := len(popped) - 1; i >= 0; i-- {
		target.Push(popped[i])
	}
}

func NewInstruction(s string) Instruction {
	rex, _ := regexp.Compile(`move (?P<count>\d+) from (?P<from>\d+) to (?P<to>\d+)`)
	match := rex.FindStringSubmatch(s)

	result := make(map[string]string)
	for i, name := range rex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	count, _ := strconv.Atoi(result["count"])
	from, _ := strconv.Atoi(result["from"])
	to, _ := strconv.Atoi(result["to"])

	return Instruction{
		Count: count,
		From:  from,
		To:    to,
	}
}

func (st Stacks) Execute(instructions []Instruction) {
	for _, i := range instructions {
		st[i.From-1].Move(&st[i.To-1], i.Count)
	}
}

func (st Stacks) Forklift(instructions []Instruction) {
	for _, i := range instructions {
		st[i.From-1].Graft(&st[i.To-1], i.Count)
	}
}
