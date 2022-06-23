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
	getLanguage() string
}
type englishBot struct {
	version  string
	language string
}
type spanishBot struct {
	version  string
	language string
}

func main() {
	eb := englishBot{version: "1.3.5", language: "English"}
	sb := spanishBot{version: "1.2.0", language: "Español"}
	printLanguage(eb)
	printGreeting(eb)
	printVersion(eb)
	printLanguage(sb)
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
	// custom logic to get english version
	return eb.version
}

func (sb spanishBot) getVersion() string {
	// custom logic to get spanish version
	return sb.version
}

func (eb englishBot) getLanguage() string {
	// custom logic to get english language
	return eb.language
}

func (sb spanishBot) getLanguage() string {
	// custom logic to get spanish language
	return sb.language
}

func printGreeting(tb templateBot) {
	fmt.Println("bot version", tb.getGreeting("Mike"))
}

func printVersion(tb templateBot) {
	fmt.Println("bot version", tb.getVersion())
}

func printLanguage(tb templateBot) {
	fmt.Println("bot Language", tb.getLanguage())
}
