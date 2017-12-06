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

func (banks State) Redistribute() int {
	states := make([]State, 0)
	cycles := 0

	for {
		// Start with bank with the highest number. If more than one bank has
		// the highest number, use the bank with the lowest position.
		bank, blocks := FindLargest(banks)

		// Reset the blocks to 0 since they will be redistributed.
		banks[bank] = 0

		// Starting with the next bank, distribute one block to each bank,
		// wrapping around to the beginning, until all the blocks have been
		// redistributed.
		nextBank := bank + 1
		for distributed := blocks; distributed > 0; distributed-- {
			if nextBank >= len(banks) {
				nextBank = 0
			}

			banks[nextBank]++
			nextBank++
		}

		// Increment cycle count
		cycles++

		// If new allocation has been seen before, break out and return
		t := make([]int, len(banks))
		copy(t, banks)
		if HasState(states, t) {
			break
		}

		states = append(states, t)
	}

	return cycles
}

func FindLargest(banks []int) (int, int) {
	max := 0
	pos := 0

	for i, blocks := range banks {
		if blocks > max {
			max = blocks
			pos = i
		}
	}

	return pos, max
}

func HasState(states []State, state State) bool {
	for _, s := range states {
		if reflect.DeepEqual(state, s) {
			return true
		}
	}

	return false
}
