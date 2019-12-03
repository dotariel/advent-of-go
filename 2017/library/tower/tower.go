package tower

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Node struct {
	Name     string
	Weight   int
	Parent   *Node
	Children []*Node
}

func New(input string) *Node {
	var root *Node
	nodes := make(map[*Node][]string)

	var findNodeByName = func(name string) *Node {
		for node := range nodes {
			if node.Name == name {
				return node
			}
		}
		return nil
	}

	for _, line := range strings.Split(input, "\n") {
		node, children := BuildNode(line)
		nodes[node] = children
	}

	for node, children := range nodes {
		for _, ch := range children {
			child := findNodeByName(ch)
			child.Parent = node
			node.Children = append(node.Children, child)
		}
	}

	for node := range nodes {
		if node.Parent == nil {
			root = node
		}
	}

	return root
}

func FindMismatch(node *Node) int {
	if !node.IsBalanced() {
		childrenBalanced := true

		for _, child := range node.Children {
			childrenBalanced = childrenBalanced && child.IsBalanced()
		}

		if childrenBalanced {
			return node.Mismatch()
		}

		for _, child := range node.Children {
			return FindMismatch(child)
		}
	}

	return -1
}

func buildString(node *Node, indent int, buffer string) string {
	buffer += fmt.Sprintf("%v%v (%v)(%v); balanced:%v\n", strings.Repeat(" ", indent*4), node.Name, node.Weight, node.TotalWeight(), node.IsBalanced())

	if node.HasChildren() {
		indent++
		for _, child := range node.Children {
			buffer = fmt.Sprintf("%v", buildString(child, indent, buffer))
		}
		indent--
	}

	return buffer
}

func BuildNode(input string) (*Node, []string) {
	exp := regexp.MustCompile("(?P<name>[a-z]+) \\((?P<weight>[0-9]+)\\)( -> (?P<children>.*))?")
	match := exp.FindStringSubmatch(input)

	result := make(map[string]string)
	for i, name := range exp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	weight, _ := strconv.ParseInt(result["weight"], 10, 64)
	node := Node{
		Name:     result["name"],
		Weight:   int(weight),
		Parent:   nil,
		Children: nil,
	}

	children := make([]string, 0)
	for _, child := range strings.FieldsFunc(result["children"], splitFunc) {
		children = append(children, child)
	}

	return &node, children
}

func (n *Node) SubWeight(acc int) int {
	if !n.HasChildren() {
		return acc
	}

	for _, child := range n.Children {
		acc += child.SubWeight(child.Weight)
	}

	return acc
}

func (n *Node) TotalWeight() int {
	return n.SubWeight(n.Weight)
}

func (n *Node) HasChildren() bool {
	return len(n.Children) > 0
}

func (n *Node) IsBalanced() bool {
	if !n.HasChildren() {
		return true
	}

	var lastWeight int
	for i, child := range n.Children {
		if i == 0 {
			lastWeight = child.TotalWeight()
		}

		if newWeight := child.TotalWeight(); newWeight != lastWeight {
			return false
		}
	}

	return true
}

func (n *Node) Walk() {
	fmt.Printf(buildString(n, 0, ""))
}

func (n *Node) Mismatch() int {
	values := make(map[int][]*Node)

	for _, child := range n.Children {
		key := child.TotalWeight()
		values[key] = append(values[key], child)
	}

	correctTotal := 0
	outlier := &Node{}
	wrongTotal := 0

	for total, nodes := range values {
		if len(nodes) == 1 {
			outlier = nodes[0]
			wrongTotal = total
		} else {
			correctTotal = total
		}
	}

	return outlier.Weight + (correctTotal - wrongTotal)
}

func splitFunc(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsPunct(r)
}
