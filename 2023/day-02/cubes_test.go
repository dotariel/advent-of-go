package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCubeSet_Power(t *testing.T) {
	cubeSet := CubeSet{"red": 2, "blue": 3, "green": 4}

	assert.Equal(t, 24, cubeSet.Power())
}
