package main

import (
	"dotariel/inputreader"
	"fmt"
	"math"
)

func main() {
	requiredFuel := 0.0
	for _, val := range inputreader.ReadFloats("input.txt") {
		requiredFuel += GetTotalFuel(val, 0.0)
	}

	fmt.Printf("%f", requiredFuel)
}

// GetModuleFuel calculates the fuel required for a module
func GetModuleFuel(mass float64) float64 {
	return math.Floor(mass/3) - 2
}

// GetTotalFuel calculates the total fuel required for a module, taking the
// module fuel's mass requirement into account
func GetTotalFuel(mass float64, total float64) float64 {
	moduleFuel := GetModuleFuel(mass)

	if moduleFuel < 0 {
		return total
	}

	total += math.Max(0, moduleFuel)

	return GetTotalFuel(moduleFuel, total)
}
