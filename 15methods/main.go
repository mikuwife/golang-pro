package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	xiaoming := User{"xiaoming", "ming@gmail", true, 12}
	fmt.Printf("xiaoming: %v\n", xiaoming)

	fmt.Printf("xiaoming details are: %+v\n", xiaoming)
	fmt.Printf("Name is %v and email is %v. \n", xiaoming.Name, xiaoming.Email)

	xiaoming.GetStatus()
	xiaoming.NewMail()
	fmt.Printf("Name is %v and email is %v. \n", xiaoming.Name, xiaoming.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
