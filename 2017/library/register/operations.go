package register

type operation func(int, int) int

// NewOperation constructs an operation from a string
func NewOperation(op string) operation {
	switch op {
	case "inc":
		return inc
	case "dec":
		return dec
	}

	return noop
}

func inc(n, x int) int {
	return n + x
}

func dec(n, x int) int {
	return n - x
}
