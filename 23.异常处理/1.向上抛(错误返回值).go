package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	res = a / b
	return
}

// 命名返回值
func server(a, b int) (res int, err error) {
	// 可以直接使用result, err（已声明）
	res, err = div(a, b)
	if err != nil {
		//把错误向上传递
		return // 等价于 return 0, err
	}

	//执行其它的逻辑

	res += 2
	return // 等价于 return result, nil
}

// 大多数情况下的首选 (匿名返回值)
//func Calculate(a, b int) (int, error) {
//	res, err := div(a, b)  // 短声明创建局部变量
//	if err != nil {
//		return 0, err  // 必须明确返回
//	}
//	res += 10
//	return res, nil
//}

func main() {
	//短声明
	res, err := server(2, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}
