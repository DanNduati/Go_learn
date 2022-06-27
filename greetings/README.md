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