package main

type Map struct {
	Source      string
	Destination string

	Ranges []Range
}

func (m Map) GetDestinationValue(id int) int {
	for _, rng := range m.Ranges {
		start := rng.Source
		end := rng.Source + rng.Length

		if id >= start && id <= end {
			return rng.Destination + (id - start)
		}
	}

	return id
}
