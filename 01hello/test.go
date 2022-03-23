package main

import "fmt"

func main() {
	// 取反
	//fmt.Println(!true)
	// 逻辑与
	//fmt.Println(true && false)
	// 逻辑或
	//fmt.Println(true || false)
	// 位运算
	a := 20             // 10100
	b := 30             // 11110
	fmt.Println(a & b)  // 10100 = 20
	fmt.Println(a | b)  // 11110 = 30
	fmt.Println(a ^ b)  // 01010 = 10
	fmt.Println(a << 2) // 1010000 = 80
	fmt.Println(a >> 2) // 101 = 50

}
