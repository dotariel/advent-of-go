package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var g = Grid{
	[]int{1, 2, 3},
	[]int{4, 5, 6},
	[]int{7, 8, 9},
}

func TestParse(t *testing.T) {
	input := `
30373
25512
65332
33549
35390
`
	g := Parse(input)

	assert.Len(t, g, 5)
}

func TestGrid_Row(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, g.Row(0))
	assert.Equal(t, []int{4, 5, 6}, g.Row(1))
	assert.Equal(t, []int{7, 8, 9}, g.Row(2))
}

func TestGrid_Column(t *testing.T) {
	assert.Equal(t, []int{1, 4, 7}, g.Column(0))
	assert.Equal(t, []int{2, 5, 8}, g.Column(1))
	assert.Equal(t, []int{3, 6, 9}, g.Column(2))
}

func TestGrid_Get(t *testing.T) {
	assert.Equal(t, 1, g.Get(0, 0))
}

func TestGrid_CountVisible(t *testing.T) {
	input := `
30373
25512
65332
33549
35390`

	g := Parse(input)

	assert.Equal(t, 21, g.CountVisible())
}

func TestObstructed(t *testing.T) {
	assert.True(t, Obstructed(5, []int{1, 2, 3, 4, 5}))
	assert.False(t, Obstructed(6, []int{1, 2, 3, 4, 5}))
}
