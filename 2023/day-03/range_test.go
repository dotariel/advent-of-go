package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange_Intersects(t *testing.T) {
	assert.True(t, Range{4, 6}.Intersects(Range{3, 5}))
	assert.True(t, Range{3, 5}.Intersects(Range{4, 6}))
	assert.True(t, Range{0, 3}.Intersects(Range{0, 1}))
	assert.True(t, Range{0, 3}.Intersects(Range{3, 4}))
	assert.False(t, Range{0, 3}.Intersects(Range{4, 5}))
	assert.False(t, Range{4, 5}.Intersects(Range{0, 3}))
}
