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
		"http://www.stackoverflow.com"}

	for _, URL := range URLs {
		checkURL(URL)
	}
}

func checkURL(URL string) {
	_, err := http.Get(URL)
	if err != nil {
		fmt.Println(URL, "might be down!")
		return
	}
	fmt.Println(URL, "is up!")
}
