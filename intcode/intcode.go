package intcode

type Program struct {
	memory   []int
	output   []int
	position int
}

type Operation func(params ...Parameter)

type Parameter struct {
	mode  int
	value int
}

const (
	ParameterModePosition  = 1
	ParameterModeImmediate = 2
)

func NewProgram(memory []int) Program {
	return Program{memory: memory, output: append([]int(nil), memory...), position: 0}
}

func (p Program) Run(noun int, verb int) int {
	p.output[1] = noun
	p.output[2] = verb

	return p.Process()[0]
}

func (p Program) Process() []int {
	for i := 0; i < len(p.output); i++ {
		code := p.output[i]

		if code == 99 {
			break
		}

		operation := p.getOperation(code)

		op1 := p.output[p.output[i+1]]
		op2 := p.output[p.output[i+2]]
		pos := p.output[i+3]

		// p.output[pos] = operation(op1, op2)
		i += 3
	}

	return p.output
}

func (p Program) getOperation(code int) Operation {
	if code == 1 {
		return p.add
	}

	return nil
}

func (p Program) add(params ...Parameter) {

	a := p.getParameterValue(params[0])
	b := p.getParameterValue(params[1])
	pos := p.getParameterValue(params[2])

	p.output[pos] = a + b
}

func (p Program) getParameterValue(param Parameter) int {
	if param.mode == ParameterModePosition {
		return p.output[param.value]
	}

	return param.value
}

// TODO: Need to keep pointer to current position in the program
