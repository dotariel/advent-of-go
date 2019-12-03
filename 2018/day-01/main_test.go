package main

import "testing"

func TestGetFrequency(t *testing.T) {
	testCases := []struct {
		changes  []int
		expected int
	}{
		{[]int{1, -2, 3, 1}, 3},
		{[]int{1, 1, 1}, 3},
		{[]int{1, 1, -2}, 0},
		{[]int{-1, -2, -3}, -6},
	}

	for _, tc := range testCases {
		diff := GetFrequency(tc.changes)

		if diff != tc.expected {
			t.Errorf("Expected %v but got %v", tc.expected, diff)
		}
	}
}

func TestGetDuplicateFrequency(t *testing.T) {
	testCases := []struct {
		changes  []int
		expected int
	}{
		{[]int{1, -2, 3, 1}, 2},
		{[]int{3, 3, 4, -2, -4}, 10},
		{[]int{-6, 3, 8, 5, -6}, 5},
		{[]int{7, 7, -2, -7, -4}, 14},
	}

	for _, tc := range testCases {
		diff := GetDuplicateFrequency(tc.changes)

		if diff != tc.expected {
			t.Errorf("Expected %v but got %v", tc.expected, diff)
		}
	}
}
