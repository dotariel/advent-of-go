package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGear(t *testing.T) {
	parts := []Part{{"9", Range{0, 0}}, {"8", Range{0, 0}}}

	gear := NewGear(parts)

	assert.Equal(t, 72, gear.Ratio())
}

func TestGear_Ratio(t *testing.T) {
	assert.Equal(t, 50, Gear([]int{5, 10}).Ratio())
}
