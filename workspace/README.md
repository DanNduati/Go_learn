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