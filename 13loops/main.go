package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for 循环遍历数组
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// for range
	// 先索引值，后值
	for _, day := range days {
		fmt.Println(day)
	}

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			rougueValue++
			//break
			continue
		}
		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

	//跳转到指定地区
lco:
	fmt.Println("Jumping at LearnCodeonline.in")
}
