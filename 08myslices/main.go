package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// fruitList = append(fruitList[1:]) //[Tomato Peach Mango Banana]
	// fruitList = append(fruitList[1:3]) //[Tomato Peach]
	fruitList = append(fruitList[:3]) //[Apple Tomato Peach]
	fmt.Println(fruitList)

	// use make function to create a slices
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 // panic: runtime error: index out of range [4] with length 4

	// if use append() there no error
	// append() 方法会重新分配内存，去包含你的所有内容
	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	// 排序
	sort.Ints(highScores)
	fmt.Println(highScores)

	//判断切片是否是拍好了序的
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index
	var course = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(course)
	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)
}
