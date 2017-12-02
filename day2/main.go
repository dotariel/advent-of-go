package main

import (
	"fmt"
	"io/ioutil"

	"github.com/dotariel/advent-of-go/library/checksum"
)

func main() {
	bytes, _ := ioutil.ReadFile("input")
	input := string(bytes)

	fmt.Println("Part 1: ", checksum.NewSpreadsheet(input).Checksum(checksum.Diff))
	fmt.Println("Part 2: ", checksum.NewSpreadsheet(input).Checksum(checksum.Factor))
}
