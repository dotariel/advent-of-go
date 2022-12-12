package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIncreases(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	assert.Equal(t, 7, GetIncreases(input, 1))
	assert.Equal(t, 5, GetIncreases(input, 3))
}
