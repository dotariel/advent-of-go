package main

import (
	"strconv"
	"strings"
)

type Grid [][]int

func Parse(s string) Grid {
	g := make([][]int, 0)

	for _, line := range strings.Split(s, "\n") {
		if len(line) == 0 {
			continue
		}

		row := make([]int, 0)

		for _, entry := range line {
			n, _ := strconv.Atoi(string(entry))
			row = append(row, n)
		}

		g = append(g, row)
	}

	return g
}

func NewGrid(size int) Grid {
	grid := make([][]int, 0)

	for i := 0; i < size; i++ {
		grid = append(grid, make([]int, size))
	}

	return grid
}

func (g Grid) Row(index int) []int {
	return g[index]
}

func (g Grid) Column(index int) []int {
	col := make([]int, 0)

	for _, row := range g {
		col = append(col, row[index])
	}

	return col
}

func (g Grid) Get(r, c int) int {
	return g[r][c]
}

func (g Grid) CountVisible() int {
	numVisible := 0

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			n := g[i][j]

			if i == 0 || i == len(g)-1 || j == 0 || j == len(g)-1 {
				numVisible++
				continue
			}

			row := g.Row(i)
			col := g.Column(j)

			before := row[0:j]
			after := row[j+1:]
			above := col[0:i]
			below := col[i+1:]

			obBefore := Obstructed(n, before)
			obAfter := Obstructed(n, after)
			obAbove := Obstructed(n, above)
			obBelow := Obstructed(n, below)

			if !obBefore || !obAfter || !obBelow || !obAbove {
				numVisible++
			}
		}
	}

	return numVisible
}

func Obstructed(n int, ns []int) bool {
	for _, val := range ns {
		if val >= n {
			return true
		}
	}

	return false
}
