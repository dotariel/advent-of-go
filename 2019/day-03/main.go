package main

import (
	"dotariel/inputreader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x     int
	y     int
	trail []*Point
}

var origin = NewPoint(0, 0)

var path1 []string
var path2 []string

func main() {
	data := inputreader.ReadStrings("input.txt", "\n")

	path1 = strings.Split(data[0], ",")
	path2 = strings.Split(data[1], ",")

	// Part1()
	Part2()
}

func Part1() {
	p1 := origin.Clone()
	p2 := origin.Clone()

	p1.Path(path1)
	p2.Path(path2)

	distance := -1
	for _, point := range p1.Intersects(p2) {
		if !point.Equals(origin) {
			if d := point.DistanceFrom(origin); distance < 0 || d < distance {
				distance = d
			}
		}
	}

	fmt.Printf("Shortest distance is: %v\n", distance)
}

func Part2() {
	p1 := origin.Clone()
	p2 := origin.Clone()

	p1.Path(path1)
	p2.Path(path2)

	steps := 0
	for _, point := range p1.Intersects(p2) {
		if d := p1.StepsFrom(point) + p2.StepsFrom(point); steps < 1 || d < steps {
			steps = d
		}
	}

	fmt.Printf("Fewest steps: %v\n", steps)
}

func NewPoint(x int, y int) *Point {
	return &Point{x: x, y: y, trail: []*Point{&Point{x: x, y: y}}}
}

func (p *Point) Move(d string) {
	runes := []rune(d)
	dir := runes[0]
	dist, _ := strconv.Atoi(string(runes[1:len(runes)]))

	for i := 0; i < dist; i++ {
		if 'R' == dir {
			p.x++
		}

		if 'L' == dir {
			p.x--
		}

		if 'U' == dir {
			p.y++
		}

		if 'D' == dir {
			p.y--
		}
		p.trail = append(p.trail, NewPoint(p.x, p.y))
	}
}

func (p *Point) Intersects(q *Point) []*Point {
	xs := make([]*Point, 0)

	for _, i := range p.trail {

		ch := make(chan []*Point)
		go func(trail []*Point) {
			localXs := make([]*Point, 0)
			for _, j := range trail {
				if i.Equals(j) {
					localXs = append(localXs, i)
				}
			}
			ch <- localXs
		}(q.trail)

		xs = append(xs, <-ch...)
	}

	return xs
}

func (p *Point) Path(movements []string) {
	for _, movement := range movements {
		p.Move(movement)
	}
}

func (p *Point) DistanceFrom(q *Point) int {
	return int(math.Abs(float64(p.x-q.x)) + math.Abs(float64(p.y-q.y)))
}

func (p *Point) Equals(q *Point) bool {
	return p.x == q.x && p.y == q.y
}

func (p *Point) Clone() *Point {
	return NewPoint(p.x, p.y)
}

func (p *Point) StepsFrom(q *Point) int {
	for idx, point := range p.trail {
		if point.Equals(q) {
			return idx
		}
	}

	return -1
}
