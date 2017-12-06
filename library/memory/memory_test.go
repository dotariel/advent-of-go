package memory

import (
	"reflect"
	"testing"
)

var redistributeTestCases = []struct {
	input    []int
	expected []int
}{
	{[]int{0, 2, 7, 0}, []int{2, 4, 1, 2}},
	{[]int{2, 4, 1, 2}, []int{3, 1, 2, 3}},
	{[]int{3, 1, 2, 3}, []int{0, 2, 3, 4}},
}

func TestRedistribute(t *testing.T) {
	for _, tt := range redistributeTestCases {
		actual := State(tt.input)
		actual.Redistribute()

		if reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("test failed for: %v; wanted: %v, got: %v", tt.input, tt.expected, actual)
		}
	}
}

var cycleCountTestCases = []struct {
	input    []int
	expected int
}{
	{[]int{0, 2, 7, 0}, 5},
}

func TestCycle(t *testing.T) {
	for _, tt := range cycleCountTestCases {
		if actual := State(tt.input).CountCycles(); actual != tt.expected {
			t.Errorf("test failed for: %v; wanted: %v, got: %v", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkRedistribute(b *testing.B) {
	b.StopTimer()
	for _, tt := range cycleCountTestCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			State(tt.input).Redistribute()
		}

		b.StopTimer()
	}
}

func TestFindLargest(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 2, 7, 0}, []int{2, 7}},
		{[]int{0, 2, 7, 7}, []int{2, 7}},
		{[]int{8, 7, 8, 1}, []int{0, 8}},
		{[]int{0, 1, 2, 3}, []int{3, 3}},
		{[]int{3, 2, 1, 1}, []int{0, 3}},
	}

	for _, tt := range testCases {
		if actualPos, actualValue := FindLargest(tt.input); actualPos != tt.expected[0] || actualValue != tt.expected[1] {
			t.Errorf("test failed for: %v; wanted: %v, got: %v|%v", tt.input, tt.expected, actualPos, actualValue)
		}
	}
}
