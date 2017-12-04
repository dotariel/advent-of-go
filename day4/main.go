package main

import (
	"fmt"
	"io/ioutil"

	"github.com/dotariel/advent-of-go/library/passphrase"
)

func main() {
	bytes, _ := ioutil.ReadFile("input")
	input := string(bytes)

	fmt.Println("Part 1: ", passphrase.CountValid(input, passphrase.Unique))
	fmt.Println("Part 2: ", passphrase.CountValid(input, passphrase.NonAnagram))
}
