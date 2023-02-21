package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitsList [4]string
	fruitsList[0] = "Apple"
	fruitsList[1] = "Tomato"
	fruitsList[3] = "Mango"
	//fmt.Println("Fruits List is :", fruitsList) //The long space indicates there is something missing.
	//fmt.Println("Fruits List is :", len(fruitsList))
	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("vegy list is ", vegList)

}

//There is no much to do with arrays in golang strangely.
// Minimally used in golang
