package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Network struct {
	Instructions []string
	Nodes
}

type Node struct {
	Key   string
	Left  string
	Right string
}

type Nodes map[string]Node

func NewNetwork(input string) Network {
	return Network{
		Instructions: getInstructions(input),
		Nodes:        getNodes(input),
	}
}

func getInstructions(input string) []string {
	instructions := make([]string, 0)

	r := regexp.MustCompile(`([RL]+)(?:\n+)`)

	if matches := r.FindStringSubmatch(input); len(matches) > 1 {
		instructions = append(instructions, strings.Split(matches[1], "")...)
	}

	return instructions
}

func getNodes(input string) map[string]Node {
	r := regexp.MustCompile(`([\w]{3}) = \(([\w]{3}), ([\w]{3})\)\n?`)

	nodes := make(map[string]Node, 0)

	for _, match := range r.FindAllStringSubmatch(input, -1) {
		nodes[match[1]] = Node{Key: match[1], Left: match[2], Right: match[3]}
	}

	return nodes
}

func (n Network) Locate(node string) (Node, error) {
	if node, exists := n.Nodes[node]; exists {
		return node, nil
	}

	return Node{}, fmt.Errorf("node '%v' not found", node)
}

func (n Network) Traverse() int {
	start, _ := n.Locate("AAA")
	hops, _ := track(n, start, 0, Stop)

	return hops
}

func (n Network) TraverseParallel() int {
	paths := make([]Node, 0)

	for _, node := range n.Nodes {
		if node.IsStartingNode() {
			paths = append(paths, node)
		}
	}

	counter := 0

	multiples := make([]int, 0)

	for _, node := range paths {
		count, _ := track(n, node, counter, StopAtEnd)
		multiples = append(multiples, count)
	}

	return LCM(multiples[0], multiples[1], multiples[2:]...)
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func track(n Network, node Node, counter int, stop StopFunc) (int, Node) {
	if stop(node) {
		return counter, node
	}

	direction := n.Instructions[counter%len(n.Instructions)]
	next, err := n.Locate(node.Next(direction))
	if err != nil {
		panic(err)
	}

	counter++

	return track(n, next, counter, stop)
}

func (n Node) IsStartingNode() bool {
	return strings.HasSuffix(n.Key, "A")
}

func (n Node) IsEndingNode() bool {
	return strings.HasSuffix(n.Key, "Z")
}

func (n Node) Next(direction string) string {
	if "L" == direction {
		return n.Left
	}

	if "R" == direction {
		return n.Right
	}

	return ""
}

type StopFunc func(n Node) bool

func Stop(n Node) bool {
	return n.Key == "ZZZ"
}

func StopAtEnd(n Node) bool {
	return n.IsEndingNode()
}
