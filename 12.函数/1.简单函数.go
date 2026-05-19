package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("hello")
}

// 传入一个参数
func param1(id string) {
	fmt.Println(id)
}

// 传入两个参数
func param2(id string, username string) {
	fmt.Println(id, username)
}

// 传染源一个切片slice
/**
func 函数名(参数名 ...类型) 返回值类型 {
    // 函数体
}
这种参数写法是 Go 语言中的可变参数（variadic parameter）
参数类型前的 ...表示这个函数可以接受零个或多个该类型的参数
在函数内部，可变参数被当作切片来处理
*/
func add(slice ...int) {
	var sum int
	for _, v := range slice {
		sum += v
	}
	fmt.Println(sum)
}

//函数返回值

func return1() {
	//没有返回值
	return

}

func return2() bool {
	//返回一个布尔值
	return true

}

func return3() (string, bool) {
	//有多个返回值
	if 1 > 2 {
		return "", false
	}
	if 1 < 2 {
		return "", true
	}
	return "", false

}

// 命名式函数返回(具名返回值)
//
//	func 函数名(参数) (返回值变量名 类型, ...) {
//	    函数体
//	   return  // 可以省略返回值
//	}
func return4() (val string, ok bool) {
	if 1 < 2 {
		fmt.Println(val, ok)
		return val, true
	}
	return "", false

}

func main() {
	sayHello()
	param1("1")
	param2("1", "dkw")
	add(1, 2)
	return4()

	//匿名函数
	//有返回值,可以用print打印
	var getName = func(name string) string {
		return name
	}
	fmt.Println(getName("dkw"))

	//没有返回值,直接调用该函数
	var setName = func(name string) {
		fmt.Println(name)
	}
	setName("dkw2")

	//高阶函数
	fmt.Println("请输入要执行的操作：")
	fmt.Println(`1：登录
2：注册
3：个人中心`)

	var index int
	fmt.Scanln(&index)

	switch index {
	case 1:
		login()
	case 2:
		register()
	case 3:
		userCenter()
	}

}
func login() {
	fmt.Println("login")
}
func register() {
	fmt.Println("register")
}
func userCenter() {
	fmt.Println("userCenter")
}
