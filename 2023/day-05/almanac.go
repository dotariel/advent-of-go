package main

import (
	"dotariel/util"
	"errors"
	"regexp"
	"strings"
)

type Almanac struct {
	Seeds    []int
	Mappings []Map
}

func NewAlmanac(s string) Almanac {
	return Almanac{
		Seeds:    parseSeeds(s),
		Mappings: parseMappings(s),
	}
}

func parseSeeds(s string) []int {
	rex := regexp.MustCompile(`(?:seeds: )([\d\s]+)`)

	if matches := rex.FindAllStringSubmatch(s, -1); len(matches) > 0 {
		return util.ToInts(strings.Fields(matches[0][1]))
	}

	return nil
}

func parseMappings(s string) []Map {
	mappings := make([]Map, 0)

	rex := regexp.MustCompile(`([\w]+)-to-([\w]+) map:(?:\n)([\d\s]+)\n?`)

	for _, match := range rex.FindAllStringSubmatch(s, -1) {
		mapping := Map{
			Source:      match[1],
			Destination: match[2],
			Ranges:      []Range{},
		}

		for _, line := range strings.Split(match[3], "\n") {
			if parts := util.ToInts(strings.Fields(line)); len(parts) == 3 {
				rng := Range{
					Destination: parts[0],
					Source:      parts[1],
					Length:      parts[2],
				}

				mapping.Ranges = append(mapping.Ranges, rng)
			}
		}

		mappings = append(mappings, mapping)
	}

	return mappings
}

func (a Almanac) FindMapBySource(source string) (Map, error) {
	for _, mapping := range a.Mappings {
		if mapping.Source == source {
			return mapping, nil
		}
	}

	return Map{}, errors.New("source not found")
}

func (a Almanac) FindMapByDestination(destination string) (Map, error) {
	for _, mapping := range a.Mappings {
		if mapping.Destination == destination {
			return mapping, nil
		}
	}

	return Map{}, errors.New("destination not found")
}

func (a Almanac) FindLocationBySeed(seed int) int {
	current := a.Mappings[0]
	id := seed

	for {
		id = current.GetDestinationValue(id)

		next, err := a.FindMapBySource(current.Destination)
		if err != nil {
			break
		}

		current = next
	}

	return id
}

func (a Almanac) GetSeedRanges() [][]int {
	seedranges := [][]int{}

	i := 0
	for {
		seedranges = append(seedranges, []int{a.Seeds[i], a.Seeds[i] + a.Seeds[i+1] - 1})

		if len(seedranges) == len(a.Seeds)/2 {
			break
		}
		i += 2
	}

	return seedranges
}
