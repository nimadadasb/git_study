package main

import "fmt"

func main() {
	fmt.Print("请输入你的名字：")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println("你的名字为：", name)
	fmt.Print("请输入你的年龄：")
	var age int
	n, err := fmt.Scanf("%d", &age)
	fmt.Print("变量数量为：", n, "\n", "有无报错：", err, "\n", "你的年龄为：", age)

}
