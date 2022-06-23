package main

import (
	"fmt"
	"net/http"
)

func main() {
	URLs := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.amazon.com",
		"http://www.stackoverflow.com",
	}

	c := make(chan string) // creating a channel

	for _, URL := range URLs {
		go checkURL(URL, c) //go routine. go keyword use only in front of functions calls.
		fmt.Println(<-c)
	}
}

func checkURL(URL string, c chan string) {
	_, err := http.Get(URL)
	if err != nil {
		// fmt.Println(URL, "might be down!")
		c <- "Sorry " + URL + " might be down!"
		return
	}
	// fmt.Println(URL, "is up!")
	c <- "Yes " + URL + " is up!"
}
