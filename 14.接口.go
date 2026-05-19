package main

import "fmt"

type SingInterface interface {
	Sing()
	//接口里的 Sing()，不是“声明方法”，而是“约束行为”
	//“任何类型，只要拥有一个 无参数、无返回值的 Sing()方法，它就是 SingInterface”
}

// 第一个实现接口的类型
type Chicken struct {
	Name string
}

// 第二个实现接口的类型
type Cat struct {
	Name string
}

// 普通函数(function)
/**
“有一个函数，恰好接收 Chicken”
属于 Chicken
接口看不到
不能多态
*/
//func Sing(ch *Chicken) {
//	fmt.Println(ch.Name, "在唱歌")
//}

// Sing 方法(method)声明：所有 *Chicken类型的变量，都有一个叫 Sing的方法
/**
“唱歌是 Chicken的一种行为能力”
可以被接口识别
可以参与多态
面向对象风格
*/
func (ch *Chicken) Sing() {
	fmt.Println(ch.Name, "在唱歌")
}

func (ca *Cat) Sing() {
	fmt.Println(ca.Name, "在唱歌")
}

// 调用接口，只写一次
func isSing(sing SingInterface) {
	//类型断言
	//check, ok := sing.(*Cat)
	//fmt.Println(check, ok)

	switch server := sing.(type) {
	case *Chicken:
		fmt.Println("chicken", server)
	case *Cat:
		fmt.Println("cat", server)
	default:
		fmt.Println("other")
	}

	sing.Sing()
}

// any=interface{},空接口可以传递任何类型
func Print(val any) {
	fmt.Println(val)
}

func main() {
	ch := Chicken{
		Name: "ikun",
	}
	ca := Cat{
		Name: "hajimi",
	}
	ch.Sing()
	ca.Sing()
	isSing(&ch)
	isSing(&ca)
	Print(ca)
}
