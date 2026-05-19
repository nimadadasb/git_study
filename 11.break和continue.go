package main

import "fmt"

func main() {
	for c := 0; c <= 10; c++ {
		fmt.Println(c)
		if c == 5 {
			continue
			//	break 打印到c==5会直接中断
		}
		//fmt.Println(c) 放在这里打印就不会打印5，会跳过

	}

}
