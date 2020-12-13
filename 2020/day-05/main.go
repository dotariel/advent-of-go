package main

import (
	"dotariel/inputreader"
)

func main() {
	list := make([]Seat, 0)

	for _, entry := range inputreader.ReadStrings("input.txt", "\n") {
		list = append(list, Decode(entry))
	}

	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			lookup := Seat{i, j}

			if !exists(list, lookup.Id()) {
				println("NOT FOUND", lookup.row, lookup.col)
			}
		}
	}
}

type Seat struct {
	row int
	col int
}

func (s Seat) Id() int {
	return (s.row * 8) + s.col
}

func Decode(s string) Seat {
	seats := make([]int, 0)
	rows := make([]int, 0)

	for i := 0; i < 128; i++ {
		seats = append(seats, i)
	}

	for i := 0; i < 8; i++ {
		rows = append(rows, i)
	}

	for _, char := range s[0:7] {
		if char == 'F' {
			seats = seats[0 : len(seats)/2]
		}

		if char == 'B' {
			seats = seats[len(seats)/2:]
		}
	}

	for _, char := range s[7:] {
		if char == 'L' {
			rows = rows[0 : len(rows)/2]
		}

		if char == 'R' {
			rows = rows[len(rows)/2:]
		}
	}

	return Seat{seats[0], rows[0]}
}

func exists(seats []Seat, id int) bool {
	for _, seat := range seats {
		if seat.Id() == id {
			return true
		}
	}
	return false
}
