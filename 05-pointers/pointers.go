package main

import "fmt"

func main() {
	age := 32 // regular variable

	fmt.Println("Age:", age)

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)

	// Using pointers
	// var agePointer *int // typically you don't need to add pointer to the variable name
	// agePointer = &age

	agePointer := &age
	fmt.Println("Age:", *agePointer)

	adultYears = getAdultYearsUsingPointer(agePointer) // alternatively, &age could be passed as the argument as well
	fmt.Println(adultYears)

	// using pointer for data mutation
	editAgeToAdultYears(agePointer)
	fmt.Println(age) // it will print 14

}

func getAdultYears(age int) int {
	return age - 18
}

func getAdultYearsUsingPointer(age *int) int {
	return *age - 18
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
