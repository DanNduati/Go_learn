<h1 align="center"><b>Hello</b></h1>

## Enable dependency tracking for your code.
When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a go.mod file that tracks the modules that provide those packages. That go.mod file stays with your code, including in your source code repository.

To enable dependency tracking for your code by creating a go.mod file, run the `go mod init` [command](https://go.dev/ref/mod#go-mod-init), giving it the name of the module your code will be in. The name is the module's module path.

In actual development, the module path will typically be the repository location where your source code will be kept. For example, the module path might be github.com/mymodule. If you plan to publish your module for others to use, the module path must be a location from which Go tools can download your module. For more about naming a module with a module path. [Managing dependencies](https://go.dev/doc/modules/managing-dependencies#naming_module)

```bash
$ go mod init example/hello
go: creating new go.mod: module example/hello
```
## hello.go
- Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
- Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
- Implement a main function to print a message to the console. A main function executes by default when you run the main package.

## Run the code
```bash
$ go run .
Hello, World!
```

## Call code in an external package
When you need your code to do something that might have been implemented by someone else, you can look for a package that has functions you can use in your code.

1. Visit pkg.go.dev and search for the package.
2. In your Go code, import the external package.
```go
import "rsc.io/quote"
```
3. Add new module requirements and sums.
```bash
$ go mod tidy
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
```

### Greetings
1. Import the greetings package
```go
import "github.com/DanNduati/Go_learn/greetings"
```
This is the package contained in the hello module
2. Get a greeting by calling the greetings package’s Hello function.
3. Run the `go mod tidy` command to synchronize the module's dependencies
4. Run your code:
```bash
$ go run .
Hello, Daniel. Welcome!
Don't communicate by sharing memory, share memory by communicating.
```
