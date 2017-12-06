package memory

import (
	"testing"
)

var cases = []struct {
	input    []int
	expected int
}{
	{[]int{0, 2, 7, 0}, 5},
}

func TestRedistribute(t *testing.T) {
	for _, tt := range cases {
		if actual := State(tt.input).Redistribute(); actual != tt.expected {
			t.Errorf("test failed for: %v; wanted: %v, got: %v", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkRedistribute(b *testing.B) {
	b.StopTimer()
	for _, tt := range cases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			State(tt.input).Redistribute()
		}

		b.StopTimer()
	}
}

func TestFindLargest(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 2, 7, 0}, []int{2, 7}},
		{[]int{0, 2, 7, 7}, []int{2, 7}},
		{[]int{8, 7, 8, 1}, []int{0, 8}},
		{[]int{0, 1, 2, 3}, []int{3, 3}},
		{[]int{3, 2, 1, 1}, []int{0, 3}},
	}

	for _, tt := range cases {
		if actualPos, actualValue := FindLargest(tt.input); actualPos != tt.expected[0] || actualValue != tt.expected[1] {
			t.Errorf("test failed for: %v; wanted: %v, got: %v|%v", tt.input, tt.expected, actualPos, actualValue)
		}
	}
}
