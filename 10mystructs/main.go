package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	xiaoming := User{"xiaoming", "ming@gmail", true, 12}
	fmt.Printf("xiaoming: %v\n", xiaoming)

	fmt.Printf("xiaoming details are: %+v\n", xiaoming)
	fmt.Printf("Name is %v and email is %v", xiaoming.Name, xiaoming.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
