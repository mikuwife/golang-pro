package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")

	//PerformGetRequest()
	//PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	// 使用这个方法-- strings.Builder --
	//可以更简单的处理 byte转换成string读取
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())

	//fmt.Println(content) // 一段byte数据
	//fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload
	resquestBody := strings.NewReader(`
		{
			"coursename": "Let's go with golang",
			"price": 0,
			"platform": "learnCodeOnline.in"
		}
    `)

	response, err := http.Post(myurl, "application/json", resquestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// form data
	data := url.Values{}
	data.Add("firstname", "miku")
	data.Add("lastname", "cat")
	data.Add("email", "251@qq.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
