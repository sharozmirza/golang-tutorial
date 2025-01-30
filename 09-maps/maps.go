package main

import "fmt"

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
}
