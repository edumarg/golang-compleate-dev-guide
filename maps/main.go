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
		"white": "#ffffff",
	}

	fmt.Println(colors)

	// other way to declare maps
	var secondaryColors map[string]string = map[string]string{}
	fmt.Println(secondaryColors)
	secondaryColors["darkRed"] = "a60000"
	secondaryColors["darkBlue"] = "#004080"
	secondaryColors["darkGreen"] = "#008000"
	fmt.Println(secondaryColors)
	delete(secondaryColors, "darkGreen") // this is the same for all ways of declaring map
	fmt.Println(secondaryColors)

	// another way to declare maps
	primaryColors := make(map[string]string)
	fmt.Println(primaryColors)
	primaryColors["Yellow"] = "ffff00"
	primaryColors["Red"] = "ff0000"
	primaryColors["Blue"] = "0000ff"
	fmt.Println(primaryColors)

	printColors(colors)

}

// iterate over a color map
func printColors(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is:", hex)
	}
}
