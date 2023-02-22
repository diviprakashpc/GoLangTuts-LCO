package main

import "fmt"

func main() {
	fmt.Println("This is how you deal with functions in Go")
	greeter()
	result := adder(3, 4)
	pro_result, pro_string := proadder(3, 4, 5, 6, 7, 8, 9) //comma ok
	fmt.Println("Result is", result)
	fmt.Println("Pro result is", pro_result)
	fmt.Println("Pro message is", pro_string)

}

func adder(x int, y int) int {
	z := x + y
	return z
}

func proadder(values ...int) (int, string) { //Variadic functions
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Pro result function"
}

func greeter() {
	fmt.Println("Namastey from Golang")
}

/*
We use func keyword and then function name and then return type.
You need to call the function to execute it, except main.
Cannont declare function inside a function.

There are immediately invoked fns and lambda fns as well.

*/
