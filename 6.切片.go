package main

import (
	"fmt"
	"slices"
)

func main() {

	/**
	make函数的用法
	*/
	//make(T, length, capacity)
	//T：只能是切片(slice)、映射(map) 或通道(chan)
	//length：长度（必需）
	//capacity：容量（可选，仅对切片和通道有效）

	var nameSli []string
	nameSli = append(nameSli, "dkw")
	fmt.Println(nameSli)

	var nameSli2 []string
	fmt.Println("nameSlice2", nameSli2 == nil, nameSli2, len(nameSli2))

	var nameSli3 []string = []string{}
	//初始化
	fmt.Println("nameSlice3", nameSli3 == nil, nameSli3, len(nameSli3))

	var nameSli4 = []string{}
	fmt.Println("nameSlice4", nameSli4 == nil, nameSli4, len(nameSli4))

	nameSli5 := []string{}
	fmt.Println("nameSlice5", nameSli5 == nil, nameSli5, len(nameSli5))

	nameSli6 := make([]string, 3)
	//通过make函数创建指定长度，指定切片的容量
	fmt.Println("nameSlice6", nameSli6 == nil, nameSli6, len(nameSli6))

	ageSli := make([]int, 3)
	//// 直接知道切片大小的情况,使用短变量声明（:=），一行代码完成声明和初始化
	fmt.Println(ageSli, len(ageSli))

	array := [3]int{7, 8, 9}
	slices1 := array[:]
	slices2 := array[0:3]
	slices3 := array[0:0]
	slices4 := array[0:2]
	slices5 := array[1:3]
	slices6 := array[1:2]
	slices.Sort(slices1)
	fmt.Println(slices1)
	fmt.Println(slices1, slices2, slices3, slices4, slices5, slices6)

	var numberSlice = []int{9, 8, 7}
	slices.Sort(numberSlice) //升序
	fmt.Println(numberSlice)
	slices.Reverse(numberSlice) //降序
	fmt.Println(numberSlice)

}
