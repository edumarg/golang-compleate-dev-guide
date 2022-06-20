package main

// package in go is a project or workspace. There are executable and reusable packages.
// Using the word main we let know that is a executable package. Any other name is for reusable package.
// Main package should have a function called main.

import "fmt"

// import is used to import other packages to have access to code written on them and use it in my project.
// This could be go standard libraries or packages build by me or any other person.

func main() {
	fmt.Println("hello world!!!")
}

// To run a project on CLI type: go run <name>.go. For example in this case go run main.go
// To build a project on CLI type: go build <name>.go. For example in this case go build main.go.
// This will create a exe file you can run in the CLI with ./<name.exe> For example in this case ./main.exe
