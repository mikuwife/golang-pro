package main

import "fmt"

func main() {
	// 声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("空切片")
	} else {
		fmt.Println("非空切片")
	}
	
	// 定义切片
	s2 := []int{}
	fmt.Printf("s2: %v\n", s2)

	var s3 []int = make([]int, 0)
	fmt.Printf("s3: %v\n", s3)
	
	var s4 []int = make([]int, 0, 0)
	s5 := []int{1, 2, 3}
	fmt.Printf("s4: %v\n", s4)
	fmt.Printf("s5: %T\n", s5)
}
