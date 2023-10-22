package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("go files manage")

	// createFile()
	readFile()
}

func createFile() {
	content := "It's a sample content need to go inside the file"

	file, error := os.Create("./sample.txt")

	if error != nil {
		panic(error)
	} else {
		length, error := io.WriteString(file, content)
		if error != nil {
			panic(error)
		} else {
			fmt.Println("lenthe of content", length)
		}
	}
	defer file.Close()
}

func readFile() {

	// for read we need to use ioutils
	content, error := os.ReadFile("./sample.txt")
	if error != nil {
		panic(error)
	} else {
		// directly calling from string
		fmt.Println(string(content))
		//
	}
}
