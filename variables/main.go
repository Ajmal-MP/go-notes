package main

import "fmt"

//out side method
// not_allowed :=  30000
// var allowed int = 3000
// var allowed = 3000

// const always remain same

const token string = "hellow this is token"


func main() {
  var username string = "ajmal"
  fmt.Println(username)
  fmt.Printf("type of username: %T", username)

  var isLogged bool = true
  fmt.Println(isLogged)
  fmt.Printf("type fo isLogged: %T \n" , isLogged)

  var smallval uint8 = 255
  fmt.Println(smallval)
  fmt.Printf("type of smallval: %T \n",smallval)

  var smallfloat float32 = 123.7373737378
  fmt.Println(smallfloat)
  fmt.Printf("Type of smallfloat: %T \n", smallfloat)

  //default values and aliases

  var sampleval int
  fmt.Println(sampleval)
  fmt.Println("Type fo sampleval: %T",sampleval)

  var samplestring string
  fmt.Println(samplestring)
  fmt.Printf("type of samplestring: %T \n",samplestring)

  // implicit assing by lexer

  var implicit_string = "my name"
  fmt.Println(implicit_string)
  
  // NO war style

  numberOfusers := 30000
  fmt.Println(numberOfusers)


  fmt.Println(token)
}