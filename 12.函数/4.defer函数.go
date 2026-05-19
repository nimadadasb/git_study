package main

import "fmt"

//关键字 defer 用于注册延迟调用
//这些调用直到 return 前才被执。因此，可以用来做资源清理
//多个defer语句，按先进后出的方式执行，谁离return近谁先执行
//defer语句中的变量，在defer声明时就决定了

func Func() {
	defer fmt.Println("defer2")
	fmt.Println("func")
	defer fmt.Println("defer1")
	return
}

func main() {
	var a int = 100
	defer func() {
		fmt.Println(a)
	}()
	//defer的变量需在defer声明前决定了，所以变量a可以，变量b不行
	defer fmt.Println("defer4")
	Func()
	defer fmt.Println("defer3")
	//var b int = 200
	return
}
