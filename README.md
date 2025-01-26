### How to run a go file?
If the file is has the main function/package in it then:<br>
`$ go run <filename>.go`

### How to run a go application?
Go to the project directory and run: `$ go build`<br>
This will create an executable file of the application and it can be run in any platform even if the platform does not have go installed.

### Go project and modules
If running `$ go build` shows error `go: cannot find main module`, that means that there is no go project in the directory. To make a folder a go project, run `$ go mod init <module-name>` in that folder. For example, `go mod init example.com/my-app`. It will add `go.mod` file within the directory, which means the folder and the subfolders are now part of a go module. Now, the `$ go build` command should work and it should create `my-app.exe` file in the directory. On Windows computers this file can be run by double clicking it and on Mac/Linux, it can be run by `$ ./my-app` command.
