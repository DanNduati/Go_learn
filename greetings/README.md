<h1 align="center"><b>Create a Go module</b></h1>

1. Start your module using the go mod init command.
```bash
$ go mod init github.com/DanNduati/Go_learn/greetings
go: creating new go.mod: module github.com/DanNduati/Go_learn/greetings
```
2. greetings.go
- Declare the greetings package to collect related functions.
- Implement a Hello function to return the greeting.

### Error Handling
- Import the Go standard library errors package so you can use its errors.New function.
- Change the function so that it returns two values: a string and an error. Your caller will check the second value to see if an error occurred.
- Add an if statement to check for an invalid request (an empty string where the name should be) and return an error if the request is invalid. The errors.New function returns an error with your message inside.
- Add nil (meaning no error) as a second value in the successful return. That way, the caller can see that the function succeeded.

## Go Slices
Slices provide a convinient and efficient means of working with sequences of typed data. Slices are analogous to arrays in other languages but have some unusual properties. 

### Arrays
The slice type is an abstraction built on top of Go's array type.
```go
var a [4]int
a[0] = 1
i := a[0]
fmt.Println(i)
// specifying an array literal
b := [2]string{"Daniel", "Nduati"}
// slices
// Build on arrays to provide great power and convinience
// Type specification for slice is: []T, where T is the type of the elements of the slice
// A slice type has no specified length
// A slice literal is declared just like an array literal, except you leave out the element count
letters := []string{"a","b","c","d"}
// A slice can be created with the built-in function make, which is of the form
// func make([]T, len, cap) []T where T stands for the element type of the slice to be created.
var s[] byte
s = make([]byte,5,5)
//s == []byte{0,0,0,0,0}
// When the capacity of argument is omitted it defaults to the specified length
s:= make([]byte,5)
len(s) == 5
cap(s) == 5
```