package main

type Connectors map[string]Connector

type Connector struct {
	Connects []Direction
}

var connectors = map[string]Connector{
	"F": {Connects: []Direction{SOUTH, EAST}},
	"7": {Connects: []Direction{SOUTH, WEST}},
	"L": {Connects: []Direction{NORTH, EAST}},
	"J": {Connects: []Direction{NORTH, WEST}},
	"-": {Connects: []Direction{WEST, EAST}},
	"|": {Connects: []Direction{NORTH, SOUTH}},
	"S": {Connects: []Direction{NORTH, SOUTH, EAST, WEST}},
}

func (c Connector) ConnectsTo(direction Direction) bool {
	for _, d := range c.Connects {
		if d == direction {
			return true
		}
	}

	return false
}

func (c Connector) ConnectsFrom(direction Direction) bool {
	for _, d := range c.Connects {
		if d == direction.Inverse() {
			return true
		}
	}

	return false
}

func (c Connector) Next(from Direction) Direction {
	for _, direction := range c.Connects {
		if from != direction {
			return direction
		}
	}

	return NONE
}
