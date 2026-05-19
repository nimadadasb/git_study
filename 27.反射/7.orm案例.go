package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type ClassModel struct {
	Name string `orm:"name"`
	Id   int    `orm:"id"`
}

func Find(obj any, query ...any) (sql string, err error) {
	//obj必须是结构体
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		err = errors.New("obj必须是结构体")
		return
	}

	var where string
	//验证条件
	if len(query) > 0 {
		q := query[0]
		//验证第一个参数的类型，必须是字符串
		qs, ok := q.(string)
		if !ok {
			err = errors.New("查询的第一个参数必须是字符串")
			return
		}
		//算问号的个数
		if strings.Count(qs, "?")+1 != len(query) {
			err = errors.New("查询参数个数不匹配")
			return
		}

		//拼接where
		for _, a := range query[1:] {
			switch s := a.(type) {
			case string:
				strings.Replace(qs, "?", fmt.Sprint("'%s", s), 1)
			case int:
				qs = strings.Replace(qs, "?", fmt.Sprint("'%d'", s), 1)
			}

			where = "where " + qs

		}
	}

	//去拼接所有的有orm的字段
	var coloums []string
	for i := 0; i < t.NumField(); i++ {
		orm := t.Field(i).Tag.Get("orm")
		if orm == "" {
			continue
		}
		coloums = append(coloums, orm)

	}

	//算表名，小写结构体的名字s
	name := strings.ToLower(t.Name()) + "s"

	sql = fmt.Sprintf("select %s from %s %s;",
		strings.Join(coloums, ","),
		name,
		where)
	return
}

func main() {
	sql, err := Find(ClassModel{}, "name = ?", "三年一班")
	fmt.Println(sql, err)
	sql, err = Find(ClassModel{}, "id = ? and name = ?", 1, "三年一班")
	fmt.Println(sql, err)
	sql, err = Find(ClassModel{})
	fmt.Println(sql, err)
}
