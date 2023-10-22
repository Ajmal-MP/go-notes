package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")
	//defining slice
	var fruitList = []string{"apple", "papaya", "banana"}
	fmt.Printf("type of fruitlist: %T \n", fruitList)
	fmt.Println(len(fruitList))

	//appending data to slice
	fruitList = append(fruitList, "orage", "shamam")
	// fmt.Println(len(fruitList))

	//appending specific slice only
	fruitList = append(fruitList[1:3])
	fmt.Println("list: ", fruitList)
	fmt.Println(len(fruitList))

	// creating slice using make

	highScore := make([]int, 4)
	highScore[0] = 10
	highScore[1] = 20
	highScore[2] = 30
	highScore[3] = 40
	// highScore[5] ==> it will throug an error

	highScore = append(highScore, 70, 60, 80)
	// sorting
	sort.Ints(highScore)
	fmt.Println("values: ", highScore)

	//checking slice is sorted or not
	fmt.Println(sort.IntsAreSorted(highScore))

	//removing from slice based on index
	var cources = []string{"go", "python", "swift", "java", "ruby"}
	index := 2
	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println(cources)

}
