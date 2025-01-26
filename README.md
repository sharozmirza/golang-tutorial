### How to run a go file?

If the file is has the main function/package in it then:<br>
`$ go run <filename>.go`

### How to run a go application?

Go to the project directory and run: `$ go build`<br>
This will create an executable file of the application and it can be run in any platform even if the platform does not have go installed.

### Go project and modules

If running `$ go build` shows error `go: cannot find main module`, that means that there is no go project in the directory. To make a folder a go project, run `$ go mod init <module-name>` in that folder. For example, `go mod init example.com/my-app`. It will add `go.mod` file within the directory, which means the folder and the subfolders are now part of a go module. Now, the `$ go build` command should work and it should create `my-app.exe` file in the directory. On Windows computers this file can be run by double clicking it and on Mac/Linux, it can be run by `$ ./my-app` command.

A module can also be run using the command `$ go run .` from the module directory.

### Value types

Go is a statically typed language. So value types are important in Go.

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

```Go
var age int     // age is 0
```

Here's a list of the novels for the different types:

* `int`: `0`
* `float`: `0.0`
* `string`: `""`
* `bool`: `false`
