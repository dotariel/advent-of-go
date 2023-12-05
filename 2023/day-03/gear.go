package main

type Gear []int

func NewGear(parts []Part) Gear {
	return Gear([]int{parts[0].Value(), parts[1].Value()})
}

func (g Gear) Ratio() int {
	return g[0] * g[1]
}
