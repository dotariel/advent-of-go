package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	entries = []int{1721, 979, 366, 299, 675, 1456}
}

func TestFindSumTwo(t *testing.T) {
	assert.Equal(t, 514579, findSumTwo((2020)))
}

func TestFindSumThree(t *testing.T) {
	assert.Equal(t, 241861950, findSumThree((2020)))
}
