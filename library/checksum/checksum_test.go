package checksum

import "testing"

func TestDiff(t *testing.T) {
	testCases := []struct {
		input    Row
		expected int
	}{
		{Row{5, 1, 9, 5}, 8},
		{Row{7, 5, 3}, 4},
		{Row{2, 4, 6, 8}, 6},
	}

	for _, tt := range testCases {
		if actual, _ := Diff(tt.input); actual != tt.expected {
			t.Errorf("row diff failed for %v; wanted %v, but got %v", actual, tt.expected, actual)
		}
	}
}

func TestFactor(t *testing.T) {
	testCases := []struct {
		input    Row
		expected int
	}{
		{Row{5, 9, 2, 8}, 4},
		{Row{9, 4, 7, 3}, 3},
		{Row{3, 8, 6, 5}, 2},
	}

	for _, tt := range testCases {
		if actual, _ := Factor(tt.input); actual != tt.expected {
			t.Errorf("row diff failed for %v; wanted %v, but got %v", actual, tt.expected, actual)
		}
	}
}

var checksumTestCases = []struct {
	input    Spreadsheet
	rowFunc  RowFunc
	expected int
}{
	{Spreadsheet([]Row{
		Row{5, 1, 9, 5},
		Row{7, 5, 3},
		Row{2, 4, 6, 8},
	}), Diff, 18},
	{Spreadsheet([]Row{
		Row{5, 9, 2, 8},
		Row{9, 4, 7, 3},
		Row{3, 8, 6, 5},
	}), Factor, 9},
}

func TestChecksum(t *testing.T) {
	for _, tt := range checksumTestCases {
		if actual := tt.input.Checksum(tt.rowFunc); actual != tt.expected {
			t.Errorf("checksum failed for %v; wanted %v, but got %v", actual, tt.expected, actual)
		}
	}
}

func TestChecksumFromInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"5 1 9 5\n7 5 3\n2 4 6 8", 18},
		{"20 150 1 2\n1 2 19\n18 75 60 8", 234},
	}

	for _, tt := range testCases {
		if actual := NewSpreadsheet(tt.input).Checksum(Diff); actual != tt.expected {
			t.Errorf("checksum from input failed for %v; wanted %v, but got %v", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkChecksumDiff(b *testing.B) {
	b.StopTimer()
	for _, tt := range checksumTestCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			tt.input.Checksum(Diff)
		}

		b.StopTimer()
	}
}

func BenchmarkChecksumFactor(b *testing.B) {
	b.StopTimer()
	for _, tt := range checksumTestCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			tt.input.Checksum(Factor)
		}

		b.StopTimer()
	}
}
