package main

import (
	"fmt"
	"math"
	"strings"
)

type Map [][]string

type Galaxy struct {
	Row, Col int
}

type Galaxies []Galaxy

func NewMap(input string) Map {
	m := [][]string{}

	for _, line := range strings.Split(input, "\n") {
		m = append(m, strings.Split(line, ""))
	}

	return m
}

func (m Map) Print() {
	for _, row := range m {
		for _, col := range row {
			fmt.Printf("%v", col)
		}
		fmt.Println()
	}
}

func (m Map) Galaxies(expansion int) Galaxies {
	galaxies := make([]Galaxy, 0)
	offset := 0

	for row := 0; row < len(m); row++ {
		hasGalaxy := false

		for col := 0; col < len(m[row]); col++ {
			if m[row][col] == "#" {
				galaxies = append(galaxies, Galaxy{Row: row + offset, Col: col})
				hasGalaxy = true
			}
		}

		if !hasGalaxy {
			offset += expansion
		}
	}

	cols := len(m[0])
	for col := 0; col < cols; col++ {
		hasGalaxy := false

		for _, g := range galaxies {
			if g.Col == col {
				hasGalaxy = true
				break
			}
		}

		if !hasGalaxy {
			for j := 0; j < len(galaxies); j++ {
				if galaxies[j].Col > col {
					galaxies[j].Col += expansion
				}
			}

			cols += expansion
			col += expansion
		}
	}

	return galaxies
}

func (p Galaxy) DistanceTo(other Galaxy) int {
	x1, x2 := float64(p.Col), float64(other.Col)
	y1, y2 := float64(p.Row), float64(other.Row)

	return int(math.Abs(x1-x2) + math.Abs(y1-y2))
}

func (gs Galaxies) Pairs() [][2]Galaxy {
	pairs := make([][2]Galaxy, 0)

	for i := 0; i < len(gs)-1; i++ {
		for j := i + 1; j < len(gs); j++ {
			pairs = append(pairs, [2]Galaxy{gs[j], gs[i]})
		}
	}

	return pairs
}

func (m Map) GetTotalDistance(expansion int) int {
	pairs := m.Galaxies(expansion).Pairs() //m.GetGalaxyPairs(expansion)
	distance := 0

	for _, pair := range pairs {
		distance += pair[0].DistanceTo(pair[1])
	}

	return distance
}
