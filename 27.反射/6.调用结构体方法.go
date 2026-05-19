package main

import (
	"fmt"
	"reflect"
)

type User1 struct {
}

// 这里使用User1这个值接收者，而不能使用*User1这个指针接收者
func (User1) Call(name string) {
	fmt.Println("被调用了")
}

func Call(obj any) {
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	//NumMethod用来返回当前类型的方法集（Method Set）中的方法个数，这里的方法集就是User1
	for i := 0; i < v.NumMethod(); i++ {
		//Method用来获取第 i 个方法的元信息
		m := t.Method(i)
		fmt.Println(m.Name)
		if m.Name != "Call" {
			continue
		}
		method := v.Method(i).Call([]reflect.Value{
			reflect.ValueOf("dkw"),
		})
		fmt.Println(method)

	}

}

func main() {
	s := User1{}
	Call(&s)

}
