package main

import (
	"dotariel/inputreader"
)

type Position struct {
	x int
	y int
}

type Slope struct {
	x int
	y int
}

type Grid struct {
	position Position
	items    [][]rune
}

func main() {
	grid := newGrid()
	grid.load(inputreader.ReadStrings("input.txt", "\n"))

	slopes := []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	trees := make([]int, 0)

	for _, slope := range slopes {
		trees = append(trees, grid.countTrees(slope))
	}

	total := 1

	for _, val := range trees {
		total *= val
	}

	println(total)
}

func newGrid() Grid {
	return Grid{position: Position{0, 0}, items: make([][]rune, 0)}
}

func (g *Grid) reset() {
	g.position = Position{0, 0}
}

func (g *Grid) load(input []string) {
	for _, line := range input {
		row := make([]rune, 0)

		for _, char := range line {
			row = append(row, char)

		}

		g.items = append(g.items, row)
	}
}

func (g *Grid) countTrees(slope Slope) int {
	trees := 0

	for {
		char, done := g.move(slope)

		if char == "#" {
			trees++
		}

		if done {
			break
		}
	}

	g.reset()

	return trees
}

func (g *Grid) move(slope Slope) (string, bool) {
	g.position.x = (g.position.x + slope.x) % g.width()
	g.position.y = g.position.y + slope.y

	done := (g.position.y == g.height()-1)

	return string(g.items[g.position.y][g.position.x]), done
}

func (g *Grid) width() int {
	return len(g.items[0])
}

func (g *Grid) height() int {
	return len(g.items)
}
