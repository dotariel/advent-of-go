package intcode

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	input  []int
	output []int
	result int
}{
	{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}, 0},
	{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}, 0},
	{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}, 0},
	{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, 0},
}

func TestIntcode(t *testing.T) {
	for _, tc := range testCases {
		if intCode := Intcode(tc.input); !reflect.DeepEqual(intCode, tc.output) {
			t.Errorf("Expected %v, but got %v", tc.output, intCode)
		}
	}
}

func TestParseInstruction(t *testing.T) {
	// 1002, 4, 3, 4
	/*
		1002      - (02) = 2 = Multiply
		4,3,4,33  - Four parameters, we need a position mode for each
		parameters[0] = 0
		parameters[1] = 1
		parameters[2] = 0 (not present)

		This will multiply 3(position) x 4(immediate) and store in 4 (position)

		1 Add       - 3 args (a, b, result) -> adds a and b and saves to result
		2 Multiply  - 3 args (a, b, result) -> multiplies a and b and saves to result
		3 Put       - 1 arg (a)             -> saves a into result
		4 Get       - 1 arg (a)			    -> returns value at a


	*/

}
