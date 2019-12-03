package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dotariel/advent-of-go/exercise"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	day, err := strconv.ParseInt(os.Args[1], 10, 8)
	if err != nil {
		usage()
		return
	}

	if err := exercise.Run(int(day)); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func usage() {
	fmt.Println("usage: advent-of-go <day>")
}
