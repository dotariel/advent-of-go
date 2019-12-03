package main

import (
	"dotariel/inputreader"
	"fmt"
)

func main() {
	inputs := inputReader.ReadInts("input.txt", ",")

	// Place in '1202 program alarm' state
	inputs[1] = 12
	inputs[2] = 2

	intCode := Intcode(inputs)

	fmt.Println(intCode[0])
}

func Intcode(input []int) []int {
	output := append([]int(nil), input...)

	for i := 0; i < len(output); i++ {
		code := output[i]

		if code == 99 {
			break
		}

		if code == 1 || code == 2 {
			op1 := output[output[i+1]]
			op2 := output[output[i+2]]
			pos := output[i+3]

			if code == 1 {
				output[pos] = op1 + op2
			} else {
				output[pos] = op1 * op2
			}

			i += 3
		}
	}

	return output
}
