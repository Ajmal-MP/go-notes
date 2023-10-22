package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number:")

	input, _ := reader.ReadString('\n')
	fmt.Println("Enterd value is, ", input)
	fmt.Printf("Type of vlaue is %T \n", input)
}
