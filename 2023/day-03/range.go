package main

type Range struct {
	start int
	end   int
}

func (r Range) Intersects(other Range) bool {
	if r.start >= other.start && r.start <= other.end {
		return true
	}

	if r.end >= other.start && r.end <= other.end {
		return true
	}

	return false

}
