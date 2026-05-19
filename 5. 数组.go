package main

import "fmt"

func main() {

	/**
	数组
	*/
	var nameArr [3]string = [3]string{"张三",
		"李四",
		"王五",
	}
	fmt.Println(nameArr)
	fmt.Println(nameArr[0])
	fmt.Println(nameArr[len(nameArr)-1])
	var namelength = len(nameArr)
	fmt.Println(namelength)

	nameArr[0] = "dkw"
	fmt.Println(nameArr)

}
