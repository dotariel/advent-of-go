package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildScore(t *testing.T) {
	assert.Equal(t, Score{1, 0, 0, 0, 0, 0}, NewScore(1, nil))
	assert.Equal(t, Score{1, 2, 3, 4, 5, 0}, NewScore(1, []int{2, 3}, []int{4, 5}))
}

func TestScore_Compare(t *testing.T) {
	assert.Equal(t, 1, NewScore(1).Compare(NewScore(0)))
	assert.Equal(t, -1, NewScore(0).Compare(NewScore(1)))
	assert.Equal(t, 0, NewScore(1).Compare(NewScore(1)))
}
