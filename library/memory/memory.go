package memory

import (
	"reflect"
	"strconv"
	"strings"
)

type State []int

func NewState(input string) State {
	state := make([]int, 0)

	for _, item := range strings.Fields(input) {
		blocks, _ := strconv.ParseInt(item, 10, 64)
		state = append(state, int(blocks))
	}

	return state
}

func (state State) Redistribute() {
	// Start with bank with the highest number. If more than one bank has
	// the highest number, use the bank with the lowest position.
	bank, blocks := FindLargest(state)

	// Reset the blocks to 0 since they will be redistributed.
	state[bank] = 0

	// Starting with the next bank, distribute one block to each bank,
	// wrapping around to the beginning, until all the blocks have been
	// redistributed.
	nextBank := bank + 1
	for distributed := blocks; distributed > 0; distributed-- {
		if nextBank >= len(state) {
			nextBank = 0
		}

		state[nextBank]++
		nextBank++
	}
}

func (state State) CountCycles() int {
	states := make([]State, 0)
	cycles := 0

	for {
		state.Redistribute()
		cycles++

		t := make([]int, len(state))
		copy(t, state)
		if Contains(states, t) {
			return cycles
		}

		states = append(states, t)
	}
}

func FindLargest(state []int) (int, int) {
	max := 0
	pos := 0

	for i, blocks := range state {
		if blocks > max {
			max = blocks
			pos = i
		}
	}

	return pos, max
}

func Contains(states []State, state State) bool {
	for _, s := range states {
		if reflect.DeepEqual(state, s) {
			return true
		}
	}

	return false
}
