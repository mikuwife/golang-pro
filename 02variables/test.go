package main

import (
	"fmt"
)

func main() {

	//// 标准声明
	//var name string
	//
	//// 批量声明
	//var (
	//	age       int
	//	firstname string
	//	lastname  string
	//)

	// 初始化
	//var name string = "mikuwife"
	//// 同时初始化多个变量
	//var fullname, age = "miku", 10

	// 零值
	//var age int
	//fmt.Println(age)

	//// 标准
	//const pi = 3.14
	//// 多个常量同时声明
	//const (
	//	p1 = 3.14
	//	p2 = 2
	//)
	//// 省略值
	//const (
	//	c1 = 1
	//	c2 // c2 = 1
	//	c3 // c3 = 1
	//)
	//
	//fmt.Println(c2)
	//fmt.Println(c3)

	//const (
	//	a1 = iota
	//	a2
	//	a3
	//)
	//
	//fmt.Println(a1)
	//fmt.Println(a2)
	//fmt.Println(a3)

	//const (
	//	a, b = iota + 1, iota + 2 // 1 , 2
	//	c, d                      // 2 , 3
	//	e, f                      // 3, 4
	//)
	//
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(d)
	//changeString()

	//s1 := "hello"
	//bytes1 := []byte(s1)
	//fmt.Println(bytes1)
	//bytes1[0] = 'H'
	//fmt.Println(bytes1)
	//fmt.Println(string(bytes1))

	//var a, b = 3, 4
	//var c float64
	//
	//c = math.Sqrt(float64(a*a + b*b))
	//fmt.Println(c)

	a := 3.133
	var b int
	//b = a
	fmt.Println(b) // cannot use a (variable of type float64) as type int in assignment
	b = int(a)
	fmt.Println(b) // 3
}

//func changeString() {
//	s := "pprof.cn博客"
//	for i := 0; i < len(s); i++ { //byte
//		fmt.Printf("%v(%c) ", s[i], s[i])
//	}
//	fmt.Println()
//	for _, r := range s { //rune
//		fmt.Printf("%v(%c) ", r, r)
//	}
//	fmt.Println()
//}
