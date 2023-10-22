package main

import "fmt"

func main() {
	var languages = make(map[string]string)
	languages["js"] = "javaScript"
	languages["py"] = "python"
	languages["rb"] = "rubby"
	languages["go"] = "golang"

	fmt.Println(languages)

	// delete from map
	delete(languages, "js")

	// looping map

	for _, value := range languages {
		fmt.Printf("key is: v, Value is %v\n", value)
	}
}
