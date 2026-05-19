package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// 有返回类型声明时 → 必须 return
func plus[T Number](n1, n2 T) T {
	sum := n1 + n2
	fmt.Println(sum)
	return sum
}

// 无返回类型声明时 → 不能有 return
// func plus[T Number](n1, n2 T) {    // 无返回值
//
//		sum := n1 + n2
//		fmt.Println(sum)
//		// 不能有 return
//	}

// 带两个泛型参数
func myPrint[T int, K string](a T, b K) (T, K) {
	fmt.Printf("int value: %v, string value: %v\n", a, b)
	return a, b
}

// 泛型结构体
type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type User struct {
	Name string `json:"name"`
	Age  int    ` json:"age"`
	Sex  string `json:"sex"`
}

type userInfo struct {
	Name     string
	Age      int `json:"age,omitempty"` //忽略空值,不希望出现零值问题
	Sex      string
	Password string `json:"-"` //敏感信息,转json的时候忽略该字段
}

func main() {
	plus(1, 2)
	var u1, u2 = uint(2), uint(3)
	plus(u1, u2)
	myPrint(42, "hello")

	userRes := Response[User]{
		Code: 1,
		Msg:  "ok",
		Data: User{
			Name: "dkw",
			Age:  21,
			Sex:  "male",
		},
	}

	userInfoRes := Response[userInfo]{
		Code: 1,
		Msg:  "ok",
		Data: userInfo{
			Name:     "dkw",
			Age:      21,
			Sex:      "male",
			Password: "11102538sb",
		},
	}

	byteDate, _ := json.Marshal(userRes)
	byteUser, _ := json.Marshal(userInfoRes)
	fmt.Println(string(byteDate))
	fmt.Println(string(byteUser))

	//反序列化
	var response Response[User]

	//正确的JSON - 注意"male"后面没有逗号,属性名全部小写，且用双引号括起来
	jsonStr := `{
		"code": 1,
		"msg":  "ok",
		"data": {
			"name": "dkw",
			"age":  21,
			"sex":  "male"
		}
}`

	err := json.Unmarshal([]byte(jsonStr), &response)
	if err != nil {
		log.Fatalf("JSON解析失败: %v", err)
	}

	fmt.Printf("解析成功: %+v\n", response.Data.Name)

	// 泛型切片
	type MySlice[T any] []T
	var mySlice = MySlice[int]{1, 2, 3}
	fmt.Println(mySlice[0] + 1)

	//泛型Map,map的key只能是基本数据类型
	type MyMap[T int, K string] map[T]K
	var myMap = MyMap[int, string]{
		1: "1",
		2: "2",
	}
	fmt.Println(myMap)

}
