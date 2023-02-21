package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang")
	presentTime := time.Now() // of type time
	fmt.Println(presentTime)
	fmt.Printf("%T \n", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //always use this time for syntax of formatting

	createdDate := time.Date(2020, time.October, 10, 23, 17, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
