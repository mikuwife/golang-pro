package main

import "fmt"

//func main() {
//	// 标准
//	//var s1 []int
//	//s2 := []int{}
//	//
//	//// make() 方法声明
//	//var s3 = make([]int, 0)
//	//fmt.Println(s1, s2, s3) // [] [] []
//	//// 初始化赋值
//	//var s4 = make([]int, 0, 0)
//	//fmt.Println(s4)
//	//s5 := []int{1, 2, 3}
//	//fmt.Println(s5)
//	//// 截取数组做成切片
//	//arr := [5]int{1, 2, 3, 4, 5}
//	//arr1 := arr[1:3]
//	//fmt.Println(arr1) // [2 3]
//
//	arr := [6]int{1, 2, 3, 4, 5, 6}
//	arr1 := arr[2]            // 3
//	arr2 := arr[:]            // [1 2 3 4 5 6]
//	arr3 := arr[0:]           // [1 2 3 4 5 6]
//	arr4 := arr[:4]           // [1 2 3 4]
//	arr5 := arr[1:3]          // [2 3]
//	arr6 := arr[1:3:len(arr)] // [2 3]
//	fmt.Println(arr1)
//	fmt.Println(arr2)
//	fmt.Println(arr3)
//	fmt.Println(arr4)
//	fmt.Println(arr5)
//	fmt.Println(arr6)
//}

func main() {
	s1 := make([]int, 0, 5)
	fmt.Printf("%p\n", &s1) // 0xc000004078
	s2 := append(s1, 1)
	fmt.Printf("%p\n", &s2) //xc000004090

	fmt.Println(s1, s2) // [] [1]
}
