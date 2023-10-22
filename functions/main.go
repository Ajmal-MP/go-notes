package main

import "fmt"

func main() {
	fmt.Println("welcome to go functions")
	//two number sum
	sum := sum(10, 20)
	fmt.Println("sum: ", sum)

	//result of number of arguments
	result := adder(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("Total sum: ", result)

	//result of fn return more than one param
	total, message := adderWithmessage(1, 2, 3, 4, 5)
	fmt.Printf("Total is: %T and message is: %v", total, message)

}

// normal return fn
func sum(value1 int, value2 int) int {
	return value1 + value2
}

// fn with n number of arguments
func adder(values ...int) int {
	result := 0
	for _, value := range values {
		result = result + value
	}
	return result
}

//fn return more than to parameters

func adderWithmessage(values ...int) (int, string) {
	result := 0
	for _, value := range values {
		result = result + value
	}

	return result, "it's properly added"
}
