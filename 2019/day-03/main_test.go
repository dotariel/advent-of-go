package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewPoint(t *testing.T) {
	p := NewPoint(5, 8)

	assert.Equal(t, 5, p.trail[0].x)
	assert.Equal(t, 8, p.trail[0].y)
}

func TestMove(t *testing.T) {
	testCases := []struct {
		move string
		x    int
		y    int
	}{
		{"R1", 1, 0},
		{"U6", 1, 6},
		{"R5", 6, 6},
		{"U1", 6, 7},
		{"D3", 6, 4},
	}

	p := origin.Clone()

	for _, tc := range testCases {
		p.Move(tc.move)
		assert.Equal(t, p.x, tc.x)
		assert.Equal(t, p.y, tc.y)
	}
}

var intersectTestCases = []struct {
	w1       []string
	w2       []string
	distance int
	steps    int
}{
	{
		[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
		[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
		159,
		610,
	},
	{
		[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
		[]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
		135,
		410,
	},
}

func TestIntersects(t *testing.T) {
	for _, tc := range intersectTestCases {
		p1 := origin.Clone()
		p2 := origin.Clone()
		p1.Path(tc.w1)
		p2.Path(tc.w2)

		distance := -1
		for _, point := range p1.Intersects(p2) {
			if !point.Equals(origin) {
				if d := point.DistanceFrom(origin); distance < 0 || d < distance {
					distance = d
				}
			}
		}

		if distance != tc.distance {
			t.Errorf("You fucked up! Wanted %v, but got %v", tc.distance, distance)
		}
	}
}

var distanceTestCases = []struct {
	x        int
	y        int
	distance int
}{
	{1, 1, 2},
	{2, 3, 5},
	{8, 12, 20},
}

func TestDistanceFrom(t *testing.T) {
	for _, tc := range distanceTestCases {
		assert.Equal(t, tc.distance, NewPoint(tc.x, tc.y).DistanceFrom(origin))
	}
}

func BenchmarkDistanceFrom(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tc := range distanceTestCases {
			NewPoint(tc.x, tc.y).DistanceFrom(origin)
		}
	}
}

func TestStepsToIntersection(t *testing.T) {
	for _, tc := range intersectTestCases {
		p1 := origin.Clone()
		p2 := origin.Clone()
		p1.Path(tc.w1)
		p2.Path(tc.w2)

		steps := 0
		for _, point := range p1.Intersects(p2) {
			if d := p1.StepsFrom(point) + p2.StepsFrom(point); steps < 1 || d < steps {
				steps = d
			}
		}
		assert.Equal(t, tc.steps, steps)
	}
}
