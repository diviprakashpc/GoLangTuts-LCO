package main

import "fmt"

func main() {
	fmt.Println("Lets learn about Defer")
	defer fmt.Println("World1")
	fmt.Println("Hello")
	defer fmt.Println("World2")
	defer fmt.Println("World3")
	myDefer() // This line never saw the defer statement so func call will happend after printing hello.
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
		//withour defer order of execution:
		// 0 1 2 3 4
		//with defer
		// defer stack [fmt.Println("0"),fmt.Println("1")......fp("4")]
		//therefore output 4 3 2 1 0
	}
}

/*
When a function is executed, it executes line by line.
When we put a defer at a line, then that line is executed at end of the function and is currently skipped.
The deffered order of the lines is in reversed order i.e LIFO.

So all the non defer lines are executed by their order first,
And then all defer lines are executed in LIFO manner.

*/
