package main

import "fmt"

const MyBirthday = 18

func main() {
	var username string = "mikucat"
	fmt.Println(username)
	fmt.Printf("Variables is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is of types: %T \n", isLoggedIn)

	var float = 6.2
	fmt.Println(float)
	fmt.Printf("Variables is of types: %T \n", float)

	birthday := 2020
	fmt.Println(birthday)
	fmt.Printf("Variables is of types: %T \n", birthday)

	// var (
	// 	a = "stromg"
	// 	b = "b"
	// )
	fmt.Println(MyBirthday)
}
