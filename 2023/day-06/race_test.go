package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`

	expected := Races{
		{7, 9},
		{15, 40},
		{30, 200},
	}
	races := Parse(input)

	assert.Equal(t, expected, races)
}

func TestParseSingle(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`

	expected := Race{71530, 940200}

	assert.Equal(t, expected, ParseSingle(input))
}

func TestRace_GetChargeTimes(t *testing.T) {

	r1 := Race{7, 9}
	r2 := Race{15, 40}
	r3 := Race{30, 200}

	assert.Equal(t, []int{2, 3, 4, 5}, r1.GetChargeTimes())
	assert.Equal(t, []int{4, 5, 6, 7, 8, 9, 10, 11}, r2.GetChargeTimes())
	assert.Equal(t, []int{11, 12, 13, 14, 15, 16, 17, 18, 19}, r3.GetChargeTimes())

	margin := len(r1.GetChargeTimes()) * len(r2.GetChargeTimes()) * len(r3.GetChargeTimes())

	assert.Equal(t, 288, margin)
}
