package main

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "https://lco.dev/"

func main() {
	fmt.Println("go web request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of the response: %T\n", response)
	defer response.Body.Close()
	// we must close the request
	// Need to read the response
	responseDatainByte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseDatainByte))
}
