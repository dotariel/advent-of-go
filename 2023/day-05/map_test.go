package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_GetDestinationValue(t *testing.T) {
	m := Map{
		Source:      "seed",
		Destination: "soil",
		Ranges: []Range{
			{
				Destination: 50,
				Source:      98,
				Length:      2,
			},
			{
				Destination: 52,
				Source:      50,
				Length:      48,
			},
		},
	}

	assert.Equal(t, 0, m.GetDestinationValue(0))
	assert.Equal(t, 1, m.GetDestinationValue(1))
	assert.Equal(t, 48, m.GetDestinationValue(48))
	assert.Equal(t, 49, m.GetDestinationValue(49))
	assert.Equal(t, 53, m.GetDestinationValue(51))
	assert.Equal(t, 98, m.GetDestinationValue(96))
	assert.Equal(t, 99, m.GetDestinationValue(97))
	assert.Equal(t, 50, m.GetDestinationValue(98))
	assert.Equal(t, 51, m.GetDestinationValue(99))
}
