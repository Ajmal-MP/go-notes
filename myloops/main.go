package main

import "fmt"

func main() {
	fmt.Println("golang loop study")
	days := []string{"sunday", "monday", "thuesday", "wednesday", "friday", "saturday"}
	//normal looping
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	// using rage
	fmt.Println("")
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println("")

	// using ok ; syntax
	for index, value := range days {
		fmt.Printf("day[%v] = %v \n", index, value)
	}

	sample()
}

func sample() {
	fmt.Println("looping based on condition")
	num := 1
	for num < 10 {
		if num == 7 {
			break
		} else if num == 5 {
			num++
			continue
			// labeling sample
		} else if num == 3 {
			goto samplelabel
		}

		fmt.Println("Number is: ", num)
		num++
	}
	//creating label
samplelabel:
	fmt.Println("this it the lable created for 3")
}
