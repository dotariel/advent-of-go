package main

type CubeSet map[string]int

func (cs CubeSet) Power() int {
	power := 1

	for _, count := range cs {
		power *= count
	}

	return power
}
