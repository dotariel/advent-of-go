package hash

import (
	"reflect"
	"testing"
)

func TestSublist(t *testing.T) {
	list := []uint8{0, 1, 2, 3, 4, 5, 6}

	testCases := []struct {
		input        []uint8
		position     int
		length       int
		expectedList []uint8
		expectedPos  []int
		error        bool
	}{
		{list, 0, 8, nil, nil, true},                                                  // length larger than list size
		{list, 7, 7, nil, nil, true},                                                  // starting position too big
		{list, -1, 7, nil, nil, true},                                                 // starting position too small
		{list, 0, 2, []uint8{0, 1}, []int{0, 1}, false},                               // simple subset from zero
		{list, 1, 4, []uint8{1, 2, 3, 4}, []int{1, 2, 3, 4}, false},                   // simple subset from non-zero start
		{list, 0, 7, []uint8{0, 1, 2, 3, 4, 5, 6}, []int{0, 1, 2, 3, 4, 5, 6}, false}, // entire list
		{list, 4, 5, []uint8{4, 5, 6, 0, 1}, []int{4, 5, 6, 0, 1}, false},             // wraps around
	}

	for _, tt := range testCases {
		actual, _, err := sublist(tt.input, tt.position, tt.length)
		if !reflect.DeepEqual(actual, tt.expectedList) {
			t.Errorf("failed %v; wanted:%v, but got:%v", tt.input, tt.expectedList, actual)
		}

		if (err != nil) != tt.error {
			t.Errorf("failed error assertion %v; wanted:%v, but got:%v", tt.input, tt.error, actual)
		}
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		input    []uint8
		expected []uint8
	}{
		{[]uint8(nil), []uint8(nil)},
		{[]uint8{4, 3, 2, 1}, []uint8{1, 2, 3, 4}},
		{[]uint8{0, 0, 1, 2, 3, 0}, []uint8{0, 3, 2, 1, 0, 0}},
	}

	for _, tt := range testCases {
		if actual := reverse(tt.input); !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("failed %v; wanted:%v, but got:%v", tt.input, tt.expected, actual)
		}
	}
}

func TestDo(t *testing.T) {
	testCases := []struct {
		input    []uint8
		position int
		length   int
		expected []uint8
		error    bool
	}{
		{[]uint8{1, 2, 3}, 4, 2, nil, true},               // bad position
		{[]uint8{1, 2, 3}, 0, 4, nil, true},               // bad length
		{[]uint8{1, 2, 3}, 0, 2, []uint8{2, 1, 3}, false}, // starting at zero
		{[]uint8{1, 2, 3}, 1, 2, []uint8{1, 3, 2}, false}, // starting after zero
		{[]uint8{1, 2, 3}, 2, 2, []uint8{3, 2, 1}, false}, // wrapping around
	}

	for _, tt := range testCases {
		actual, err := do(tt.input, tt.position, tt.length)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("failed %v; wanted:%v, but got:%v", tt.input, tt.expected, actual)
		}

		if (err != nil) != tt.error {
			t.Errorf("failed error assertion %v; wanted:%v, but got:%v", tt.input, tt.error, actual)
		}
	}
}
