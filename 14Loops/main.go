package main

import "fmt"

func main() {
	fmt.Println("This is a loop tutorial hero")
	//days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Println("day is", day, "At index", index)
	// } // ofc we can use comma ok syntax. so if we dont want any value we use underscore.

	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue == 2 {
			goto lco
		}
		if rougeValue == 8 {
			break
		}
		if rougeValue == 3 {
			rougeValue++
			continue
		}
		fmt.Println("Value is", rougeValue)
		rougeValue++
	}

	//GO TO statement
lco:
	fmt.Println("Jumping at learn code online ")
}
