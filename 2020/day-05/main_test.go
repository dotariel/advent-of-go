package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	assert.True(t, reflect.DeepEqual(Decode("FBFBBFFRLR"), Seat{44, 5}))
}

func TestSeat_Id(t *testing.T) {
	testCases := []struct {
		seat     Seat
		expected int
	}{
		{Seat{44, 5}, 357},
		{Seat{70, 7}, 567},
		{Seat{14, 7}, 119},
		{Seat{102, 4}, 820},
		{Seat{89, 2}, 714},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.seat.Id(), tc.expected)
	}
}
