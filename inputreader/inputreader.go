package inputreader

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Iterator = func(s string)

func ReadAll(file string) string {
	bytes, _ := ioutil.ReadFile(file)

	return string(bytes)
}

func ReadStrings(file string, delim string) []string {
	vals := make([]string, 0)

	Read(file, delim, func(s string) {
		vals = append(vals, s)
	})

	return vals
}

func ReadInts(file string, delim string) []int {
	vals := make([]int, 0)

	Read(file, delim, func(s string) {
		i, _ := strconv.Atoi(s)
		vals = append(vals, i)
	})

	return vals
}

func ReadFloats(file string, delim string) []float64 {
	vals := make([]float64, 0)

	Read(file, delim, func(s string) {
		i, _ := strconv.ParseFloat(s, 64)
		vals = append(vals, i)
	})

	return vals
}

func Read(file string, delim string, iterator Iterator) {
	bytes, _ := os.ReadFile(file)

	for _, row := range strings.Split(string(bytes), delim) {
		iterator(row)
	}
}
