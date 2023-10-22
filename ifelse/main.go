package main

import (
	"fmt"
)

func main() {
	fmt.Println("if else go study")
	mark := 10
	var result string

	if mark < 10 {
		result = "failed"
	} else if mark > 10 {
		result = "passed"
	} else {
		result = "your mark is 10 not bad"
	}
	fmt.Printf("result: %v\n", result)

	// directly giving
	if 9%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// multiple condition in one if
	if mark := 3; mark < 10 {
		fmt.Println("done")
	} else {
		fmt.Println("not done")
	}
}
