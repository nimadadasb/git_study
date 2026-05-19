package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	IsMan bool   `json:"is_man"`
}

// ParseJson 获取结构体里的值
func ParseJson(obj any) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	//NumField是用来返回结构体字段个数，因为obj就是我传入的结构体的值，对应下面main函数传入的结构体s
	for i := 0; i < v.NumField(); i++ {
		//reflect.Value.Field(i)函数用来获取结构体第i个字段的值
		tf := t.Field(i)
		jsonTag := tf.Tag.Get("json")
		if jsonTag == "-" {
			jsonTag = tf.Name
		}
		fmt.Println(tf.Name, tf.Tag, jsonTag)
		fmt.Println(v.Field(i))
	}
}

func main() {
	s := Student{
		Name:  "dkw",
		Age:   21,
		IsMan: true,
	}
	ParseJson(s)

}
