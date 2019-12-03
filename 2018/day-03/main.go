package main

import (
	"dotariel/inputreader"
	"fmt"
	"regexp"
	"strconv"
)

type Claim map[string]int

type Coord struct {
	x int
	y int
}

type Rect struct {
	topLeft     Coord
	bottomRight Coord
}

func main() {
	inputs := inputreader.ReadStrings("input.txt")
	claims := make([]Claim, 0)

	for _, input := range inputs {
		claims = append(claims, ParseClaim(input))
	}

	overlap := GetOverlap(claims)
	fmt.Println(overlap)

	return
	for _, claim := range claims {
		claim["overlapped"] = 0

		for _, other := range claims {
			if claim.OverlapsWith(other) && claim["id"] != other["id"] {
				claim["overlapped"] = 1
				break
			}
		}

		if claim["overlapped"] == 0 {
			fmt.Printf("Found it! %v\n", claim["id"])
		}
	}

}

func ParseClaim(s string) Claim {
	r := regexp.MustCompile(`#(?P<id>\d+) @ (?P<left>\d+),(?P<top>\d+): (?P<width>\d+)x(?P<height>\d+)`)
	res := r.FindStringSubmatch(s)
	names := r.SubexpNames()
	claim := make(map[string]int)

	for i, _ := range res {
		if i != 0 {
			val, _ := strconv.Atoi(res[i])
			claim[names[i]] = val
		}
	}

	return Claim(claim)
}

func GetOverlap(claims []Claim) int {
	overlap := 0
	coords := make(map[Coord]int)

	for _, claim := range claims {
		for _, coord := range claim.GetCoords() {
			if claimed, exists := coords[coord]; !exists {
				coords[coord] = 1
			} else if claimed == 1 {
				coords[coord]++
				overlap++
			}
		}
	}

	return overlap
}

func (c Claim) GetCoords() []Coord {
	coords := make([]Coord, 0)

	for x := c["left"]; x < c["left"]+c["width"]; x++ {
		for y := c["top"]; y < c["top"]+c["height"]; y++ {
			coords = append(coords, Coord{x, y})
		}
	}

	return coords
}

func (c Claim) OverlapsWith(other Claim) bool {
	a := c.GetBoundingRect()
	b := other.GetBoundingRect()

	if a.topLeft.y > b.bottomRight.y || a.bottomRight.y < b.topLeft.y {
		return false
	}

	if a.topLeft.x > b.bottomRight.x || a.bottomRight.x < b.topLeft.x {
		return false
	}

	return true
}

func (c Claim) GetBoundingRect() Rect {
	coords := c.GetCoords()
	return Rect{coords[0], coords[len(coords)-1]}
}

func Contains(coords []Coord, coord Coord) bool {
	for _, val := range coords {
		if val == coord {
			return true
		}
	}

	return false
}
