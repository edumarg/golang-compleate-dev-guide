package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	fmt.Println("response status: ", resp.StatusCode)
	io.Copy(lw, resp.Body)

}

// custom writer using writer interface https://pkg.go.dev/io@go1.18.3#Writer
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("just wrote this many bytes", len(bs))
	return len(bs), nil
}
