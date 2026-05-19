package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入你的年龄:")
	fmt.Scanln(&age)

	//中断式 卫语句
	//if age <= 0 {
	//	fmt.Println("年龄错误")
	//	return
	//}
	//
	//if age <= 18 {
	//	fmt.Println("未成年")
	//	return
	//}
	//
	//if age > 18 {
	//	fmt.Println("已成年")
	//	return
	//}

	// 使用 else if(推荐)
	if age <= 0 {
		fmt.Println("年龄错误")
	} else if age <= 12 {
		fmt.Println("儿童")
	} else if age <= 18 {
		fmt.Println("青少年")
	} else if age <= 60 {
		fmt.Println("成年人")
	} else {
		fmt.Println("老年人")
	}

}
