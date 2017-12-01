package captcha

import "testing"

var wrapAroundTestCases = []struct {
	name     string
	input    string
	fn       AddFunc
	expected int
}{
	{"first digit matches second and third matches fourth", "1122", WrapAroundAdd, 3},
	{"each digit matches the next", "1111", WrapAroundAdd, 4},
	{"no digits match the next", "1234", WrapAroundAdd, 0},
	{"last digit matches the first wrapping around", "91212129", WrapAroundAdd, 9},
	{"list contains 4 items and all digits match digit 2 items ahead", "1212", HalfwayAroundAdd, 6},
	{"every comparison is between a 1 and a 2", "1221", HalfwayAroundAdd, 0},
	{"both 2s match each other, but no other digit has a match", "123425", HalfwayAroundAdd, 4},
}

var halfWayAroundTestCases = []struct {
	name     string
	input    string
	fn       AddFunc
	expected int
}{
	{"list contains 4 items and all digits match digit 2 items ahead", "1212", HalfwayAroundAdd, 6},
	{"every comparison is between a 1 and a 2", "1221", HalfwayAroundAdd, 0},
	{"both 2s match each other, but no other digit has a match", "123425", HalfwayAroundAdd, 4},
	{"simple case", "123123", HalfwayAroundAdd, 12},
	{"another simple case", "12131415", HalfwayAroundAdd, 4},
}

func TestInverseCaptchaWithWrapAround(t *testing.T) {
	for _, tc := range wrapAroundTestCases {
		if actual := InverseCaptcha(tc.input, tc.fn); actual != tc.expected {
			t.Errorf("`'%v' failed; wanted %v, but got %v", tc.name, tc.expected, actual)
		}
	}
}

func TestInverseCaptchaWithHalfwayAround(t *testing.T) {
	for _, tc := range halfWayAroundTestCases {
		if actual := InverseCaptcha(tc.input, tc.fn); actual != tc.expected {
			t.Errorf("`'%v' failed; wanted %v, but got %v", tc.name, tc.expected, actual)
		}
	}
}

func BenchmarkInverseCaptchaWithWrapAround(b *testing.B) {
	b.StopTimer()
	for _, tt := range wrapAroundTestCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			InverseCaptcha(tt.input, tt.fn)
		}

		b.StopTimer()
	}
}

func BenchmarkInverseCaptchaWithHalfwayAround(b *testing.B) {
	b.StopTimer()
	for _, tt := range halfWayAroundTestCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			InverseCaptcha(tt.input, tt.fn)
		}

		b.StopTimer()
	}
}
