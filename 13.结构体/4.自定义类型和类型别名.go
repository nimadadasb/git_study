package main

import "fmt"

type Mycode int
type Yourcode = int

// 自定义类型可以绑方法
func (m Mycode) Code() {

}

//类型别名不可以绑方法
//func(y Yourcode) Code(){
//
//}

const myCode Mycode = 1
const yourCode Yourcode = 1

func main() {
	fmt.Printf("%v,%T\n", myCode, myCode)     //自定义类型的打印类型就是自定义类型本身
	fmt.Printf("%v,%T\n", yourCode, yourCode) //类型别名的打印类型还是原始类型，这里就是int
	
}
