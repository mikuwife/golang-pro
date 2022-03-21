package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coutsename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	//parsing 解析url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())

	// 此时获取的数据格式不够友好
	fmt.Println(result.RawQuery) // coutsename=reactjs&paymentid=ghbj456ghb
	// 将查询数据格式化
	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println(qparams)

	for _, val := range qparams {
		fmt.Println(val)
	}

	// 反过来构建一个url
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "loc.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String() // https://loc.dev/tutcss
	fmt.Println(anotherURL)

}
