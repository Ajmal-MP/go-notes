package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("you are in and you can go")
	case 2:
		fmt.Println("you can move 2 spcace")
	case 3:
		fmt.Println("you can move 3 spcace")
		// execut case 3 and also execute case 4
		fallthrough
	case 4:
		fmt.Println("you can move 4 spcace")
	case 5:
		fmt.Println("you can move 5 spcace")
	case 6:
		fmt.Println("you can move 6 space and repeat")
	default:
		fmt.Println("what was that !")
	}

}
