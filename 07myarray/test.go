package main

import "fmt"

//func main() {
//	// 指定类型
//	var a1 [5]int = [5]int{1, 2, 3}
//
//	// 省略类型
//	var a2 = [5]int{1, 2, 3, 4, 5}
//
//	// 根据初始值判断数组大小
//	var a3 = [...]int{1, 2, 3, 4, 5, 6, 7}
//
//	// 给指定下标的数组元素初始化
//	var strs = [5]string{3: "hello", 4: "miku"}
//
//	// 输出
//	fmt.Println(a1)      // [1 2 3 0 0]
//	fmt.Println(a2)      // [1 2 3 4 5]
//	fmt.Println(len(a3)) // 7   [1 2 3 4 5 6 7]
//	fmt.Println(strs[3]) // hello
//
//}

//func main() {
//	var array1 [5][3]int
//	array2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
//	array3 := [...][2]int{{1, 2}, {1, 2}}
//	fmt.Println(array1)
//	fmt.Println(array2)
//	fmt.Println(array3)
//}

//func main() {
//	arr1 := [5]int{1, 2, 3, 4, 5}
//	arr2 := arr1
//	arr2[0] = 2
//	fmt.Printf("arr1: %p\n", &arr1) // 0xc00000e4b0
//	fmt.Printf("arr2: %p\n", &arr2) // 0xc00000e4e0
//	fmt.Println(arr1)               //[1 2 3 4 5]
//	fmt.Println(arr2)               //[2 2 3 4 5]
//}

//func main() {
//	arr1 := [5]int{1, 2, 3, 4, 5}
//	fmt.Println(len(arr1)) // 5
//	fmt.Println(cap(arr1)) // 5
//}

//func main() {
//	a := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
//	for _, val := range a {
//		for _, va := range val {
//			fmt.Println(va)
//		}
//	}
//}

//func printArr(arr *[5]int) {
//
//}
//
//func main() {
//	var arr1 [5]int
//	printArr(&arr1)
//}

//func getSum(arr *[5]int) int {
//	sum := 0
//	for _, val := range arr {
//		sum += val
//	}
//	return sum
//}
//
//func main() {
//	arr1 := [5]int{1, 2, 3, 4, 5}
//	result := getSum(&arr1)
//	fmt.Println(result)
//}

func main() {
	target := 3
	arr1 := [5]int{1, 2, 3, 1, 5}
	for i := 0; i < len(arr1); i++ {
		another := target - arr1[i]
		for j := i + 1; j < len(arr1); j++ {
			if arr1[j] == another {
				fmt.Println(i, j)
			}
		}
	}
}
