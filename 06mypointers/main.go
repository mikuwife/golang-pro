package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	var ptr *int
	fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	var ptr1 = &myNumber
	// fmt.Println("Value of actual pointer is", ptr1)
	// fmt.Println("Value of actual pointer is", *ptr1)

	*ptr1 = *ptr1 * 2
	fmt.Println("New value is:", myNumber)
}
