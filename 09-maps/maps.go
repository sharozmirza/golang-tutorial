package main

import "fmt"

// declare a type alias for a map since declaring a map is too long
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

	// nil map
	var mm map[string]int
	fmt.Println(mm == nil)

	// initialize mm
	mm = map[string]int{"apple": 1, "banana": 2}
	fmt.Println(mm)

	// create a map using type alias
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	courseRatings.output()
}
