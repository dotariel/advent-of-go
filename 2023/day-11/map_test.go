package main

import (
	"dotariel/inputreader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewMap(t *testing.T) {
	m := NewMap(inputreader.ReadAll("sample_input.txt"))

	assert.Len(t, m, 10)
	assert.Len(t, m[0], 10)
}

func TestImage_GetGalaxies(t *testing.T) {
	m := NewMap(inputreader.ReadAll("sample_input.txt"))

	assert.Len(t, m.Galaxies(1), 9)
}

func TestPoint_DistanceTo(t *testing.T) {
	assert.Equal(t, 9, Galaxy{1, 6}.DistanceTo(Galaxy{5, 11}))
	assert.Equal(t, 15, Galaxy{4, 0}.DistanceTo(Galaxy{9, 10}))
	assert.Equal(t, 17, Galaxy{0, 2}.DistanceTo(Galaxy{12, 7}))
	assert.Equal(t, 5, Galaxy{0, 11}.DistanceTo(Galaxy{5, 11}))
}

func TestImage_GetGalaxyPair(t *testing.T) {
	m := NewMap(inputreader.ReadAll("sample_input.txt"))
	pairs := m.Galaxies(1).Pairs()

	assert.Len(t, pairs, 36)
}

func TestImage_GetTotalDistance(t *testing.T) {
	m := NewMap(inputreader.ReadAll("sample_input.txt"))

	assert.Equal(t, 374, m.GetTotalDistance(1))
	assert.Equal(t, 1030, m.GetTotalDistance(10-1))
	assert.Equal(t, 8410, m.GetTotalDistance(100-1))

}
