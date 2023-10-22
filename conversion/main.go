package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Mango Store")
	fmt.Println("Enter a number:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	numRating, error := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Total value is: ", numRating+1)
	}
}
