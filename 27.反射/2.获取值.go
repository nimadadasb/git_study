package main

import (
	"fmt"
	"reflect"
)

// 通过反射获取值
func getValue(obj any) {
	//ValueOf获取对象的值
	v := reflect.ValueOf(obj)
	//Kind判断具体的类型的值
	switch v.Kind() {
	case reflect.String:
		fmt.Println("string", v.String())
	case reflect.Int:
		fmt.Println("int", v.Int())
	case reflect.Struct:
		//结构体比较特殊，没有像string，int那样可以直接拿到的函数
		fmt.Println("struct")
	}
}

func main() {
	getValue(21)
	getValue("dkw")
	getValue(struct {
		Name string
	}{Name: "dkw"})
}
