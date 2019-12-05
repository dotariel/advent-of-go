package main

import (
	"dotariel/inputreader"
)

func main() {
	changes := inputreader.ReadInts("input.txt", "\n")

	println(GetFrequency(changes))
	println(GetDuplicateFrequency(changes))
}

func GetFrequency(changes []int) int {
	diff := 0
	for _, i := range changes {
		diff += i
	}

	return diff
}

func GetDuplicateFrequency(changes []int) int {
	frequencies := make([]int, 0)
	frequency := 0

	for {
		for _, f := range changes {
			frequency += f

			if exists(frequency, frequencies) {
				return frequency
			}

			frequencies = append(frequencies, frequency)
		}
	}

	return frequency

}

func exists(val int, vals []int) (exists bool) {
	for _, v := range vals {
		if v == val {
			exists = true
			break
		}
	}
	return
}
