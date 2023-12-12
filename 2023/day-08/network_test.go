package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	input1 = `RL
	
AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

	input2 = `LLR
	
AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

	input3 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`
)

func Test_NewNetwork(t *testing.T) {
	expected := Network{
		Instructions: []string{"R", "L"},
		Nodes: map[string]Node{
			"AAA": {Key: "AAA", Left: "BBB", Right: "CCC"},
			"BBB": {Key: "BBB", Left: "DDD", Right: "EEE"},
			"CCC": {Key: "CCC", Left: "ZZZ", Right: "GGG"},
			"DDD": {Key: "DDD", Left: "DDD", Right: "DDD"},
			"EEE": {Key: "EEE", Left: "EEE", Right: "EEE"},
			"GGG": {Key: "GGG", Left: "GGG", Right: "GGG"},
			"ZZZ": {Key: "ZZZ", Left: "ZZZ", Right: "ZZZ"},
		},
	}

	assert.Equal(t, expected, NewNetwork(input1))
}

func TestNetwork_Locate(t *testing.T) {
	network := NewNetwork(input1)

	testCases := []struct {
		input string
		error error
		node  Node
	}{
		{input: "AAA", error: nil, node: Node{"AAA", "BBB", "CCC"}},
		{input: "XXX", error: errors.New("node 'XXX' not found"), node: Node{}},
	}

	for _, tc := range testCases {
		node, err := network.Locate(tc.input)

		assert.Equal(t, tc.error, err)
		assert.Equal(t, tc.node, node)
	}
}

func TestNetwork_Traverse(t *testing.T) {
	assert.Equal(t, 6, NewNetwork(input2).Traverse())
}

func TestNetwork_TraverseParallel(t *testing.T) {
	network := NewNetwork(input3)

	assert.Equal(t, 6, network.TraverseParallel())
}

func Test_NewNetwork_Alt(t *testing.T) {
	network := NewNetwork(input3)

	expected := Network{
		Instructions: []string{"L", "R"},
		Nodes: map[string]Node{
			"11A": {Key: "11A", Left: "11B", Right: "XXX"},
			"11B": {Key: "11B", Left: "XXX", Right: "11Z"},
			"11Z": {Key: "11Z", Left: "11B", Right: "XXX"},
			"22A": {Key: "22A", Left: "22B", Right: "XXX"},
			"22B": {Key: "22B", Left: "22C", Right: "22C"},
			"22C": {Key: "22C", Left: "22Z", Right: "22Z"},
			"22Z": {Key: "22Z", Left: "22B", Right: "22B"},
			"XXX": {Key: "XXX", Left: "XXX", Right: "XXX"},
		},
	}

	assert.Equal(t, expected, network)
}

func TestNode_IsStartingNode(t *testing.T) {
	assert.True(t, Node{Key: "11A"}.IsStartingNode())
	assert.True(t, Node{Key: "22A"}.IsStartingNode())
	assert.False(t, Node{Key: "XXX"}.IsStartingNode())
}

func TestNode_IsEndingNode(t *testing.T) {
	assert.True(t, Node{Key: "22Z"}.IsEndingNode())
	assert.False(t, Node{Key: "11A"}.IsEndingNode())
	assert.False(t, Node{Key: "22A"}.IsEndingNode())
}
