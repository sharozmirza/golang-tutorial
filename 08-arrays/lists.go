package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	// array slicing

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)

	fmt.Println(prices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	// building a dynamic list

	list := []float64{10.99, 8.99}
	fmt.Println(list[0:1])
	list[1] = 9.99

	list = append(list, 5.99)
	list = list[1:]
	fmt.Println(list)

	// removing elements

	// Create a slice
	slice := []int{10, 20, 30, 40, 50}

	// Remove the element at index 2 (value 30)
	indexToRemove := 2
	slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)

	// Print the updated slice
	fmt.Println("Slice after removal:", slice) // Output: [10 20 40 50]
}
