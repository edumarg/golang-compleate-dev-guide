package main

// map is a type that has key: values pairs.
// In a map all the different keys should be of a same type
// and all the different values assign to the keys also should all be of a same type
// Keys and values can be of different types but all keys same type and all values same type

import "fmt"

func main() {
	// creating a map with keys type string (inside braces) and values type string (outside braces)
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "00FF00",
		"blue":  "#0000FF",
	}
	fmt.Println(colors)
}
