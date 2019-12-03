package stream

import "testing"

var garbageTestCases = []struct {
	name     string
	input    string
	expected bool
}{
	{"empty garbage", `<>`, true},
	{"gargabe containing random characters", `<random characters>`, true},
	{"the extra < are ignored", `<<<<>`, true},
	{"the first > is canceled", `<{!>}>`, true},
	{"the second ! is canceled, allowing the > to terminate the garbage", `<!!>`, true},
	{"the second ! and the first > are canceled", `<!!!>>`, true},
	{"ends at the first >", `<{o"i!a,<{i<a>`, true},
}

func TestIsGarbage(t *testing.T) {
	for _, tt := range garbageTestCases {
		if _, actual, _ := Parse(tt.input); actual != tt.expected {
			t.Errorf("failed test case '%v' -> '%v'; wanted:%v, but got:%v", tt.name, tt.input, tt.expected, actual)
		}
	}
}

var parseTestCases = []struct {
	name     string
	input    string
	expected int
}{
	{"1 group", `{}`, 1},
	{"3 groups", `{{{}}}`, 3},
	{"also 3 groups", `{{},{}}`, 3},
	{"6 groups", `{{{},{},{{}}}}`, 6},
	{"1 group (which itself contains garbage)", `{<{},{},{{}}>}`, 1},
	{"also 1 group", `{<a>,<a>,<a>,<a>}`, 1},
	{"5 groups", `{{<ab>},{<ab>},{<ab>},{<ab>}}`, 5},
	{"2 groups (since all but the last '<' are canceled", `{{<!>},{<!>},{<!>},{<a>}}`, 2},
}

func TestGroupCount(t *testing.T) {
	for _, tt := range parseTestCases {
		if actual, _, _ := Parse(tt.input); len(actual) != tt.expected {
			t.Errorf("failed '%v' -> '%v'; wanted:%v, but got:%v", tt.name, tt.input, tt.expected, actual)
		}
	}
}

var scoreTestCases = []struct {
	name     string
	input    string
	expected int
}{
	{"score of 1", `{}`, 1},
	{"score of 1 + 2 + 3 = 6", `{{{}}}`, 6},
	{"score of 1 + 2 + 2 = 5", `{{},{}}`, 5},
	{"score of 1 + 2 + 3 + 3 + 3 + 4 = 16", `{{{},{},{{}}}}`, 16},
	{"score of 1", `{<a>,<a>,<a>,<a>}`, 1},
	{"score of 1 + 2 + 2 + 2 + 2 = 9", `{{<ab>},{<ab>},{<ab>},{<ab>}}`, 9},
	{"score of 1 + 2 + 2 + 2 + 2 = 9", `{{<!!>},{<!!>},{<!!>},{<!!>}}`, 9},
	{"score of 1 + 2 = 3", `{{<a!>},{<a!>},{<a!>},{<ab>}}`, 3},
}

func TestScore(t *testing.T) {
	for _, tt := range scoreTestCases {
		groups, _, _ := Parse(tt.input)
		actual := 0
		for _, group := range groups {
			actual += group
		}

		if actual != tt.expected {
			t.Errorf("failed '%v' -> '%v'; wanted:%v, but got:%v -> %v", tt.name, tt.input, tt.expected, actual, groups)
		}
	}
}

var removeGarbageTestCases = []struct {
	input    string
	expected int
}{
	{`<>`, 0},
	{`<1>`, 1},
	{`<<<<>`, 3},
	{`<{!>}>`, 2},
	{`<!!>`, 0},
	{`<!!!>>`, 0},
	{`<{o"i!a,<{i<a>`, 10},
}

func TestRemoveGarbage(t *testing.T) {
	for _, tt := range removeGarbageTestCases {
		if _, _, actual := Parse(tt.input); actual != tt.expected {
			t.Errorf("failed '%v'; wanted:%v, but got:%v", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkParse(b *testing.B) {
	b.StopTimer()
	for _, tt := range garbageTestCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			Parse(tt.input)
		}

		b.StopTimer()
	}
}
