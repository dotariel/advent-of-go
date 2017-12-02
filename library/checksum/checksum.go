package checksum

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Spreadsheet represents a collection of Rows of digits
type Spreadsheet []Row

// Row represents a collection of digits
type Row []int

// Diff calculates the difference between the largest and smallest digits in
// the list
func Diff(r Row) (int, error) {
	sort.Ints(r)
	min, max := r[0], r[len(r)-1]
	return (max - min), nil
}

// Factor calculates the result of dividing the digits in the row that
// evenly divide each other.
func Factor(r Row) (int, error) {
	factors := make([]int, 0)

	for _, n := range r {
		for _, other := range r {
			if n != other && n%other == 0 {
				factors = append(factors, n)
				factors = append(factors, other)
				break
			}
		}
	}

	if len(factors) != 2 {
		return 0, fmt.Errorf("wrong number of factors found; wanted %v, but found %v", 2, len(factors))
	}

	return (factors[0] / factors[1]), nil
}

// Checksum calculates the sum of all the differences for each Row
func (sheet Spreadsheet) Checksum(fn RowFunc) (checksum int) {
	for _, row := range sheet {
		if sum, err := fn(row); err == nil {
			checksum += sum
		}
	}

	return
}

type RowFunc func(r Row) (int, error)

// NewSpreadsheet constructs a Spreadsheet from a string. The format of the input
// string must be rows separated by '\n' and values in each row separated by one
// or more spaces.
func NewSpreadsheet(input string) Spreadsheet {
	sheet := Spreadsheet{}

	for _, line := range strings.Split(input, "\n") {
		row := Row{}

		for _, number := range strings.Fields(line) {
			n, _ := strconv.ParseInt(number, 10, 64)
			row = append(row, int(n))
		}

		sheet = append(sheet, row)
	}

	return sheet
}
