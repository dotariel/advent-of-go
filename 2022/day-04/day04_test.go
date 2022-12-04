package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	actual := Parse("11-73,29-73")

	expected := []RangeDef{
		{Lower: 11, Upper: 73},
		{Lower: 29, Upper: 73},
	}

	assert.Equal(t, expected, actual)
}

func TestRangeDef_ContainsRangeDef(t *testing.T) {
	rd1 := NewRangeDef("2-8")
	rd2 := NewRangeDef("3-7")

	assert.True(t, rd1.Contains(rd2))
	assert.True(t, rd1.Contains(rd1))
	assert.False(t, rd2.Contains(rd1))
}

func TestRangeDef_ToRange(t *testing.T) {
	rd := NewRangeDef("2-8")
	expected := Range([]int{2, 3, 4, 5, 6, 7, 8})

	assert.Equal(t, expected, rd.ToRange())
}

func TestRange_Overlaps(t *testing.T) {
	r1 := Range([]int{1, 2, 3})
	r2 := Range([]int{3, 4, 5})
	r3 := Range([]int{4, 5, 6})

	assert.True(t, r1.Overlaps(r2))
	assert.True(t, r2.Overlaps(r1))
	assert.False(t, r1.Overlaps(r3))
	assert.True(t, r2.Overlaps(r3))
}
