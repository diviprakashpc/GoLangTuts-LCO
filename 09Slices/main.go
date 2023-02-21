package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Lets talk about slices")
	var fruitsList = []string{"Apple", "Orange", "Lemon"}
	fmt.Printf("The type of fruitsList is : %T \n", fruitsList)
	fruitsList = append(fruitsList, "Mango")
	fruitsList = append(fruitsList, "Peach")
	// fmt.Println(fruitsList)
	fruitsList = append(fruitsList[1:3]) // start - end(non inclusive) // ITs prints element with idx 1 , 2. 3 is exclusive.
	fmt.Println(fruitsList)

	highScores := make([]int, 4)
	highScores[0] = 100
	highScores[1] = 945
	highScores[2] = 115
	highScores[3] = 700
	//highScores[4] = 777 // will give error as size is only 4
	highScores = append(highScores, 555, 666, 321)
	// it reallocates the newly required memory.
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// Removing a value from slice based on the index

	var courses = []string{"ReactJS", "Javascript", "Swift", "Python", "Ruby"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}

/*
They are under the hood arrays with abstraction and more features. Mostly used in golang

*/
