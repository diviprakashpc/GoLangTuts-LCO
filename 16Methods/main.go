package main

import (
	"fmt"
)

func main() {
	fmt.Println("Lets learn about methods in go")
	Divya := User{"Divya", "diviprakash3@gmail.com", true, 16}
	fmt.Println("Divya's email is", Divya.Email)
	fmt.Println("Divya's age is", Divya.Age)
	Divya.GetStatus()
	Divya.SetNewMail("diviprakashpc@gmail.com")
	fmt.Println("New Divya's details are", Divya)

}

type User struct { // R. Public properties are declared with capital letter
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active : ", u.Status)
}

func (u User) SetNewMail(email string) {
	u.Email = email
	fmt.Println("Divya's new mail is", u.Email)
}

/*
When these functions go into classes or structs here , we call these fns methods.
The SetNewMail did not set. So we can see it is updating a copy and not actual object
*/
