package main

import (
	"errors"
	"fmt"
	"strings"
)

type Field [][]*Tile

func NewField(input string) Field {
	field := Field{}

	for y, line := range strings.Split(input, "\n") {
		tiles := make([]*Tile, 0)

		for x, symbol := range strings.Split(line, "") {
			tile := Tile{symbol: symbol, point: Point{x, y}}

			if connector, exists := connectors[symbol]; exists {
				tile.connector = &connector
			}

			tiles = append(tiles, &tile)
		}

		field = append(field, tiles)
	}

	return field
}

func (f Field) Print() {
	for y, tiles := range f {
		for x, tile := range tiles {
			fmt.Printf("(%v,%v)=%v ", x, y, tile)
		}

		fmt.Printf("\n")
	}
}

func (f Field) Get(p Point) *Tile {
	if p.x < 0 || p.y < 0 || p.y > len(f)-1 || p.x > len(f[0])-1 {
		return nil
	}

	return f[p.y][p.x]
}

func (f Field) Walk() Path {
	path := make([]Point, 0)

	origin, err := f.Start()
	if err != nil {
		panic(err)
	}

	direction := NONE
	steps := 0

	for _, d := range []Direction{NORTH, EAST, SOUTH, WEST} {
		if tile, _ := f.Next(origin, d); tile != nil {
			direction = d
			break
		}
	}

	if direction == NONE {
		panic("no where to go!")
	}

	current := origin
	for {
		steps++
		path = append(path, current.point)

		current, direction = f.Next(current, direction)
		if current == origin {
			break
		}
	}

	path = append(path, origin.point)

	return path
}

func (f Field) Next(tile *Tile, direction Direction) (*Tile, Direction) {
	candidate := f.Get(direction.From(tile.point))

	if candidate == nil {
		return nil, NONE
	}

	if candidate.connector == nil {
		return nil, NONE
	}

	if !candidate.connector.ConnectsFrom(direction) {
		return nil, NONE
	}

	return candidate, candidate.connector.Next(direction.Inverse())
}

func (f Field) Start() (*Tile, error) {
	for _, tiles := range f {
		for _, tile := range tiles {
			if tile.symbol == "S" {
				return tile, nil
			}
		}
	}

	return nil, errors.New("start tile not found")
}

func (f Field) PrintPath(p Path) {
	for _, tiles := range f {
		for _, tile := range tiles {

			if p.Exists(tile.point) {
				fmt.Printf(" x ")
			} else {
				fmt.Print(" . ")
			}
		}

		fmt.Print("\n\n")
	}
}

func (f Field) GetEnclosedCount() int {
	return 0
}
