package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch Case tutorial ")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of Dice is : ", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move two spots")
	case 3:
		fmt.Println("You can move three spots")
		fallthrough
	case 4:
		fmt.Println("You can move four spots")
		fallthrough
	case 5:
		fmt.Println("You can move five spots")

	case 6:
		fmt.Println("You can move six spots and roll dice again")

	default:
		fmt.Println("What was that ")

	}
}
