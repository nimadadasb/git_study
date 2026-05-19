package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Name1 string `big:"-"`
	Name2 string
}

func Setstruct(obj any) {
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		big := t.Field(i).Tag.Get("big")
		//过滤没有big标签的值
		if big == "" {
			continue
		}
		value.SetString(strings.ToUpper(value.String()))
	}

}

func main() {
	s := User{
		Name1: "dkw1",
		Name2: "dkw2",
	}
	Setstruct(&s)
	fmt.Println(s)
}
