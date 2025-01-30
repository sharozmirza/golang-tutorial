# REST API

## `net/http`

The `net/http` package is part of Go's standard library and provides functions for building HTTP servers and clients. It is commonly used for creating web servers, handling HTTP requests and responses, and managing various aspects of HTTP communication.

Key Features of `net/http`:

* **HTTP Server**: It provides utilities to create and run an HTTP server.
* **HTTP Client**: It also allows making HTTP requests (GET, POST, etc.) to other servers.
* **Routing**: It offers basic routing through the `http.HandleFunc()` function, where you can map URLs to specific handler functions.
* **Request and Response Handling**: It provides structures like `http.Request` and `http.Response` for working with HTTP requests and responses.

## `github.com/gin-gonic/gin`

Gin is a web framework for Go that builds on `net/http` and provides a more feature-rich, higher-level API for building web applications and APIs. It is designed for speed and flexibility and is commonly used in Go projects requiring web services.

Key Features of Gin:

* **Routing**: Offers a more advanced and easier routing system compared to `net/http`.
* **Middleware**: Supports adding middleware functions that can be executed before or after the main handler, enabling logging, authorization, and other features.
* **JSON Handling**: Makes working with JSON more convenient, especially for building RESTful APIs.
* **Error Handling**: Simplifies error handling and returns structured responses.
* **Performance**: Gin is one of the fastest web frameworks for Go, as it is built for high-performance applications.

### `go get -u github.com/gin-gonic/gin`

* Use the above command to download and install `gin` package.
* `-u` stands for update. It ensures that all the dependencies of the package are updated. This can include transitive dependencies.
* When the above command is run, it adds the package and its dependencies to the `require` section of the `go.mod` file.
* It also updates (or creates if does not exist) the `go.sum` file.
  * The `go.sum` file is used by Go to ensure the integrity and security of dependencies in a Go project. It is automatically created and maintained by Go when you run commands like `go get`, `go build`, or `go mod tidy`. The primary purpose of the `go.sum` file is to store cryptographic hashes of the content of dependencies, helping verify that the downloaded dependencies have not been tampered with.

