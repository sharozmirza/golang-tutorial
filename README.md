## Table of Contents

* [Go Project and Modules](#go-project-and-modules)
* [Go Packages](#go-packages)
* [Value Types](#value-types)
  * [Basic Types](#basic-types)
  * [Null Values](#null-values)
* [Variable Declarations](#variable-declarations)
* [Functions](#functions)
  * [Using Functions As Values](#using-functions-as-values)
  * [Variadic Functions](#variadic-functions)
* [Conditional Statements](#conditional-statements)
* [Loops](#loops)
* [Error Handling](#error-handling)
* [Pointers](#pointers)
* [Structs](#structs)
  * [Using `structs` with Pointers](#using-structs-with-pointers)
  * [Structs Methods](#structs-methods)
  * [Constructor Functions](#constuctor-functions)
  * [Struct Embedding](#struct-embedding)
  * [Type Alias and Named Type](#type-alias-and-named-type)
  * [Struct Tags](#struct-tags)
* [Interfaces](#interfaces)
  * [Embedded Interface](#embedded-interface)
  * [Empty Interface](#empty-interface)
* [Generics](#generics)
* [Arrays](#arrays)
  * [Accessing Array Elements](#accessing-array-elements)
  * [Iterating Over An Array](#iterating-over-an-array)
  * [Multidimensional Arrays](#multidimensional-arrays)
  * [Array Slicing](#array-slicing)
  * [Length and Capacity of a Slice](#length-and-capacity-of-a-slice)
  * [Dynamic Arrays (Lists) With Slices](#dynamic-arrays-lists-with-slices)
  * [Removing Elements Using Slices](#removing-elements-using-slices)
* [Maps](#maps)
  * [`make()`](#make)

### How to run a go file?

If the file is has the main function/package in it then:<br>
`$ go run <filename>.go`

### How to run a go application?

Go to the project directory and run: `$ go build`<br>
This will create an executable file of the application and it can be run in any platform even if the platform does not have go installed.

### Go Project and Modules

If running `$ go build` shows error `go: cannot find main module`, that means that there is no go project in the directory. To make a folder a go project, run `$ go mod init <module-name>` in that folder. For example, `go mod init example.com/my-app`. It will add `go.mod` file within the directory, which means the folder and the subfolders are now part of a go module. Now, the `$ go build` command should work and it should create `my-app.exe` file in the directory. On Windows computers this file can be run by double clicking it and on Mac/Linux, it can be run by `$ ./my-app` command.

A module can also be run using the command `$ go run .` from the module directory.

## Go Packages

In Go, packages are a fundamental way to organize and structure code. A package is a collection of Go source files in the same directory, and it is the basic unit of code organization and reuse in Go.

* Go comes with a large set of standard library packages, which provide common functionality like input/output (I/O), string manipulation, networking, and more. For example, `fmt`, `math`, `os`, `net/http`, and `encoding/json`.
* Every Go file must be a part of a package. That's how all the files are linked together behind the scene by the Go compiler.
* Files in the same directory of the `main` package will be part of the `main` package. The fuctions from these files can be called within the other files in the `main` package without using the import statement. The first letter of the function names does not need to be capitalized in this case.
* Every package must go in a subfolder within a project. The folder will have the same name as the package.
* To import your own package in an application, the full path that includes the module path will need to be added.". For example, if the module name is `example.com/my-app`, then to import a package, `my-package`, into another file of the application, add `example.com/my-app/my-package` to the `import` statement.
* In Go, there is no "export" keyword. The functions, constants, variables that start with an upper case character are available in other packages.
* To install a 3P package, use `go get <path-to-the-package>` command. For example, `go get github.com/Pallinder/go-randomdata`. Then the package will be downloaded and added to the project. The downloaded package is stored globally in GOPATH on the system. Also, `require <path-to-the-pakage>:<package-version>` will be added to the `go.mod` file. `go.mod` file lists all the 3P dependencies of a project. If a project is cloned from a repository, use `go get` command to download all the dependencies listed in the `go.mod` file of that project.
* To find Go packages, go to: https://pkg.go.dev/

## Value Types

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

### Null Values

All Go value types come with a so-called **"null value"** which is the value stored in a variable if no other value is explicitly set. For example, the following `int` variable would have a default value of `0` (because `0` is the null value of `int`, `int32`).

```go
var age int     // age is 0
```

Here's a list of the null values for different types:

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

If multiple values are returned by a function and one of the values is not wanted/handled, `_` can be used instead of keeping the variable unused. For example:

```go
func myFunction(param1 int, param2 string) (retVal1 float64, retVal2 int) {
    // do something
    return
}

val1, _ := myFunction(arg1, arg2)
```

### Using Functions As Values

In Go, functions can be assigned to variables, passed as arguments, and returned from other functions. This is possible because functions in Go are first-class citizens, meaning they can be treated as values just like any other type (e.g., integers, strings).

An example where a function is assigned to a variable and invoked:

```go
// Define a function
func add(a, b int) int {
    return a + b
}

func main() {
    var funcVar func(int, int) int  // Assign the function to a variable
    funcVar = add // Assign the add function to funcVar

    // Call the function through the variable
    result := funcVar(3, 4)
    fmt.Println("Result:", result) // Output: Result: 7
}
```

Example of passing a function as an argument:

```go
// Define a function that takes another function as an argument
func operate(a, b int, op func(int, int) int) int {
    return op(a, b)
}

// Define a function to add
func add(a, b int) int {
    return a + b
}

// Define a function to multiply
func multiply(a, b int) int {
    return a * b
}

func main() {
    // Call operate with different functions
    fmt.Println("Addition:", operate(3, 4, add))       // Output: Addition: 7
    fmt.Println("Multiplication:", operate(3, 4, multiply)) // Output: Multiplication: 12
}
```

Example of returning a function from another function using an **anonymous function**:

* An **anonymous function** is a function that is defined without a name. In Go, anonymous functions can be created just like named functions, but they are often used where the function is needed only temporarily and doesn’t need a name. They are often used in callbacks, or as arguments to other functions.
* A **closure** is a function that captures variables from its outer function and retains access to those variables even after the function that created them has finished executing. In Go, closures are created when an anonymous function is defined within another function, and it captures the variables from the outer function's scope.
  * The function `multiplier(factor int)` in the example below returns a closure that multiplies its argument `x` by the `factor` captured from the outer fucntion's scope. Two closures, `double` and `triple`, are created, capturing `2` and `3` as the `factor`, respectively. The variable `factor` is defined in the outer function's scope, and the closures retain access to it even after `multiplier()` has finished executing. That's why calling `double(5)` multiplies 5 by 2, and `triple(5)` multiplies 5 by 3.

```go
// Define a function that returns another function
func multiplier(factor int) func(int) int {
    // Returns an anonymous function (closure)
    return func(x int) int {
        return x * factor
    }
}

func main() {
    // Get a function that multiplies by 2
    double := multiplier(2)
    // Call the returned function
    fmt.Println(double(5)) // Output: 10

    // Get a function that multiplies by 3
    triple := multiplier(3)
    // Call the returned function
    fmt.Println(triple(5)) // Output: 15
}
```

### Variadic Functions

A variadic function is a function that accepts a variable number of arguments of the same type. Instead of defining a fixed number of parameters, variadic functions use an ellipsis (...) before the type to indicate that they can take zero or more arguments.

```go
// A variadic function to calculate the sum of numbers
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Variadic function with a regular argument and a variadic argument
func greet(prefix string, words ...string) {
    fmt.Println(prefix)
    for _, word := range words {
        fmt.Println(word)
    }
}

func main() {
    fmt.Println(sum(10, 20, 30, 40))   // Output: 100

    greet("Hi:", "How", "are", "you")
    // Output:
    // Hi:
    // How
    // are
    // you
}
```

A slice can also be passed to a variadic function by using the spread operator (`...`), which unpacks the slice into individual arguments.

```go
func printNumbers(numbers ...int) {
    for _, num := range numbers {
        fmt.Println(num)
    }
}

func main() {
    nums := []int{1, 2, 3, 4, 5}

    // Passing a slice to the variadic function
    printNumbers(nums...)  // Output: 1 2 3 4 5
}
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

## Error Handling

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

Pointers are variables that store addresses instead of values. Using pointers making unnecessary copies of a value can be avoided. It can also directly mutate values by accessing the address of a variable.

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

### Using `structs` with Pointers

Normally dereferencing a pointer (e.g. `(*p).FirstName = "Jane"`) is needed to get access to the value stored in it. But for structs, Go allows the usage of pointers directly to access the value stored in it.

```go
func updateName(p *Person) {
    p.FirstName = "Jane" // Modifies the original struct
}

func main() {
    person := Person{"John", "Doe"}
    updateName(&person) // Passing a pointer to the struct

    fmt.Println(person.FirstName) // Output: Jane
}
```

### Structs Methods

A function that is associated with a struct is called a method. A struct method is defined outside of the struct literals (`{}`)

```go
// A struct method of the Person struct. Since it is a struct method, it doesn't have any parameter related to the struct. It has a receiver type (p Person) instead
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.FirstName, p.LastName)
}

func main() {
    person := Person{"John", "Doe"}
    person.Greet() // Output: Hello, my name is John Doe
}
```

A mutator method (a setter method) is a method that modifies the fields of an object or struct. In Go, since function arguments are passed by value by default, when an instance of a struct is passed to a function or method, Go creates a copy of that instance, and changes are made to the copy. Therefore, it does not affect the original instance. By passing a pointer to the instance, the method modifies with the original instance, not a copy of it.

```go
// Mutator method with a pointer receiver (pass-by-reference)
func (p *Person) SetFirstName(firstName string) {
    p.FirstName = firstName
}
```

### Constuctor Functions

In Go, there isn't a built-in concept of a constructor like in some object-oriented languages (e.g., Python or Java). However, a constructor for a struct can be simulated by creating a function that returns an instance of the struct, often initialized with specific values.

```go
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
}
```

A constructor function can also be used for data validation before instantiating a struct.

```go
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
}
```

### Struct Embedding

Struct embedding in Go is a feature that allows one struct to be embedded inside another. This provides a way to achieve inheritance-like behavior without explicitly using classes or inheritance, as in traditional object-oriented programming languages. When a struct is embedded in another struct, the fields and methods of the embedded struct can be accessed directly from the outer struct.

Simple struct embedding:

```go
// Define a struct for Address
type Address struct {
    Street, City, State string
}

// Define a method on Address
func (a Address) FullAddress() string {
    return a.Street + ", " + a.City + ", " + a.State
}

// Define a struct for Person that embeds Address
type Person struct {
    FirstName string
    LastName  string
    Address   // Embedded struct
}

// Method on Person that overrides the inherited FullAddress method
func (p Person) FullAddress() string {
    return p.FirstName + " " + p.LastName + " lives at " + p.Street + ", " + p.City + ", " + p.State
}

func main() {
    // Create an instance of Person, including an Address
    person := Person{
        FirstName: "John",
        LastName:  "Doe",
        Address:   Address{"123 Main St", "Anytown", "CA"},
    }

    // Access fields or methods directly from the embedded Address struct
    fmt.Println("Name:", person.FirstName, person.LastName)
    fmt.Println("Street:", person.Street) // Accessing Address directly
    fmt.Println("City:", person.City)     // Accessing Address directly
    fmt.Println("State:", person.State)   // Accessing Address directly
    fmt.Println(person.FullAddress())     // Output: 123 Main St, Anytown, CA

    // Call the method on the Person struct (which overrides the Address method)
    fmt.Println(person.FullAddress())     // Output: John Doe lives at 123 Main St, Anytown, CA
}
```

### Type Alias and Named Type

In Go, the `type` keyword can be used to create a **type alias**. It does not create a new, distinct type. It's simply an alias for the original type. An alias provides an alternate name for an existing type, which can be useful for clarity, readability, or creating custom types that are based on existing ones.

```go
// Create an alias for the built-in int type
type MyInt = int

func main() {
    var a MyInt = 10
    var b int = 20

    fmt.Println(a)  // Output: 10
    fmt.Println(b)  // Output: 20
}
```

If a completely new type is needed to be defined based on an existing type (not just an alias), `type` can be used *without* the `=` sign. This creates a **distinct, named type**, which behaves like the original type but is treated as a separate type. This is useful when a method needs to associated with a built-in variable type (or a type that is imported from another package) like associatind a methnod with a `struct` type.

```go
// Create a named type that is based on int
type MyInt int

// Define a method on a built-in type (int)
// Error: cannot define new methods on non-local type
// func (m int) MultiplyByTwo() int {
//     return m * 2
// }

// Define a method on the new named type (MyInt)
func (m MyInt) MultiplyByTwo() MyInt {
    return m * 2
}

func main() {
    var a MyInt = 10    // to use custom type, the value needs to be set explicitly
    b := 20

    fmt.Println(a)  // Output: 10
    fmt.Println(b)  // Output: 20

    // The types are distinct, so this would result in an error
    // a = b  // Error: cannot use b (type int) as type MyInt in assignment

    // Call the method on MyInt
    result := a.MultiplyByTwo()
    fmt.Println("Result:", result) // Output: Result: 20
}
```

### Struct Tags

Struct tags are metadata annotations that can be associated with the fields of a struct. They allow additional information to be attached to the struct fields, which can be used by external libraries or frameworks to perform tasks like JSON serialization, databse ORM mapping, validation, or reflection.

Example: JSON serialization

```go
type Person struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Age       int    `json:"age"`
}

func main() {
    // Create an instance of Person
    p := Person{FirstName: "John", LastName: "Doe", Age: 30}

    // Convert the struct to JSON
    jsonData, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(jsonData)) // Output: {"first_name":"John","last_name":"Doe","age":30}
}
```

Example: field validation

```go
import "github.com/go-playground/validator/v10"

type Person struct {
    FirstName string `validate:"required"`
    Age       int    `validate:"gte=18"`
}

func main() {
    validate := validator.New()

    p := Person{FirstName: "", Age: 17}

    // Validate struct
    err := validate.Struct(p)
    if err != nil {
        fmt.Println(err) // Output: FirstName: required
    }
}
```

## Interfaces

In Go, an `interface` is a type that specifies a set of method signatures without providing the implementation of those methods. The method signatures can be implemented by any type, regardless of their concrete definitions. Interfaces are central to Go’s design and are key to achieving polymorphism and abstraction. There’s no need for explicit declarations like `implements` or `extends` to implement an interface.

Syntax for defining an interface:

```go
type InterfaceName interface {
    Method1() returnType
    Method2(argType) returnType
    // More methods
}
```

Example:

```go
// Define an interface
type Speaker interface {
    Speak() string
}

// Define a struct that implements the interface
type Person struct {
    Name string
}

// Implement the Speak method for Person
func (p Person) Speak() string {
    return "Hello, my name is " + p.Name
}

// Another type implementing the interface
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return "I am an animal named " + a.Name
}

func main() {
    // Create instances of Person and Animal
    person := Person{Name: "John"}
    animal := Animal{Name: "Dog"}

    // Use the Speak method
    fmt.Println(person.Speak()) // Output: Hello, my name is John
    fmt.Println(animal.Speak()) // Output: I am an animal named Dog
}
```

### Embedded Interface

An embedded interface in Go is an interface type that is included (embedded) within another interface. When an interface embeds another interface, it inherits all the methods of the embedded interface, which allows for more flexible and modular designs.

```go
// Define a basic interface with a method
type Speaker interface {
    Speak() string
}

// Define a more complex interface that embeds the Speaker interface
type Greeter interface {
    Speaker  // Embedding the Speaker interface
    Greet() string
}
```

### Empty Interface

An empty interface (`interface{}`) is a type that holds any value. Since every type in Go implicitly satisfies the empty interface, it is often used to store values of any type and can be used for type assertions.

```go
var i interface{} = 42

// Type assertion
value, ok := i.(int)
if ok {
    fmt.Println("The value is an int:", value) // Output: The value is an int: 42
} else {
    fmt.Println("The value is not an int")
}
```

Type assertion with switch statement:

```go
func printType(value interface{}) {
    switch value.(type) {
    case int:
        fmt.Println("Integer:", value)
    case string:
        fmt.Println("String:", value)
    case float64:
        fmt.Println("Float64:", value)
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
    printType(42)        // Output: Integer: 42
    printType("hello")   // Output: String: hello
    printType(3.14)      // Output: Float64: 3.14
    printType(true)      // Output: Unknown type
}
```

## Generics

Generics allows a flexible way to write functions, types, and data structures that can work with any data type.

```go
// Define a generic function to add two values of type int or float64 or string
func add[T int | float64 | string](a, b T) T {
    return a + b
}

// Define a generic function to swap two values of any type
func Swap[T any](a, b T) (T, T) {
    return b, a
}

func main() {
    // Adding two integers
    result := add(1, 2)
    fmt.Println("Result:", result)// Output: Result: 3

    // Adding two float64s
    result := add(1.3, 2.5)
    fmt.Println("Result:", result)// Output: Result: 3.8

    // Swapping two integers
    x, y := 1, 2
    x, y = Swap(x, y)
    fmt.Println("Swapped integers:", x, y) // Output: Swapped integers: 2 1

    // Swapping two strings
    str1, str2 := "hello", "world"
    str1, str2 = Swap(str1, str2)
    fmt.Println("Swapped strings:", str1, str2) // Output: Swapped strings: world hello

    // Swapping two float64 values
    f1, f2 := 3.14, 2.71
    f1, f2 = Swap(f1, f2)
    fmt.Println("Swapped floats:", f1, f2) // Output: Swapped floats: 2.71 3.14
}
```

## Arrays

In Go, an array is a fixed-size collection of elements of the same type. The size of an array is part of its type, meaning `[5]int` and `[10]int` are distinct types. Arrays are value types, so when an array is assigned or passed, a copy of the array is made.

```go
// Declare an array of 5 integers
var numbers [5]int

// Initialize the array
numbers = [5]int{1, 2, 3, 4, 5}

// Declare and initialize an array of 3 strings
fruits := [3]string{"Apple", "Banana", "Cherry"}

fmt.Println(numbers) // Output: [1 2 3 4 5]
fmt.Println(fruits)  // Output: [Apple Banana Cherry]
```

### Accessing Array Elements

```go
numbers := [4]int{10, 20, 30, 40}

// Access and print the third element (index 2)
fmt.Println(numbers[2]) // Output: 30

// Modify the second element (index 1)
numbers[1] = 25
fmt.Println(numbers) // Output: [10 25 30 40]
```

### Iterating Over An Array

```go
numbers := [3]int{1, 2, 3}

// Iterate over the array and print each element
for i := 0; i < len(numbers); i++ {
    fmt.Printf("Element at index %d: %d\n", i, numbers[i])
}

// Output
// Element at index 0: 1
// Element at index 1: 2
// Element at index 2: 3

fruits := [3]string{"Apple", "Banana", "Cherry"}

// Iterate over the array using range
for index, value := range fruits {
    fmt.Printf("Index: %d, Value: %s\n", index, value)
}

// Output
// Index: 0, Value: Apple
// Index: 1, Value: Banana
// Index: 2, Value: Cherry
```

### Multidimensional Arrays

```go
matrix = [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Print the 2D array
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        fmt.Printf("%d ", matrix[i][j])
    }
    fmt.Println()
}

// Output
// 1 2 3 
// 4 5 6 
// 7 8 9 
```

### Array Slicing

A slice of an array is created by specifying a range of indices within the array. The syntax for array slicing is as follows:

```go
array[start:end]
```

* `start`: The index at which to begin the slice (inclusive).
* `end`: The index at which to end the slice (exclusive).
* If `start` is omitted, it defaults to `0` (the beginning of the array).
* If `end` is omitted, it defaults to the length of the array.
* Negative indexes can't be used
* Highest index that can be used is the array size
* When a slice is created, it points to the underlying array. The slice doesn't store its data or creates a copy of the original array; it simply references part of the array. Therefore, modifying a slice will modify the original array (and vice versa).

### Length and Capacity of a Slice

* **Length** refers to the number of elements in the slice. It represents the current size of the slice.
* **Capacity** refers to the maximum number of elements the slice can hold without allocating additional memory. It represents the total size of the underlying array starting from the first element of the slice.
* A slice can access elements beyond its length (up to its capacity) because it shares the same underlying array, but it cannot access elements before its start index in the original array.

```go
arr := [6]int{10, 20, 30, 40, 50, 60}

fmt.Println("Original Array:", arr) // Output: [10 20 30 40 50 60]

// Create a slice from the array (from index 1 to 4, exclusive)
slice1 := arr[1:4]

fmt.Println("Slice 1 (arr[1:4]):", slice1) // Output: [20 30 40]
fmt.Println("Length of Slice 1:", len(slice1)) // Output: 3
fmt.Println("Capacity of Slice 1:", cap(slice1)) // Output: 5 (from index 1 to the end of the array)

// Create another slice from the first slice (from index 1 to 3, exclusive)
slice2 := slice1[1:3]

fmt.Println("Slice 2 (slice1[1:3]):", slice2) // Output: [30 40]
fmt.Println("Length of Slice 2:", len(slice2)) // Output: 2
fmt.Println("Capacity of Slice 2:", cap(slice2)) // Output: 4 (from index 2 to the end of the array)

// Access elements beyond the length of slice2 (using the underlying array)
// Explanation: slice2 has access to the underlying array's elements beyond its length.
fmt.Println("Accessing slice2[2]:", slice2[2]) // Output: 50

// Modify an element in slice2 and observe the change in the original array
slice2[1] = 99
fmt.Println("After modifying slice2[1]:")
fmt.Println("Slice 2:", slice2) // Output: [30 99]
fmt.Println("Slice 1:", slice1) // Output: [20 30 99]
fmt.Println("Original Array:", arr) // Output: [10 20 30 99 50 60]
```

### Dynamic Arrays (Lists) With Slices

In Go, there is no concept of a *dynamic array* or *list* in the same way as some other languages. However, dynamic/resizable behavior can be achieved using slices.

```go
// Create a dynamic array (slice) with one initial element
dynamicArray := []int{10}

// Append elements to the slice (dynamic resizing)
dynamicArray = append(dynamicArray, 20, 30)

fmt.Println("Dynamic Array:", dynamicArray)     // Output: [10 20 30]
fmt.Println("Second Element:", dynamicArray[1]) // Output: 20
```

### Removing Elements Using Slices

Go does not provide a built-in function specifically for removing elements from an array/slice. However, tt can be achieved by by using slicing or other techniques.

```go
slice := []int{10, 20, 30, 40, 50}

// Remove the element at index 2 (value 30)
indexToRemove := 2
slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)

fmt.Println("Slice after removal:", slice) // Output: [10 20 40 50]
```

* The above example shows how to remove an element from a slice by index. Removing an element by value or removing multiple elements can be achieved using loops.

## Maps

A `map` is a built-in data type that allows the storage of key-value pairs, where each key is unique. The key type  must be a comparable type, meaning the key must support comparison using the `==` operator.

* Types that cannot be used as keys: slices, maps, functions, channels
* Types that can be used as keys:
  * Basic types: `int`, `float64`, `string`, `bool`, etc.
  * Pointers
  * Arrays
  * Structs: Structs can be used as keys, provided all the fields in the struct are themselves comparable. If the struct contains a field of a non-comparable type (e.g., a function, or a map), the struct cannot be used as a key.

```go
// Declation Syntax
var m map[KeyType]ValueType

// Initialize a map using map literal
m = map[string]int{"apple": 1, "banana": 2}

// Initialize a map using make
m = make(map[string]int)
```

Basic map usage:

```go
var mm map[string]int // This is a nil map
fmt.Println(mm == nil) // Output: true

// Initialize map mm
mm = make(map[string]int)

// Declare and initialize a map m
m := map[string]int{"apple": 5, "banana": 3, "cherry": 8}

// Access values
fmt.Println("Apple:", m["apple"])   // Output: Apple: 5

// Check if a key exists
value, exists := m["grape"]
if exists {
    fmt.Println("Grape:", value)
} else {
    fmt.Println("Grape not found") // Output: Grape not found
}

// Iterate over a map
for key, value := range m {
    fmt.Println(key, ":", value)
}

// Delete a key
delete(m, "banana")
```

### make()

The make() function is a built-in function used to initialize slices or maps. It allocates memory and initializes the data structure to be used, which is crucial because these types do not automatically allocate memory when declared. The `make()` function provides a way to create these types with a specified size or capacity. It can be an efficient way for saving memory if the size/capacity of the a slice or map is known before creating them.

Syntax for slices:

```go
make([]Type, length, capacity)
```

* `length`: The number of elements the slice should initially hold.
* `capacity`: (Optional parameter) The total size that the slice can grow to before a new allocation is needed. If not specified, it defaults to the same value as length.

Syntax for maps:

```go
make(map[KeyType]ValueType, initialCapacity)
```

* `initialCapacity`: (Optional parameter) The initial number of key-value pairs the map can hold before it grows. If omitted, Go uses a default capacity.

```go
// Create a slice of integers with length 5 and capacity 10
slice := make([]int, 5, 10)

fmt.Println("Slice:", slice)        // Output: Slice: [0 0 0 0 0]
fmt.Println("Length:", len(slice))  // Output: Length: 5
fmt.Println("Capacity:", cap(slice)) // Output: Capacity: 10

// Create a map with string keys and int values, with an initial capacity of 5
m := make(map[string]int, 5)
```
