package main

import "fmt"

type Class struct {
	className string
}
type Student struct {
	Class //组合
	Name  string
}

// Study method与receiver，(s Student)就是接收者，这个方法属于 Student 类型，使得调用方式为s1.Study()，而不是Study(s1)
func (s Student) Study() {
	fmt.Printf("%s 正在学习\n", s.Name)
}

// Info 值接收者
func (s Student) Info() {
	fmt.Printf("名字: %s 班级:%s\n", s.Name, s.className)
}

// SetName 指针接收者
func (s *Student) SetName(name string) {
	s.Name = name

}

/**
用 值接收者 当：
struct 很小
不需要修改原对象
只是读取数据
 用 指针接收者 当：
需要修改原对象
struct 很大
想避免拷贝
*/

func main() {
	c1 := Class{
		className: "123",
	}
	s1 := Student{
		Name:  "dkw",
		Class: c1,
	}
	/**
	不能写成下面这种看起来理所当然的格式：
		s1 := Student{
		Name:  "dkw",
		classname: "123",
	因为go是强类型+显示构造，Student的真实构造为：
	Student
	├── Class
	│   └── className string
	└── Name string
	}
	s.className能直接访问，是因为字段提升（Field Promotion）
	*/
	s1.Study()
	s1.Info()

	s1.SetName("dkw1")
	s1.Info()

}
