package main

import "fmt"

func main() {
	demo1()
	fmt.Print("hello world")

	// 先定义，再赋值
	var name string
	name = "枫枫"
	fmt.Println(name)

	// var 变量名 类型 = "变量值"
	var userName string = "枫枫"
	fmt.Println(userName)

	/**
	var name = "枫枫",变量类型省略
	*/

	/**
	name := "枫枫",简短声明
	*/
}
