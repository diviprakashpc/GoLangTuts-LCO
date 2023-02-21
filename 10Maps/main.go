package main

import (
	"fmt"
)

func main() {
	fmt.Println("Lets learn about maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages", languages)
	fmt.Println("JS is shorthand for ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages", languages)

	// Loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For Key %v Value is %v\n", key, value)
	}
	//put underscore in place of this  comma ok syntax for value you dont want.
}
