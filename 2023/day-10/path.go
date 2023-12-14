package main

import "math"

type Path []Point

func (p Path) Max() int {
	return (len(p) / 2)
}

func (p Path) Exists(point Point) bool {
	for _, pt := range p {
		if pt == point {
			return true
		}
	}
	return false
}

func (p Path) Area() int {
	n := len(p)
	res := 0

	for i := range p {
		x1, y1 := p[i].x, p[i].y
		x2, y2 := p[(i+1)%n].x, p[(i+1)%n].y

		res += x1*y2 - x2*y1
	}

	return int(math.Abs(float64(res))) >> 1
}

func (p Path) NumInteriorPoints() int {
	return p.Area() - (len(p) / 2) + 1
}
