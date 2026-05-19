package main

import (
	"fmt"
	"runtime/debug"
)

func read() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			debug.PrintStack()

		}
	}()
	var list []int = []int{1, 2}
	//故意制造panic
	fmt.Println(list[2]) //数组越界

}

func main() {
	read()

	//正常逻辑
	fmt.Println("hello")
}
