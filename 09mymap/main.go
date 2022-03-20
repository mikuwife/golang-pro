package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)

	// loops
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// 按照指定顺序遍历map
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)

	for i:=0; i<100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}


}
