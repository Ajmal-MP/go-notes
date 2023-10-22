package main

import "fmt"

func main() {
	fmt.Println("goland differ")
	//differ execute the line in end of the function
	// it's LIFO
	defer fmt.Println("Ajmal")
	defer fmt.Println("is")
	defer fmt.Println("Name")
	fmt.Println("My")
	myloop()
}

func myloop() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) //stack 0,1,2,3,4
	}
}
