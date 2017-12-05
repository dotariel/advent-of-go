package spiral

import (
	"testing"
)

func TestDistance(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		// {1, 0},
		// {2, 1},
		// {3, 2},
		// {4, 1},

		// {5, 2},
		// {6, 1},
		// {7, 2},
		// {8, 1},
		// {9, 2},

		// {10, 3},
		// {11, 2},
		// {12, 3},
		{13, 4},
		// {14, 3},
		// {15, 2},
		// {16, 3},
	}

	for _, tt := range testCases {
		if actual := Distance(tt.input); actual != tt.expected {
			t.Errorf("distance test failed for %v; wanted:%v, got:%v", tt.input, tt.expected, actual)
		}
	}
}

func TestDo(t *testing.T) {
	for i := 0; i < 100; i++ {

	}
}

/*
  square (4)  ->  4  [2,4,6,8]
	square (9)  ->  8  [3,5,7,9,11,15,19,23]
	square (16) -> 12  [10,12,14,16,18,20,22,24,28,34,40,46]
*/
