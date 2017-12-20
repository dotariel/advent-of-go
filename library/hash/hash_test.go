package hash

import (
	"reflect"
	"testing"
)

var sublistTestCases = []struct {
	name         string
	position     int
	length       uint8
	expectedList []uint8
	expectedPos  []int
	error        bool
}{
	{"length larger than list size", 0, 8, nil, nil, true},
	{"starting position too big", 7, 7, nil, nil, true},
	{"starting position too small", -1, 7, nil, nil, true},
	{"simple subset from zero", 0, 2, []uint8{0, 1}, []int{0, 1}, false},
	{"simple subset from non-zero start", 1, 4, []uint8{1, 2, 3, 4}, []int{1, 2, 3, 4}, false},
	{"entire list", 0, 7, []uint8{0, 1, 2, 3, 4, 5, 6}, []int{0, 1, 2, 3, 4, 5, 6}, false},
	{"wraps around", 4, 5, []uint8{4, 5, 6, 0, 1}, []int{4, 5, 6, 0, 1}, false},
}

var sublistInput = []uint8{0, 1, 2, 3, 4, 5, 6}

func TestSublist(t *testing.T) {
	for _, tt := range sublistTestCases {
		actual, _, err := sublist(sublistInput, tt.position, tt.length)
		if !reflect.DeepEqual(actual, tt.expectedList) {
			t.Errorf("failed '%v'; wanted:%v, but got:%v", tt.name, tt.expectedList, actual)
		}

		if (err != nil) != tt.error {
			t.Errorf("failed error assertion '%v'; wanted:%v, but got:%v", tt.name, tt.error, actual)
		}
	}
}

func BenchmarkSublist(b *testing.B) {
	b.StopTimer()
	for _, tt := range sublistTestCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			sublist(sublistInput, tt.position, tt.length)
		}

		b.StopTimer()
	}
}

var reverseTestCases = []struct {
	name     string
	input    []uint8
	expected []uint8
}{
	{"nil list", []uint8(nil), []uint8(nil)},
	{"simple list", []uint8{4, 3, 2, 1}, []uint8{1, 2, 3, 4}},
	{"another simple list", []uint8{0, 0, 1, 2, 3, 0}, []uint8{0, 3, 2, 1, 0, 0}},
}

func TestReverse(t *testing.T) {
	for _, tt := range reverseTestCases {
		if actual := reverse(tt.input); !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("failed '%v'; wanted:%v, but got:%v", tt.name, tt.expected, actual)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	b.StopTimer()
	for _, tt := range reverseTestCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			reverse(tt.input)
		}

		b.StopTimer()
	}
}

var twistTestCases = []struct {
	name     string
	input    []uint8
	position int
	length   uint8
	expected []uint8
	error    bool
}{
	{"bad position", []uint8{1, 2, 3}, 4, 2, nil, true},
	{"bad length", []uint8{1, 2, 3}, 0, 4, nil, true},
	{"starting at zero", []uint8{1, 2, 3}, 0, 2, []uint8{2, 1, 3}, false},
	{"starting after zero", []uint8{1, 2, 3}, 1, 2, []uint8{1, 3, 2}, false},
	{"wrapping around", []uint8{1, 2, 3}, 2, 2, []uint8{3, 2, 1}, false},
}

func TestTwist(t *testing.T) {
	for _, tt := range twistTestCases {
		actual, err := twist(tt.input, tt.position, tt.length)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("failed '%v'; wanted:%v, but got:%v", tt.name, tt.expected, actual)
		}

		if (err != nil) != tt.error {
			t.Errorf("failed error assertion '%v'; wanted:%v, but got:%v", tt.name, tt.error, actual)
		}
	}
}

func BenchmarkTwist(b *testing.B) {
	b.StopTimer()
	for _, tt := range twistTestCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			twist(tt.input, tt.position, tt.length)
		}

		b.StopTimer()
	}
}

var hashTestCases = []struct {
	name     string
	list     []uint8
	lengths  []uint8
	expected []uint8
	error    bool
}{
	{"simple example", []uint8{0, 1, 2, 3, 4}, []uint8{3, 4, 1, 5}, []uint8{3, 4, 2, 1, 0}, true},
}

func TestHash(t *testing.T) {
	for _, tt := range hashTestCases {
		if actual := Hash(tt.list, tt.lengths); !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("failed '%v'; wanted:%v, but got:%v", tt.name, tt.expected, actual)
		}
	}
}

func BenchmarkHash(b *testing.B) {
	b.StopTimer()
	for _, tt := range hashTestCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			Hash(tt.list, tt.lengths)
		}

		b.StopTimer()
	}
}

var parseTestCases = []struct {
	input    string
	expected []uint8
}{
	{"", nil},
	{"1", []uint8{1}},
	{"4,234,1,9,15", []uint8{4, 234, 1, 9, 15}},
}

func TestParse(t *testing.T) {

	for _, tt := range parseTestCases {
		if actual := Parse(tt.input); !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("failed '%v'; wanted:%v, but got:%v", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkParse(b *testing.B) {
	b.StopTimer()
	for _, tt := range parseTestCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			Parse(tt.input)
		}

		b.StopTimer()
	}
}
