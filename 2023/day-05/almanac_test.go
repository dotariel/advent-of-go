package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var almanacInput = `seeds: 79 14 55 13
	
seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

var simpleMap = `seed-to-soil map:
50 98 2
52 50 48`

func TestParseSeeds(t *testing.T) {
	assert.Equal(t, []int{79, 14, 55, 13}, parseSeeds(almanacInput))
}

func TestParseMappings(t *testing.T) {
	expected := []Map{
		Map{
			Source:      "seed",
			Destination: "soil",
			Ranges: []Range{
				Range{
					Destination: 50,
					Source:      98,
					Length:      2,
				},
				Range{
					Destination: 52,
					Source:      50,
					Length:      48,
				},
			},
		},
	}

	assert.Equal(t, expected, parseMappings(simpleMap))
}

func TestNewAlmanac(t *testing.T) {
	almanac := NewAlmanac(almanacInput)

	assert.Equal(t, []int{79, 14, 55, 13}, almanac.Seeds)
	assert.Len(t, almanac.Mappings, 7)
}

func TestAlmanac_FindMapBySource(t *testing.T) {
	almanac := NewAlmanac(almanacInput)

	testCases := []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity"}

	for _, tc := range testCases {
		mapping, _ := almanac.FindMapBySource(tc)
		assert.Equal(t, tc, mapping.Source)
	}
}

func TestAlmanac_FindMapByDestination(t *testing.T) {
	almanac := NewAlmanac(almanacInput)

	testCases := []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}

	for _, tc := range testCases {
		mapping, _ := almanac.FindMapByDestination(tc)
		assert.Equal(t, tc, mapping.Destination)
	}
}

func TestAlmanac_FindLocationBySeed(t *testing.T) {
	a := NewAlmanac(almanacInput)

	assert.Equal(t, 82, a.FindLocationBySeed(79))
	assert.Equal(t, 43, a.FindLocationBySeed(14))
	assert.Equal(t, 86, a.FindLocationBySeed(55))
	assert.Equal(t, 35, a.FindLocationBySeed(13))
	assert.Equal(t, 46, a.FindLocationBySeed(82)) // from Part 2
	assert.Equal(t, 999, a.FindLocationBySeed(999))
}
