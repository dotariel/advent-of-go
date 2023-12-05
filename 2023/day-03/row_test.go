package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestRow_HasPartInRange(t *testing.T) {
// 	testCases := []struct {
// 		row      Row
// 		expected []Part
// 	}{
// 		{row: Row{"467..114.."}, expected: []Part{{"467"}}},
// 		{row: Row{"...*......"}, expected: []Part{}},
// 		{row: Row{"..35..633."}, expected: []Part{{"35"}}},
// 	}

// 	for _, tc := range testCases {
// 		assert.Equal(t, tc.expected, tc.row.FindPartsNear(3))
// 	}
// }

func TestRow_GetSearchRange(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected []Range
	}{
		{input: "...*......", expected: []Range{{2, 4}}},
		{input: "*.........", expected: []Range{{0, 1}}},
		{input: ".........*", expected: []Range{{8, 9}}},
		{input: ".*....*...", expected: []Range{{0, 2}, {5, 7}}},
	} {
		assert.Equal(t, tc.expected, Row{tc.input}.GetSearchRanges())
	}
}

func TestRow_FindPartsThatIntersect(t *testing.T) {

}
