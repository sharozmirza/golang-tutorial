package main

import (
	"fmt"
	"math"
)

func main() {

	// Multuple variables can be declared and assigned values to them in the same line. For example:
	// 		var investmentAmount, name = 1000, "example-name"

	// If an explicit type assignment is used like below then all of them must be the same type
	var investmentAmount, years float64 = 1000, 10

	// Recommendation is to use := as much as possible to declare and assign a variable
	// Type would be identified by go automatically
	expectedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println(futureValue)
}
