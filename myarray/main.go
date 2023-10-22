package main

import (
	"fmt"
)

func main() {
	fmt.Println("welocme to goland array")
	//  entering one by one
	var fruintlist [5]string
	fruintlist[0] = "apple"
	fruintlist[1] = "orange"
	fruintlist[2] = "banana"
	fmt.Println("fruit list: ", fruintlist)

	fmt.Println("length of list: ", len(fruintlist))

	// entering all

	var veglist = [3]string{"carrot", "beans", "potato"}
	fmt.Println("veg list: ", veglist)
	fmt.Println("veg list array number", len(veglist))
}
