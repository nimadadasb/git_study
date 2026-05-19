package main

import "fmt"

func main() {

	var age int
	fmt.Println("请输入你的年龄:")
	fmt.Scanln(&age)

	/**
	Go 中 fallthrough 的用法和示例
	1. 基本概念
	在 Go 的 switch 语句中：
	默认情况下，匹配到一个 case 后会自动跳出​ switch（不执行后面的 case）
	使用 fallthrough关键字可以继续执行下一个 case
	*/
	switch {
	case age <= 0:
		fmt.Println("年龄错误")
		fmt.Println("请输入正确的年龄")
	case age <= 12:
		fmt.Println("儿童")
		fallthrough
	case age <= 18:
		fmt.Println("青少年")
		if age == 18 {
			fmt.Println("恭喜成年！")
		}
		fallthrough
	case age <= 60:
		fmt.Println("成年人")
	default:
		fmt.Println("老年人")
	}

	var week int
	fmt.Println("请输入星期：")
	fmt.Scanln(&week)
	switch week {
	case 1:
		fmt.Printf("今天星期%d\n", week)
	case 2:
		fmt.Printf("今天星期%d\n", week)
	case 3:
		fmt.Printf("今天星期%d\n", week)
	case 4:
		fmt.Printf("今天星期%d\n", week)
	case 5:
		fmt.Printf("今天星期%d\n", week)
	case 6:
		fmt.Printf("今天星期%d\n", week)
	case 7:
		fmt.Printf("今天星期%d\n", &week)
	default:
		fmt.Println("输入错误")
	}
}
