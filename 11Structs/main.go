package main

import "fmt"

func main() {
	fmt.Println("Welcome to Struct Tutorial")
	Divya := User{"Divya", "diviprakash3@gmail.com", true, 22}
	fmt.Println(Divya)
	fmt.Printf("Divya details are : %+v\n", Divya) // Incase of struct we use +v for complete struct showing fields with their values.
	fmt.Printf("Divya's name is %v and his email is %v\n", Divya.Name, Divya.Email)
}

type User struct { //capital U
	Name   string
	Email  string // If field / method has capital letter starting it is has public modifier
	Status bool
	Age    int
}

/*
Like classes that are in other languages
There is no inheritence in golang
NO concept of Super and parent.
*/
