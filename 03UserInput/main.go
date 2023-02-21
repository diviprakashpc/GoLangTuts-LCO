package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var welcome string = "Welcome to User Input"
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Type of reader is : %T \n", reader)
	fmt.Println("Enter the rating of pizza : ")

	//comma ok || error ok

	// WE DONY HAVE TRY CATCH . SO NOTHING TO CATCH IF AN ERROR OCCURS
	// BECAUSE THE DESIGN PARADIGM OF GO EXPECTS US TO TAKE PROBLEM OR ERROR AS VARIABLE

	//input, err := reader.ReadString('\n')
	// If it is successfull the input value comes in input . or if there is error value goes into err
	// or
	input, _ := reader.ReadString('\n')
	// when we dont care about error

	//_, error := reader.ReadString('\n'), _ := reader.ReadString('\n')
	// when we dont care about result.

	// in bracket tells the indicator of when to stop reading. Here till next line is encountered.
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Type of rating is %T \n", input)
}
