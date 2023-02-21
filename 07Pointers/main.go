package main

import "fmt"

func main() {
	fmt.Println("In Pointer Tutorial")

	/*This astrik says ptr is a variable of type pointer(*) which points to a memory location holding int. Default vaue is <nil>
	var ptr *int
	fmt.Println(ptr)
	Right now this pointer is not pointing to any memory address.
	*/
	myNumber := 23
	var ptr = &myNumber
	/* This is a pointer that references to a memory*/
	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr) //So *ptr extracts out the value at that pointer.

	*ptr = *ptr * 2
	fmt.Println("New value is", myNumber)
	/* ptr is pointing/referencing at memory location of myNumber right?
	      And  *ptr holds value at ptr, ie. it holds myNumber.
		  What it did was take the value at ptr  and multiply it by 2 and store it as new value at ptr
	*/
}

/*
Why do we use pointers in go?
Ans: When we pass variable / constants , sometimes a copy of that variable gets passed instead of the actual value.
So any change made to that var. using the function wont effect the actual variable value i.e. in parent function its value
wont get effected.
For example in java we cannot directly pass a primitive data type as it would pass its copy and any update made to
tha0t varible passed in the function wont effect the actual variable.
This was not case with arrays or non - primitive data types. This is because they were passed as reference.
So a reference pointing to the memory addres of actual array was passed. So anychange made in that function
to the passed array reference was made on array of that memory address.

Therefore in go, to make sure that on passing the value in a function we get our actual value updated
We use pointers.
Instead of passing the variable, we pass a pointier (reference) pointing to the memory address of that variable.
So in that function we again take hold of the actual value from the passed memory address and update it.

*/
