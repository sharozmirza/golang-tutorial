/*
* There must be at least one package clause in your program, but you can have multiple packages
* A single package can also be split into multiple files
* The idea behind the packages is to organize code
* If there are mulple packages we can use features from Package A to Package B for example
* We can export and import packages
 */

// main is a special package name that tells go that
// this package will be the main entry point of the application
package main

/*
* fmt (format) package, it is part of go standard library
* Here are the list of pakcages that come with go standard library: https://pkg.go.dev/std
 */
import "fmt"

// This function must be named main(), since this is where the code execution starts
// If you are building an executable program a main() is required in that application
// There can be two files which use the main package, but there cannot be two main() functions in an app
// If you develop a 3P library or a package like fmt, a main() is not required
// The reason for this is this package is not meant to be executed as a program
// This is just imported into another application/module/project to use the functions from that package in your program
func main() {
	// String must use double quotes ("") or ``
	fmt.Print("Hello World!")   // prints % at the end of the line
	fmt.Println("Hello World!") // prints no % at the end of the line
}
