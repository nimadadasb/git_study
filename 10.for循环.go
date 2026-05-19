package main

import (
	"fmt"
)

func main() {

	var sum int
	for i := 0; i <= 10; i++ {
		sum += i

	}
	fmt.Println(sum)

	//死循环
	//for a := 0; true; a++ {
	//	fmt.Println(time.Now())
	//	time.Sleep(2 * time.Second)
	//}

	//while模式
	var sum1 int
	var i int = 1
	for i <= 100 {
		sum1 += i
		i++
	}
	fmt.Println(sum1)

	//do-while模式
	var sum2 int
	var n int = 1
	for {
		sum2 += n
		n++
		if n > 100 {
			break
		}
	}
	fmt.Println(sum2)

	//遍历切片
	var slice = []string{"a", "b", "c", "d"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(i, slice[i])
	}

	for index, item := range slice {
		fmt.Println(index, item)
	}

	//遍历map
	var userMap = map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}
	for key, value := range userMap {
		fmt.Println(key, value)
	}

	//九九乘法表
	//1*1=1
	//2*1=2  2*2=4
	//...
	//9*1=9  ...  9*9=81

	for number := 1; number <= 9; number++ {
		for j := 1; j < number+1; j++ {
			fmt.Printf("%d*%d=%d\t", number, j, number*j)
		}
		fmt.Println()

	}

}
