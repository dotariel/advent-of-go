package spiral

import "fmt"

func Distance(n float64) float64 {

	pos := 1
	value := 1

	// Loop forever, and break out only once the target is found.
	for square := 1; square < 4; square++ {
		// Each i is constructing a "square". During square construction,
		// determine if the target number is in the square, and return its
		// distance.

		// Each square has i*8 items in it
		items := square * 8
		min := square
		max := square * 2

		for j := 0; j < items; j++ {

			// value = (j % 4) + 1
			mod := (j + 1) % (8 - (2 * square))

			// ((items/2 - j) % items)

			pos++
			fmt.Printf("square:%v, items:%v, j:%v, pos:%v, mod:%v, min:%v, max:%v, value:%v\n", square, items, j, pos, mod, min, max, value)
		}

		// i++
	}
	return float64(value)

}
