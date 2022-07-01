# Basics of multi-module workspaces in Go
## 1. Create a module for your code
1. Create a directory for your code called workspace.
```bash
$ mkdir workspace
$ cd workspace
```
2. Initialize the module

```bash
$ mkdir hello
$ cd hello
$ go mod init example.com/hello
go: creating new go.mod: module example.com/hello
```

Add a dependency on the golang.org/x/example module by using go get.
```bash
$ go get golang.org/x/example
```
Create hello.go in the hello directory
Now, run the hello program:
```bash
$ go run example.com/hello
olleH
```
## 2. Create the workspace
1. Initialize the workspace
In the workspace directory, run:
```bash
$ go work init ./hello
```
2. Run the program in the workspace directory
```bash
$ go run example.com/hello
olleH
```

## 3. Download and modify the golang.org/x/example module
1. Clone the repo
```bash
$ git clone https://go.googlesource.com/example
```
2. Add the module to the workspace
```bash
$ go work use ./example
```
this command adds a new module to the go.work file. It will now look like this:
```
go 1.18

use (
    ./hello
    ./example
)
```
The module now includes both the example.com/hello module and the golang.org/x/example module.
3. Add the new function.
Implement a new ToUpper() function to uppercase a string to the golang.org/x/example/stringutil package.
4. Modify the hello program to use the function.
## 4. Run the code in the workspace
From the workspace directory, run
```bash
$ go run example.com/hello
olleH
DANIEL
```