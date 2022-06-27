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

## Maps
A go map type is of the form:
```go
map[KeyType]ValueType
```
where KeyType may be any type that is comparable (more on this later), and ValueType may be any type at all, including another map! Example: 
```go
var m map[string]int
```
This variable m is a map of string keys to int values

Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn’t point to an initialized map. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don’t do that. To initialize a map, use the built in make function:

```go
m = make(map[string]int)
```
The make function allocates and initializes a hash map data structure and returns a map value that points to it.

### Working with maps
Go provides a familiar syntax for working with maps. This statement sets the key "route" to the value 66:
```go
m["age"] = 69
```
This statement sets the key "age" to the value 69.
```go
i := m["age"]
```
This statement retrieves the value stored under the key "route" and assigns it to a new variable i.

If the requested key doesn’t exist, we get the value type’s `zero value`. In this case the value type is `int`, so the zero value is 0:
```go
j := m["height"]
// j==0
```
The built in len function returns on the number of items in a map:
```go
n := len(m)
```
The built in delete function removes an entry from the map:
```go
delete(m, "route")
```
The delete function doesn’t return anything, and will do nothing if the specified key doesn’t exist.

A two-value assignment tests for the existence of a key:
```go
i, ok := m["age"]
```
In this statement, the first value (i) is assigned the value stored under the key "age". If that key doesn’t exist, i is the value type’s zero value (0). The second value (ok) is a `bool` that is `true` if the key exists in the map, and `false` if not.

To test for a key without retrieving the value, use an underscore in place of the first value:
```go
_,ok = m["age"]
```

To iterate over the contents of a map, use the `range` keyword
```go
for key,value := range m{
    fmt.Println("Key:",key,"Value:",value)
}
```

To initialize a map with some data use the map literal
```go
bmi := map[string]int{
    "height": 2,
    "weight": 64,
}