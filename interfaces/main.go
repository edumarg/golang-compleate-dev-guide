package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

func (eb englishBot) getGreeting() string {
	// custom logic to get english greeting
	return "Hello, how are you?"
}

func (sb spanishBot) getGreeting() string {
	// custom logic to get spanish greeting
	return "Hola, como estas?!!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
