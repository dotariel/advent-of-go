package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pipe = connectors["|"]
var f = connectors["F"]

func TestConnector_ConnectsFrom(t *testing.T) {
	assert.True(t, pipe.ConnectsFrom("S"))
	assert.True(t, pipe.ConnectsFrom("N"))

	assert.True(t, f.ConnectsFrom("N"))
	assert.True(t, f.ConnectsFrom("W"))
}

func TestConnector_Next(t *testing.T) {
	assert.Equal(t, NORTH, pipe.Next(SOUTH))
	assert.Equal(t, SOUTH, pipe.Next(NORTH))
	assert.Equal(t, EAST, f.Next(SOUTH))
	assert.Equal(t, SOUTH, f.Next(EAST))
}
