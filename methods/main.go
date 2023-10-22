package main

import "fmt"

func main() {
	fmt.Println("go methods")
	ajmal := User{"ajmal", 21, "m.pajmalaju07@gmail.com"}
	ajmal.changeEmail()
	ajmal.userDetails()

}

type User struct {
	Name  string
	Age   int
	Email string
}

func (u User) userDetails() {
	fmt.Println("User Details", u)
}

func (u User) changeEmail() {
	u.Email = "newemail@gmail.com"
	fmt.Println("New details: ", u)
}
