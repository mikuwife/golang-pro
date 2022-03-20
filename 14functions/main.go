package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	sum, _ := proAdder(1, 2, 3, 4)
	fmt.Println("Sum is:", sum)
	//fmt.Println("res string is:", res)
}

func adder(x int, y int) int {
	return x + y
}

// 函数参数不确定时===返回不同类型的值
func proAdder(values ...int) (int, string) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum, "Hi Pro result function"
}

func greeter() {
	fmt.Println("Namstey from golang")
}
