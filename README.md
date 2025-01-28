## Table of Contents

* [Value types](#value-types)
* [Variable Declarations](#variable-declarations)
* [Functions](#functions)

### How to run a go file?

If the file is has the main function/package in it then:<br>
`$ go run <filename>.go`

### How to run a go application?

Go to the project directory and run: `$ go build`<br>
This will create an executable file of the application and it can be run in any platform even if the platform does not have go installed.

### Go project and modules

If running `$ go build` shows error `go: cannot find main module`, that means that there is no go project in the directory. To make a folder a go project, run `$ go mod init <module-name>` in that folder. For example, `go mod init example.com/my-app`. It will add `go.mod` file within the directory, which means the folder and the subfolders are now part of a go module. Now, the `$ go build` command should work and it should create `my-app.exe` file in the directory. On Windows computers this file can be run by double clicking it and on Mac/Linux, it can be run by `$ ./my-app` command.

A module can also be run using the command `$ go run .` from the module directory.

## Value types

#### Basic Types

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

#### Null values

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