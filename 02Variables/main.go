package main

import "fmt"

//jwtToken :=3000000 //We cannot allowed this type of declaration in global scope
const loginToken string = "88045695464754654643"

func main() {
	var username string = "divya"
	fmt.Println(username) //fp shortcut
	fmt.Print("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn) //fp shortcut
	fmt.Print("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float32 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	//default valurd and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)

	// implicit type

	var website = "learncodeonline.com"
	fmt.Println(website) // lexer decides type for us
	//	website = 5 // cant do this unlike js, as lexer has already given it a type nad go is very type sensitive.

	//  no var style

	numberOfUser := 300000.0 // lexer auto detect type
	fmt.Println(numberOfUser)

	fmt.Println(loginToken)
	fmt.Printf("Variable is of type : %T \n", loginToken)
}
