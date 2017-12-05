package main

import (
	"os"
	"strconv"

	"github.com/dotariel/advent-of-go/exercise"
)

func main() {
	day, _ := strconv.ParseInt(os.Args[1], 10, 8)
	exercise.Run(int(day))
}
