package main

import "fmt"

func main() {
	var sample int = 20
	// var ptr = &sample
	fmt.Println(*&sample + 20)
}
