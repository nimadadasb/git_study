package main

import (
	"fmt"
	"reflect"
)

// 类型判断
func getType(obj any) {
	//TypeOf获取对象的类型
	t := reflect.TypeOf(obj)
	//Kind判断具体的类型的值
	switch t.Kind() {
	case reflect.String:
		fmt.Println("string")
	case reflect.Int:
		fmt.Println("int")
	case reflect.Struct:
		fmt.Println("struct")
	}
}

func main() {
	getType(21)
	getType("dkw")
	getType(struct {
		Name string
	}{})
}
