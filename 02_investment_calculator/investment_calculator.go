package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// Scan() can't (easily) fetch multi-word input values
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculatedFutureValues(investmentAmount, expectedReturnRate, years)

	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (adjusted for inflation):", futureRealValue)

	/*
	 * More info about the formatting options can be found here: https://pkg.go.dev/fmt@go1.23.5
	 */

	// Same output can be achieved by using Printf() function
	// fmt.Printf("Future Value: %v\nFuture Value (adjusted for inflation): %v\n", futureValue, futureRealValue)

	// It will print a shorter float number
	// fmt.Printf("Future Value: %f\nFuture Value (adjusted for inflation): %f\n", futureValue, futureRealValue)

	// %.0f will remove the decimal point and round the value
	// fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for inflation): %.1f\n", futureValue, futureRealValue)

	// Using backticks for line breaking
	// fmt.Printf(`Future Value: %.1f
	// Future Value (adjusted for inflation): %.1f`, futureValue, futureRealValue)
	// fmt.Println()

	// Use of Sprintf() function
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func calculatedFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}
