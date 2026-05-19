package main

import (
	"fmt"
	"time"
)

func awaitAdd(awaitSecond int) func(...int) int {
	time.Sleep(time.Duration(awaitSecond) * time.Second)
	return func(numberSlice ...int) (sum int) {
		for _, v := range numberSlice {
			sum += v
		}
		return sum

	}
}

func Copy(name string) {
	fmt.Printf("Copy:%p\n", &name)
	//&是 取地址运算符，用于获取变量在内存中的地址，返回一个指针。
}

// go所有参数都是值传递，这是在传指针
func Set(name *string) {
	fmt.Printf("Set:%p\n", name)
}

func main() {
	//t1 := time.Now()

	//sum := awaitAdd(2)(1, 2, 3)
	//链式调用
	//第一次调用：awaitAdd(2)
	//传入参数：2
	//返回值：一个函数（类型为 func(...int) int）
	//第二次调用：对返回的函数调用 (1, 2, 3)
	//传入参数：1, 2, 3
	//返回值：6（求和结果）
	//subTime := time.Since(t1)

	//fmt.Println(sum, subTime)

	var name = "dkw"
	fmt.Printf("%p\n", &name)
	Copy(name)
	//拷贝后的内存地址不同于原先的内存地址

	Set(&name)
	//传递指针值后的内存地址等同于原先的内存地址
}
