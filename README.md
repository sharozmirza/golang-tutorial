## Table of Contents

* [Go Packages](#go-packages)
* [Value types](#value-types)
* [Variable Declarations](#variable-declarations)
* [Functions](#functions)
* [Loops](#loops)
* [Conditional Statements](#conditional-statements)
* [Error Handling](#error-handling)
* [Pointers](#pointers)
* [Structs](#structs)

### How to run a go file?

If the file is has the main function/package in it then:<br>
`$ go run <filename>.go`

### How to run a go application?

Go to the project directory and run: `$ go build`<br>
This will create an executable file of the application and it can be run in any platform even if the platform does not have go installed.

### Go project and modules

If running `$ go build` shows error `go: cannot find main module`, that means that there is no go project in the directory. To make a folder a go project, run `$ go mod init <module-name>` in that folder. For example, `go mod init example.com/my-app`. It will add `go.mod` file within the directory, which means the folder and the subfolders are now part of a go module. Now, the `$ go build` command should work and it should create `my-app.exe` file in the directory. On Windows computers this file can be run by double clicking it and on Mac/Linux, it can be run by `$ ./my-app` command.

A module can also be run using the command `$ go run .` from the module directory.

## Go Packages

In Go, packages are a fundamental way to organize and structure code. A package is a collection of Go source files in the same directory, and it is the basic unit of code organization and reuse in Go.

* Go comes with a large set of standard library packages, which provide common functionality like input/output (I/O), string manipulation, networking, and more. For example, `fmt`, `math`, `os`, `net/http`, and `encoding/json`.
* Every Go file must be a part of a package. That's how all the files are linked together behind the scene by the Go compiler.
* Files in the same directory of the `main` package will be part of the `main` package. The fuctions from these files can be called within the other files in the `main` package without using the import statement. The first letter of the function names does not need to be capitalized in this case.
* Every package must go in a subfolder within a project. The folder will have the same name as the package.
* To import your own package in a application, you will need to add the full path that includes the module path. For example, if the module name is `example.com/my-app`, then to import a package, `my-package`, into another file of the application, you will need to add `example.com/my-app/my-package` in the `import` statement.
* In Go, there is no "export" keyword. The functions, constants, variables that start with an upper case character are available in other packages.
* To install a 3P package, use `go get <path-to-the-package>` command. For example, `go get github.com/Pallinder/go-randomdata`. Then the package will be downloaded and added to the project. The downloaded package is stored globally in GOPATH on the system. Also, `require <path-to-the-pakage>:<package-version>` will be added to the `go.mod` file. `go.mod` file lists all the 3P dependencies of a project. If a project is cloned from a repository, use `go get` command to download all the dependencies listed in the `go.mod` file of that project.
* To find Go packages, go to: https://pkg.go.dev/

## Value types

### Basic Types

Go comes with some built-in basic types:

* `int`: A number without decimal places. Example: `-5`, `10`, `12`
* `float64`: A number with decimal places. Example: `-5.2`, `10.123`, `12.9`
* `string`: A text value wrapped by double quotes or backticks. Example: `"Hello World!"`, `` `Hi everyone` ``
* `bool`: `true` or `false`
* `uint`: Unsigned integer, which means a strictly non-negative integer. Example: `0`, `10`, `255`
* `int32`: A 32-bit signed integer, which is an integer with a specific range from -2,147,483,648 to 2,147,483,647. Example: `-1234567890`, `1234567890`
* `rune`: An alias for `int32`, represents a Unicode code point (i.e. a single character), and is used when dealing with Unicode characters
* `uint32`: A 32-bit unsigned integer, an integer that can represent values from 0 to 4,294,967,295
* `int64`: A 64-bit signed integer, an integer that can represent a larger range of values, specifically from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
* There are more types like `int8` or `uint8` which work in the same way (integer with smaller number range)

### Null values

All Go value types come with a so-called **"null value"** which is the value stored in a variable if no other value is explicitly set. For example, the following `int` variable would have a default value of `0` (because `0` is the null value of `int`, `int32`).

```go
var age int     // age is 0
```

Here's a list of the novels for the different types:

* `int`: `0`
* `float`: `0.0`
* `string`: `""`
* `bool`: `false`

## Variable Declarations

Go is a statically typed language. So value types are important in Go.

Option 1 - implicit type assignment

```go
const myConstant = 123
var myInt = 5
var myDouble = 5.5
var myString = "Hello"
var myString2 = `Word`
var myBoolean = true
var var1, var2 = 123, 123.45
```

Option 2 - explicit type assignment

```go
const myConstant float64 =  123     // go will read it as 123.0
var myInt int = -5
var myDouble float64 = 5.5
var myString string = "Hello"
var myString2 string = `Word`
var myBoolean bool = true
var var1, var2 float64 = 12, 45   // for a single line explicit type assignment all variables must be the same type
```

Option 3 - Use `:=` sign (recommended)

```go
// Recommendation is to use := as much as possible to declare and assign a variable
// Type would be identified by go automatically
myInt := 6
myDouble := 5.5
myString := "Hello"
myString2 := `Word`
myBoolean := true
var1, var2 = 12.0, 45
```

* Value formatting options can be found here: https://pkg.go.dev/fmt@go1.23.5

## Functions

A function with no parameter

```go
func myFunction() {
    // do something
}
```

A function with one parameter

```go
func myFunction(text string) {
    // do something
}
```

A function with multiple parameters of same type (option #1)

```go
func myFunction(text string, text2 string) {
    // do something
}
```

A function with multiple parameters of same type (option #2)

```go
func myFunction(text, text2 string) {
    // do something
}
```

A function with multiple parameters of different types

```go
func myFunction(number int, text string) {
    // do something
}
```

A function with single return values

```go
func myFunction(number int, text string) float64 {
    // do something
    return retVal
}
```

A function with multiple return values

```go
func myFunction(param1 int, param2 string) (float64, int) {
    // do something
    return retVal1, retVal2
}

func myFunction2(param1, param2 float64) (float64, float64) {
    // do something
    return retVal1, retVal2
}

// calling myFunction()
val1, val2 := myFunction(arg1, arg2)
```

Alternative return value syntax. In this case, `retVal1` and `retVal2` don't need to be declared within `myFunction()`

```go
func myFunction(param1 int, param2 string) (retVal1 float64, retVal2 int) {
    // do something
    return
}

// calling myFunction()
val1, val2 := myFunction(arg1, arg2)
```

If a function returns multiple values and you don't want to use/handle one of the values, `_` can be used instead of keeping the variable unused. For example:

```go
func myFunction(param1 int, param2 string) (retVal1 float64, retVal2 int) {
    // do something
    return
}

val1, _ := myFunction(arg1, arg2)
```

## Conditional Statements

If else statement:

```go
if option <= 1 {
    // do something
}

if option1 == 1 && option2 >= 2 {
    // do something
}

if option1 < 1 || option2 >= 2 {
    // do something
}

myBool := option != 1

if option1 == 1 {
    // do something
} else if option2 == 2 {
    // do something else
} else {
    // do default
}
```

Switch statement:<br>

* In other languages, a break `break` command needs to be added at the end of each case to make sure no other cases get triggered. That's not needed when working with Go. Only one case becomes active in a `switch` statement in Go.

```go
switch option {
    case 1:
        // do something
    case 2:
        // do something else
    default:
        // default behavior
}
```

## Loops

Only for loop is available in Go unlike some other languages.

```go
for i := 0; i < 10; i++ {
    //do something
}
```

Infinite loop:

```go
for {
    // do something
}
```

* Do not use `switch` statements within a loop in Go, use `if else` statement instead. If a `switch` statement is used within a loop and the `default` case has a `break` command in it, it will not break out of the loop. It will break out of the `switch` statement instead. And using `return` instead of `default` may cause unexpected behavior like exiting out of the application.

## Error handling

Error handling in Go is quite different from many other languages. Instead of using exceptions, Go uses a more explicit approach, relying on return values to indicate errors. Many functions in Go that could fail return an error as the last return value. If the function completes successfully, the error will be `nil`. If something goes wrong, the error will be non-nil. `nil` is a special value that represents the absence of a value or a "null" state.

```go
package main

import (
    "fmt"
    "errors"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

Go also provides a way to handle unexpected errors using `panic`. `panic` triggers an immediate stop to the program’s normal execution. This is typically reserved for truly exceptional situations (such as programming bugs) rather than normal error handling.

```go
func main() {
    // do something
    panic("something went wrong")
    fmt.Println("This will not be reached")
}
```

## Pointers

Pointers are variables that store addresses instead of values. Using pointers unnecessary you can avoid unnecessary value copies. It can also directly mutate values by accessing the address of a variable.

```go
myVar := 123    // the value of myVar is stored in an address in computer memory

// pointer definition and value assignment (method 1)
var myVarPointer *int
myVarPointer = &myVar   // retrieves and stores the address of myVar

// pointer definition and value assignment (method 2)
myVarPointer := &myVar  // retrieves and stores the address of myVar

fmt.Println("Address of myVar", myVarPointer)   // prints the address of myVar
fmt.Println("Value of myVar", *myVarPointer)    // pointer as a value, prints the value of myVar
```

When a variable is passed to a function, Go creates a copy of the value of that variable in memory and passes that copy to that function. So until the function execution is done and the copied value is cleaned up by Go's garbage collector, the value exists twice in memory. For very large and complex valus, this may take up too much memory space.

Another advantage of using pointers is, when a pointer of a variable is passed as a function argument, the function can directly edit the underlying value, no return value is required. But it can also make the code less understandable or cause unexpected behavior if not used properly.

For a pointer `nil` represents the absence of an address value or the pointer's null value.

## Structs

In Go, `structs` are used to define custom data types that group together variables (fields) of different types. A struct is similar to a class in other programming languages like Java or C++, but Go does not have classes or inheritance. Instead, it uses structs to represent data and behaviors associated with it.

```go
...
...

// A lower case letter can be use to name a custom type as well. In that case it would not be available outside of that package

// Define a struct 
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func main() {
    // Create an instance of the struct
    var person Person

    person = Person{    // typically no space is added between the struct name and the {
        FirstName: "John",
        LastName:  "Doe",
        Age:       30,
    }

    // Access struct fields
    fmt.Println(person.FirstName) // Output: John
    fmt.Println(person.LastName)  // Output: Doe
    fmt.Println(person.Age)       // Output: 30
}
```

Other instantiation options:

```go
// option 1 - with field names (order doesn't matter)
person := Person{
    FirstName: "John",
    LastName:  "Doe",
    Age:       30,
}

// option 2 - without field names (order matters)
person := Person{"John", "Doe", 30}

// option 3 - null struct
person := Person{}

// option 4 - partial instantiation
person := Person{
    FirstName: "John",
    Age:       30,
}
```

#### Using `structs` with pointers

Normally dereferencing a pointer (e.g. `(*p).FirstName = "Jane"`) is needed to get access to the value stored in it. But for structs, Go allows the usage of pointers directly to access the value stored in it.

```go
...
...

func updateName(p *Person) {
    p.FirstName = "Jane" // Modifies the original struct
}

func main() {
    person := Person{"John", "Doe"}
    updateName(&person) // Passing a pointer to the struct

    fmt.Println(person.FirstName) // Output: Jane
}
```

#### Structs methods

A function that is associated with a struct is called a method. A struct method is defined outside of the struct literals (`{}`)

```go
...
...

// A struct method of the Person struct. Since it is a struct method, it doesn't have any parameter related to the struct.
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.FirstName, p.LastName)
}

func main() {
    person := Person{"John", "Doe"}
    person.Greet() // Output: Hello, my name is John Doe
}
````

A mutator method (a setter method) is a method that modifies the fields of an object or struct. In Go, since function arguments are passed by value by default, when an instance of a struct is passed to a function or method, Go creates a copy of that instance, and changes are made to the copy. Therefore, it does not affect the original instance. By passing a pointer to the instance, the method modifies with the original instance, not a copy of it.

```go
...
...

// Mutator method with a pointer receiver (pass-by-reference)
func (p *Person) SetFirstName(firstName string) {
    p.FirstName = firstName
}

...
...
```

#### Constuctor functions

In Go, there isn't a built-in concept of a constructor like in some object-oriented languages (e.g., Python or Java). However, you can simulate a constructor for a struct by creating a function that returns an instance of the struct, often initialized with specific values.

```go
...
...

// Constructor function for Person
// Returning a value would have worked as well, but by returning a pointer a copy creation was prevented
func NewPerson(firstName, lastName string, age int) *Person {
    return &Person{
        FirstName: firstName,
        LastName:  lastName,
        Age:       age,
    }
}

func main() {
    // Create a new Person using the constructor
    person := NewPerson("John", "Doe", 30)

    ...
    ...
}
```

A constructor function can also be used for data validation before instantiating a struct.

```go
...
...

// Constructor function with data validation for Person
func NewPerson(firstName, lastName string, age int) (*Person, error) {
    if firstName == "" || lastName == "" {
        return nil, errors.New("first name and last name are required")
    }

    return &Person{
        FirstName: firstName,
        LastName:  lastName,
        Age:       age,
    }, nil
}

func main() {
    person, err := NewPerson("John", "Doe", 30)

    if err != nil {
        fmt.Println(err)
        return 
    }

    ...
    ...
}
```
