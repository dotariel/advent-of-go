package main

import (
	"regexp"
	"strings"
	"unicode"
)

type Schematic []Row

func NewSchematic(input string) Schematic {
	arr := make([]Row, 0)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		arr = append(arr, Row{line})
	}

	return arr
}

func (sch Schematic) FindParts() []Part {
	parts := make([]Part, 0)

	r := regexp.MustCompile(`[\d]+`)

	for idx, row := range sch {
		for _, match := range r.FindAllStringIndex(row.string, -1) {
			segment := Part{row.string[match[0]:match[1]], Range{match[0], match[1] - 1}}

			if sch.checkRow(idx, match) {
				parts = append(parts, segment)
			}

			if idx > 0 && sch.checkRow(idx-1, match) {
				parts = append(parts, segment)
			}

			if idx < len(sch)-1 && sch.checkRow(idx+1, match) {
				if sch.checkRow(idx+1, match) {
					parts = append(parts, segment)
				}
			}
		}
	}

	return parts
}

func (sch Schematic) FindGears() []Gear {
	gears := make([]Gear, 0)

	for idx, row := range sch {
		for _, rng := range row.GetSearchRanges() {
			intersects := make([]Part, 0)

			intersects = append(intersects, row.FindPartsThatIntersect(rng)...)

			if idx > 0 {
				intersects = append(intersects, sch[idx-1].FindPartsThatIntersect(rng)...)
			}

			if idx < len(sch)-1 {
				intersects = append(intersects, sch[idx+1].FindPartsThatIntersect(rng)...)
			}

			if len(intersects) > 1 {
				gears = append(gears, NewGear(intersects))
			}
		}
	}

	return gears
}

func (sch Schematic) checkRow(rowindex int, match []int) bool {
	row := sch[rowindex].string
	start := max(0, match[0]-1)
	end := min(len(row)-1, match[1]+1)
	segment := row[start:end]

	return hasSymbol(segment)
}

func isSymbol(r rune) bool {
	return !(unicode.IsDigit(r) || r == '.')
}

func isGearSymbol(r rune) bool {
	return r == '*'
}

func hasSymbol(s string) bool {
	for _, r := range s {
		if isSymbol(r) {
			return true
		}
	}

	return false
}
