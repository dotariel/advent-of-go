package main

import "testing"
import "reflect"

func TestIntcode(t *testing.T) {
	testCases := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, tc := range testCases {
		if intCode := Intcode(tc.input); !reflect.DeepEqual(intCode, tc.output) {
			t.Errorf("Expected %v, but got %v", tc.output, intCode)
		}
	}
}
