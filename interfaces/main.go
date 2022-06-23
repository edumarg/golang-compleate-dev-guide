package main

import (
	"fmt"
)

// interface creates a new type that has as "honorary members" to it.
// This "Honorary members" are methods(functions), that are share by different functions that may return
// different values but the logic of the function is similar or has the same objective.
type templateBot interface { // interface name
	getGreeting(string) string //<function name>(list of arguments types) list of return types
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

func (eb englishBot) getGreeting(name string) string {
	// custom logic to get english greeting
	return "Hello " + name + ", how are you?"
}

func (sb spanishBot) getGreeting(name string) string {
	// custom logic to get spanish greeting
	return "Hola " + name + ", como estas?"
}

func printGreeting(tb templateBot) {
	fmt.Println(tb.getGreeting("Mike"))
}
