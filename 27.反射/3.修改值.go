package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改值
// obj接收的是变量的地址,value是要设置的新值
func setValue(obj any, value any) {
	//ValueOf获取对象的值
	//v1 不是字符串，而是指向字符串的指针,所以后面需要使用Elem()解引用，得到真正的字符串变量
	v1 := reflect.ValueOf(obj).Elem()
	v2 := reflect.ValueOf(value)
	//判断修改后的数据与原数据的类型是否相同
	//Elem()的作用是“取指针指向的元素 / 解引用（dereference）”
	if v1.Kind() != v2.Kind() {
		return

	}
	switch v1.Kind() {
	case reflect.String:
		v1.SetString(value.(string))
	case reflect.Int:
		v1.SetInt(v2.Int())
	}
}

func main() {
	var name = "dkw"
	var age = 21
	setValue(&name, "dkw1")
	setValue(&age, 22)
	fmt.Println(name, age)
}
