package main

import "fmt"

func main() {
	fmt.Println("Struct study")
	ajmal := User{"Ajmal", 21, "m.pajmalaju07@gmail.com", true}
	// fmt.Println(ajmal)
	// detaid output
	fmt.Printf("Details: %+v\n", ajmal)
	//perticular field
	fmt.Printf("Name: %v, Email %v\n", ajmal.Name, ajmal.Email)
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}
