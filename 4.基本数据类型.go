package main

import "fmt"

func main() {

	/**
	整数型
	*/
	//var age = 12

	var u8 uint8 = 255
	//	uint8表示定义的数据范围为0~255，255=2的8次方-1,uint8也叫做byte ，一个字节=8个bit位
	fmt.Println(u8)
	var i8 int8 = -128
	//  int8表示定义的数据范围为-128~127，也就是-2的7次方~2的7次方-1
	fmt.Println(i8)

	/**
	浮点型
	*/

	var f1 float32 = -32768
	//  float42最大范围约为3.4e38
	fmt.Println(f1)

	var f2 float64 = 567890
	//  float64最大范围约为1.8e308
	fmt.Println(f2)

	/**
	字符型
	*/

	var a byte = 'a' //ascii里面的字符
	fmt.Printf("%c %d\n", a, a)

	var a1 uint8 = 97
	fmt.Printf("%c %d\n ", a1, a1)

	/**
	字符串类型
	*/

	var z rune = '一'
	fmt.Printf("%c %d\n", z, z)

	//转义字符
	fmt.Println("枫枫\t知道")              // 制表符
	fmt.Println("枫枫\n知道")              // 回车
	fmt.Println("\"枫枫\"知道")            // 双引号
	fmt.Println("枫枫\r知道")              // 回到行首
	fmt.Println("C:\\pprof\\main.exe") // 反斜杠

	//多行字符串
	var s = `今天
天气
真好
`
	fmt.Println(s)

	/**
	布尔型
	*/

	//布尔型数据只有 true（真）和 false（假）两个值
	//布尔类型变量的默认值为false
	//Go 语言中不允许将整型强制转换为布尔型
	//布尔型无法参与数值运算，也无法与其他类型进行转换

	var b bool = true
	fmt.Println(b)

	//var c = b + 1
	//fmt.Println(c)

	/**
	零值问题
	*/
	//如果我们给一个基本数据类型只声明不赋值
	//那么这个变量的值就是对应类型的零值，例如int就是0，bool就是false，字符串就是""
	var a0 int
	var a2 float32
	var a3 string
	var a4 bool

	fmt.Printf("%#v\n", a0)
	fmt.Printf("%#v\n", a2)
	fmt.Printf("%#v\n", a3)
	fmt.Printf("%#v\n", a4)
}
