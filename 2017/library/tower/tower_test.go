package tower

import (
	"reflect"
	"testing"
)

func TestBuildNode(t *testing.T) {
	testCases := []struct {
		input    string
		expected *Node
		children []string
	}{
		{
			"mqdjo (83)",
			&Node{Name: "mqdjo", Weight: 83},
			[]string{},
		},
		{
			"blah (12) -> foo, bar, baz",
			&Node{Name: "blah", Weight: 12},
			[]string{"foo", "bar", "baz"},
		},
		{
			"ugml (68) -> gyxo, ebii, jptl",
			&Node{Name: "ugml", Weight: 68},
			[]string{"gyxo", "ebii", "jptl"},
		},
	}

	for _, tt := range testCases {
		if actual, _ := BuildNode(tt.input); !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("failed: %v -> wanted:%v, but got:%v", tt.input, tt.expected, actual)
		}
	}
}

func TestNew(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			`pbga (66)
			 xhth (57)
			 ebii (61)
			 havc (66)
			 ktlj (57)
			 fwft (72) -> ktlj, cntj, xhth
			 qoyq (66)
			 padx (45) -> pbga, havc, qoyq
			 tknk (41) -> ugml, padx, fwft
			 jptl (61)
			 ugml (68) -> gyxo, ebii, jptl
			 gyxo (61)
			 cntj (57)`,
			"tknk",
		},
	}

	for _, tt := range testCases {
		if actual := New(tt.input).Name; actual != tt.expected {
			t.Errorf("failed: %v -> wanted:%v, but got:%v", tt.input, tt.expected, actual)
		}
	}
}
func TestNode_HasChildren(t *testing.T) {
	testCases := []struct {
		input    *Node
		expected bool
	}{
		{&Node{}, false},
		{&Node{Children: nil}, false},
		{&Node{Children: []*Node{}}, false},
		{&Node{Children: []*Node{&Node{}, &Node{}}}, true},
	}

	for _, tt := range testCases {
		if actual := tt.input.HasChildren(); actual != tt.expected {
			t.Errorf("failed %v; wanted:%v but got:%v", tt.input, tt.expected, actual)
		}
	}
}

func TestNode_SubWeight(t *testing.T) {
	testCases := []struct {
		input    *Node
		expected int
	}{
		{&Node{Weight: 10}, 0},
		{&Node{Weight: 10, Children: []*Node{
			&Node{Weight: 5},
			&Node{Weight: 8},
		}}, 13},
		{&Node{Weight: 99, Children: []*Node{
			&Node{Weight: 2, Children: []*Node{
				&Node{Weight: 3},
				&Node{Weight: 4},
			}},
			&Node{Weight: 3, Children: []*Node{
				&Node{Weight: 6},
				&Node{Weight: 1},
			}},
		}}, 19},
	}

	for _, tt := range testCases {
		if actual := tt.input.SubWeight(0); actual != tt.expected {
			t.Errorf("failed %v; wanted:%v but got:%v", tt.input, tt.expected, actual)
		}
	}
}

func TestNode_IsBalanced(t *testing.T) {
	testCases := []struct {
		input    *Node
		expected bool
	}{
		{&Node{Weight: 0}, true},
		{&Node{Weight: 0, Children: []*Node{
			&Node{Weight: 5},
			&Node{Weight: 8},
		}}, false},
		{&Node{Weight: 0, Children: []*Node{
			&Node{Weight: 5},
			&Node{Weight: 5},
		}}, true},
		{&Node{Weight: 0, Children: []*Node{
			&Node{Weight: 15, Children: []*Node{
				&Node{Weight: 6},
				&Node{Weight: 4},
				&Node{Weight: 4},
			}},
			&Node{Weight: 13, Children: []*Node{
				&Node{Weight: 5},
				&Node{Weight: 5},
				&Node{Weight: 4},
			}},
		}}, false},
		{&Node{Weight: 0, Children: []*Node{
			&Node{Children: []*Node{&Node{Weight: 6}, &Node{Weight: 4}, &Node{Weight: 4}}},
			&Node{Children: []*Node{&Node{Weight: 0}, &Node{Weight: 0}, &Node{Weight: 0}}},
		}}, false},
		{&Node{Weight: 10, Children: []*Node{
			&Node{Weight: 6, Children: []*Node{
				&Node{Weight: 3},
				&Node{Weight: 7},
			}},
			&Node{Weight: 6, Children: []*Node{
				&Node{Weight: 4},
				&Node{Weight: 6},
			}},
		}}, true},
		{&Node{Weight: 10, Children: []*Node{
			&Node{Weight: 6, Children: []*Node{
				&Node{Weight: 4},
				&Node{Weight: 7},
			}},
			&Node{Weight: 6, Children: []*Node{
				&Node{Weight: 4},
				&Node{Weight: 6},
			}},
		}}, false},
	}

	for _, tt := range testCases {
		if actual := tt.input.IsBalanced(); actual != tt.expected {
			t.Errorf("failed %v; wanted:%v but got:%v", tt.input, tt.expected, actual)
		}
	}
}

func TestFindMismatch(t *testing.T) {
	input := `pbga (66)
			 xhth (57)
			 ebii (61)
			 havc (66)
			 ktlj (57)
			 fwft (72) -> ktlj, cntj, xhth
			 qoyq (66)
			 padx (45) -> pbga, havc, qoyq
			 tknk (41) -> ugml, padx, fwft
			 jptl (61)
			 ugml (68) -> gyxo, ebii, jptl
			 gyxo (61)
			 cntj (57)`
	tree := New(input)

	expected := 60

	if actual := FindMismatch(tree); actual != expected {
		t.Errorf("mismatch test failed; wanted:%v, got:%v", expected, actual)
	}
}
