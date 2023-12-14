package main

type Direction string

var (
	NORTH = Direction("N")
	SOUTH = Direction("S")
	EAST  = Direction("E")
	WEST  = Direction("W")
	NONE  = Direction("")
)

func (d Direction) Inverse() Direction {
	if d == NORTH {
		return SOUTH
	}

	if d == SOUTH {
		return NORTH
	}

	if d == WEST {
		return EAST
	}

	if d == EAST {
		return WEST
	}

	return NONE
}

func (d Direction) From(p Point) Point {
	if d == NORTH {
		return Point{p.x, p.y - 1}
	}

	if d == SOUTH {
		return Point{p.x, p.y + 1}
	}

	if d == WEST {
		return Point{p.x - 1, p.y}
	}

	if d == EAST {
		return Point{p.x + 1, p.y}
	}

	return p
}
