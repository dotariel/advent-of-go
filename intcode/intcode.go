package intcode

type Program struct {
	InitialState []int
}

func NewProgram(state []int) Program {
	return Program{InitialState: state}
}

func (p Program) Run(noun int, verb int) int {
	memory := append([]int(nil), p.InitialState...)
	memory[1] = noun
	memory[2] = verb

	return Intcode(memory)[0]
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
