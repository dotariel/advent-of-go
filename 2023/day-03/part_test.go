package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart_Value(t *testing.T) {
	assert.Equal(t, 0, Part{"xxx", Range{0, 0}}.Value())
	assert.Equal(t, 567, Part{"567", Range{0, 0}}.Value())
}

// func TestPart_Intersects(t *testing.T) {
// 	// 1234.......
// 	part := Part{"1234", Range{0, 3}}

// 	assert.True(t, part.Intersects(Range{0, 0}))
// 	assert.True(t, part.Intersects(Range{0, 1}))
// 	assert.True(t, part.Intersects(Range{2, 3}))
// }
