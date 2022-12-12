package main

func IsIncrease(previous, current int) bool {
	return current > previous
}

func GetIncreases(ns []int, offset int) int {
	previous := 0
	increases := 0

	for i := 0; i < len(ns)-(offset-1); i++ {
		val := Sum(ns[i : i+offset])

		if val > previous {
			if i > 0 {
				increases++
			}
		}

		previous = val
	}

	return increases
}

func Sum(ns []int) int {
	sum := 0

	for _, n := range ns {
		sum = sum + n
	}

	return sum
}
