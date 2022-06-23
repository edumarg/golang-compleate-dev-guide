package main

import (
	"fmt"
	"net/http"
	"time"
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
		// fmt.Println(<-c)
	}

	// for i := 0; i < len(URLs); i++ {
	// 	// fmt.Println(<-c)
	// } // Will get the channel values according to the number of links

	for l := range c {
		go func(link string) {
			time.Sleep(3 * time.Second)
			checkURL(link, c)
		}(l) // function literal

	} // infinite loop
}

func checkURL(URL string, c chan string) {
	_, err := http.Get(URL)
	if err != nil {
		fmt.Println(URL, "might be down!")
		c <- URL
		return
	}
	fmt.Println(URL, "is up!")
	c <- URL
}
