# Go Type Assertion Panic
This repository demonstrates a common error in Go: type assertion panics.  The example code attempts to assert an interface{} value of the wrong type into an integer. This will cause the program to crash. The solution shows how to handle this using type switches and error checking.  

## How to Run
1. Clone this repository.
2. Navigate to the repository directory using your terminal.
3. Run the `bug.go` file using `go run bug.go`. Observe the panic.
4. Run the `bugSolution.go` file using `go run bugSolution.go`. Observe the correct handling.

## License
MIT