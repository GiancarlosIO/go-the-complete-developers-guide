package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type loggWritter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	lw := loggWritter{}

	io.Copy(lw, resp.Body)
}

func (loggWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Justo wrote this many bytes:", len(bs))

	return len(bs), nil
}
