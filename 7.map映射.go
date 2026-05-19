package main

import "fmt"

func main() {
	var userMap map[int]string = map[int]string{
		1: "dkw",
		2: "张三",
		4: "",
	}
	fmt.Println(userMap)
	fmt.Println((userMap[1]))
	fmt.Printf("%#v\n", userMap[0])
	// 一行代码判断并获取值
	if value, ok := userMap[1]; ok {
		fmt.Printf("key=1 存在，值是: %s\n", value)
	} else {
		fmt.Printf("key=1 不存在，返回的零值是: %#v\n", value)
	}

	if value, ok := userMap[0]; ok {
		fmt.Printf("key=0 存在，值是: %s\n", value)
	} else {
		fmt.Printf("key=0 不存在，返回的零值是: %#v\n", value)
	}

	if value, ok := userMap[4]; ok {
		fmt.Printf("key=4 存在，值是: %#v\n", value)
	} else {
		fmt.Printf("key=4 不存在，返回的零值是: %#v\n", value)
	}

	if value, ok := userMap[2]; ok {
		fmt.Printf("key=2 存在，值是: %#v\n", value)
	} else {
		fmt.Printf("key=2 不存在，返回的零值是: %#v\n", value)
	}
	if value, ok := userMap[3]; ok {
		fmt.Printf("key=3 存在，值是: %#v\n", value)
	} else {
		fmt.Printf("key=3 不存在，返回的零值是: %#v\n", value)
	}

	userMap[1] = "dkw2"
	delete(userMap, 4)
	fmt.Println(userMap)

}
