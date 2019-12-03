package main

import "testing"

func TestGetModuleFuel(t *testing.T) {
	testCases := []struct {
		mass float64
		fuel float64
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, tc := range testCases {
		required := GetModuleFuel(tc.mass)

		if required != tc.fuel {
			t.Errorf("Expected %v, but got %v", tc.fuel, required)
		}
	}
}

func TestGetTotalFuel(t *testing.T) {
	testCases := []struct {
		mass float64
		fuel float64
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, tc := range testCases {
		required := GetTotalFuel(tc.mass, 0)

		if required != tc.fuel {
			t.Errorf("Expected %v, but got %v", tc.fuel, required)
		}
	}
}
