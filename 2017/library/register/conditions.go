package register

type condition func(int, int) bool

// NewCondition creates a new condition from a string
func NewCondition(c string) condition {
	if c == "<" {
		return lt
	}

	if c == "<=" {
		return lte
	}

	if c == ">" {
		return gt
	}

	if c == ">=" {
		return gte
	}

	if c == "!=" {
		return neq
	}

	return eq
}

func noop(n, x int) int {
	return n
}

func eq(a, b int) bool {
	return a == b
}

func neq(a, b int) bool {
	return !eq(a, b)
}

func gt(a, b int) bool {
	return a > b
}

func gte(a, b int) bool {
	return a >= b
}

func lt(a, b int) bool {
	return a < b
}

func lte(a, b int) bool {
	return a <= b
}
