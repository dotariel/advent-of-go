package main

import (
	"regexp"
)

type Row struct {
	string
}

func (row Row) GetSearchRanges() []Range {
	r := regexp.MustCompile(`\*`)

	ranges := make([]Range, 0)

	for _, match := range r.FindAllStringIndex(row.string, -1) {
		ranges = append(ranges, Range{max(0, match[0]-1), min(len(row.string)-1, match[1])})
	}

	return ranges
}

func (row Row) FindPartsThatIntersect(rng Range) []Part {
	parts := make([]Part, 0)

	r := regexp.MustCompile(`[\d]+`)

	for _, match := range r.FindAllStringIndex(row.string, -1) {
		ps := row.string[match[0]:match[1]]
		part := Part{ps, Range{match[0], match[1] - 1}}

		if part.Intersects(rng) {
			parts = append(parts, part)
		}
	}

	return parts
}
