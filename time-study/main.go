package main

import (
	"fmt"
	"time"
)

func main() {
	// Welcome to time study
	current_time := time.Now()
	// fmt.Println(current_time)
	fmt.Println(current_time.Format("01-01-2022 Monday 12:00:"))

	createDate := time.Date(2022, time.April, 12, 12, 12, 12, 12, time.UTC)
	fmt.Println(createDate)
	fmt.Printf("%T", createDate)
}
