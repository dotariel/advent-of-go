package main

import "testing"
import "reflect"

func TestParseClaim(t *testing.T) {
	input := "#123 @ 3,2: 5x4"
	expected := Claim{
		"id":     123,
		"left":   3,
		"top":    2,
		"width":  5,
		"height": 4,
	}

	if claim := ParseClaim(input); !reflect.DeepEqual(expected, claim) {
		t.Errorf("Expected %v to deeply equal %v", expected, claim)
	}
}

func TestOverlap(t *testing.T) {
	inputs := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}

	expected := 4
	claims := make([]Claim, 0)

	for _, input := range inputs {
		claims = append(claims, ParseClaim(input))
	}

	if overlap := GetOverlap(claims); overlap != expected {
		t.Errorf("Expected %v to equal %v", overlap, expected)
	}
}

func TestOverlapsWith(t *testing.T) {
	testCases := []struct {
		a        string
		b        string
		overlaps bool
	}{
		{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", true},
		{"#1 @ 1,3: 4x4", "#3 @ 5,5: 2x2", false},
		{"#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2", false},
	}

	for _, tc := range testCases {
		a := ParseClaim(tc.a)
		b := ParseClaim(tc.b)

		if overlaps := a.OverlapsWith(b); overlaps != tc.overlaps {
			t.Errorf("Expected overlap of %v with %v to be %v", tc.a, tc.b, tc.overlaps)
		}
	}
}
