package main

import (
	"fmt"
)

// interface creates a new type that has as "honorary members" to it.
// This "Honorary members" are methods(functions), that are share by different functions that may return
// different values but the logic of the function is similar or has the same objective.
type templateBot interface { // interface name
	getGreeting(string) string //<function name>(list of arguments types) list of return types
	getVersion() string
}
type englishBot struct {
	version string
}
type spanishBot struct {
	version string
}

func main() {
	eb := englishBot{version: "1.3.5"}
	sb := spanishBot{version: "1.2.0"}
	printGreeting(eb)
	printVersion(eb)
	printGreeting(sb)
	printVersion(sb)

}

func (eb englishBot) getGreeting(name string) string {
	// custom logic to get english greeting
	return "Hello " + name + ", how are you?"
}

func (sb spanishBot) getGreeting(name string) string {
	// custom logic to get spanish greeting
	return "Hola " + name + ", como estas?"
}

func (eb englishBot) getVersion() string {
	// custom logic to get english greeting
	return eb.version
}

func (sb spanishBot) getVersion() string {
	// custom logic to get spanish greeting
	return sb.version
}

func printGreeting(tb templateBot) {
	fmt.Println("bot version", tb.getGreeting("Mike"))
}

func printVersion(tb templateBot) {
	fmt.Println("bot version", tb.getVersion())
}
