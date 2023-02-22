package main

import "fmt"

func main() {
	fmt.Println("Welcome to If-Else tutorial")
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 && loginCount < 20 {
		result = "Special User"
	} else {
		result = "VIP User"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is not less than 10")
	}

	// if err != nil {

	// } else {

	// }

}
